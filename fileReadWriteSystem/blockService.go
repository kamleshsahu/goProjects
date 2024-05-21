package main

import (
	"awesomeProject/fileReadWriteSystem/entity"
	"awesomeProject/fileReadWriteSystem/interfaces"
	"bytes"
	"encoding/gob"
	"errors"
	"os"
)

type blockService struct {
	BD *entity.BlockDevice
}

const (
	BlockSize   = 64
	NumBlocks   = 10
	MagicNumber = "VFS"
)

func (bs *blockService) FOpen(Filename string, mode string) (int, error) {
	//bs.BD.Lock.Lock()
	//defer bs.BD.Lock.Unlock()

	inode, exists := bs.BD.Inodes[Filename]
	if !exists {
		if mode == "r" {
			return 0, errors.New("file not found")
		}

		blocks, err := bs.allocateBlocks(1)
		if err != nil {
			return 0, err
		}

		inode = entity.Inode{
			Filename:   Filename,
			Size:       0,
			StartBlock: blocks[0],
			BlockCount: 1,
		}
		bs.BD.Inodes[Filename] = inode
	}

	fd := bs.BD.FileCounter
	bs.BD.FileCounter++
	bs.BD.FdTable[fd] = &entity.FileDescriptor{
		Inode:    &inode,
		Position: 0,
		Mode:     mode,
	}

	return fd, nil
}

func (bs *blockService) FRead(fd int, size int) ([]byte, error) {
	bs.BD.Lock.Lock()
	defer bs.BD.Lock.Unlock()

	descriptor, exists := bs.BD.FdTable[fd]
	if !exists {
		return nil, errors.New("invalid file descriptor")
	}

	inode := descriptor.Inode
	data := make([]byte, 0, size)
	totalRead := 0
	startBlock := descriptor.Position / bs.BD.BlockSize
	startOffset := descriptor.Position % bs.BD.BlockSize

	for i := startBlock; i < inode.BlockCount; i++ {
		blockNum := inode.StartBlock + i
		blockData, err := bs.readBlock(blockNum)
		if err != nil {
			return nil, err
		}

		if i == startBlock {
			blockData = blockData[startOffset:]
		}

		bytesToRead := len(blockData)
		if totalRead+bytesToRead > size {
			bytesToRead = size - totalRead
		}

		data = append(data, blockData[:bytesToRead]...)
		totalRead += bytesToRead

		if totalRead >= size {
			break
		}
	}
	println("reading from blocks", startBlock)

	descriptor.Position += totalRead
	return data[:totalRead], nil
}

func (bs *blockService) FWrite(fd int, data []byte) error {
	bs.BD.Lock.Lock()
	defer bs.BD.Lock.Unlock()

	descriptor, exists := bs.BD.FdTable[fd]
	if !exists {
		return errors.New("invalid file descriptor")
	}

	inode := descriptor.Inode
	bytesWritten := 0
	dataSize := len(data)

	requiredBlocks := (dataSize + bs.BD.BlockSize - 1) / bs.BD.BlockSize
	blocks, err := bs.allocateBlocks(requiredBlocks)
	if err != nil {
		return err
	}

	println("writing to blocks", blocks[0])

	for _, blockNum := range blocks {
		blockData := data[bytesWritten:]
		if len(blockData) > bs.BD.BlockSize {
			blockData = blockData[:bs.BD.BlockSize]
		}

		err = bs.writeBlock(blockNum, blockData)
		if err != nil {
			return err
		}

		bytesWritten += len(blockData)
	}

	inode.BlockCount += requiredBlocks
	inode.Size += bytesWritten
	bs.BD.Inodes[inode.Filename] = *inode
	return bs.saveInodes()
}

func (bs *blockService) FClose(fd int) error {
	bs.BD.Lock.Lock()
	defer bs.BD.Lock.Unlock()

	_, exists := bs.BD.FdTable[fd]
	if !exists {
		return errors.New("invalid file descriptor")
	}

	delete(bs.BD.FdTable, fd)
	return nil
}

func (bs *blockService) Rename(oldName, newName string) error {
	bs.BD.Lock.Lock()
	defer bs.BD.Lock.Unlock()

	inode, exists := bs.BD.Inodes[oldName]
	if !exists {
		return errors.New("file not found")
	}

	inode.Filename = newName
	bs.BD.Inodes[newName] = inode
	delete(bs.BD.Inodes, oldName)
	return bs.saveInodes()
}

func (bs *blockService) Remove(Filename string) error {
	bs.BD.Lock.Lock()
	defer bs.BD.Lock.Unlock()

	inode, exists := bs.BD.Inodes[Filename]
	if !exists {
		return errors.New("file not found")
	}

	for i := 0; i < inode.BlockCount; i++ {
		blockNum := inode.StartBlock + i
		delete(bs.BD.BlockMap, blockNum)
		bs.BD.FreeBlocks = append(bs.BD.FreeBlocks, blockNum)
	}

	delete(bs.BD.Inodes, Filename)
	return bs.saveInodes()
}

func (bs *blockService) readBlock(blockNum int) ([]byte, error) {
	file, err := os.Open(bs.BD.Filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.Seek(int64(blockNum*bs.BD.BlockSize+len(MagicNumber)), 0)
	if err != nil {
		return nil, err
	}

	data := make([]byte, bs.BD.BlockSize)
	n, err := file.Read(data)
	if err != nil {
		return nil, err
	}
	if n < bs.BD.BlockSize {
		return data[:n], nil
	}

	return data, nil
}

func (bs *blockService) writeBlock(blockNum int, data []byte) error {
	file, err := os.OpenFile(bs.BD.Filename, os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Seek(int64(blockNum*bs.BD.BlockSize+len(MagicNumber)), 0)
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (bs *blockService) allocateBlocks(count int) ([]int, error) {
	//bs.BD.Lock.Lock()
	//defer bs.BD.Lock.Unlock()

	if len(bs.BD.FreeBlocks) < count {
		return nil, errors.New("not enough free blocks available")
	}

	blocks := bs.BD.FreeBlocks[:count]
	bs.BD.FreeBlocks = bs.BD.FreeBlocks[count:]

	for _, blockNum := range blocks {
		bs.BD.BlockMap[blockNum] = true
	}

	return blocks, nil
}

func (bs *blockService) saveInodes() error {
	file, err := os.OpenFile(bs.BD.Filename, os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Seek(int64(bs.BD.NumBlocks*bs.BD.BlockSize+len(MagicNumber)), 0)
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err = encoder.Encode(bs.BD.Inodes)
	if err != nil {
		return err
	}

	_, err = file.Write(buffer.Bytes())
	return err
}

func (bs *blockService) loadInodes() error {
	file, err := os.Open(bs.BD.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Seek(int64(bs.BD.NumBlocks*bs.BD.BlockSize+len(MagicNumber)), 0)
	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(file)
	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(&buffer)
	err = decoder.Decode(&bs.BD.Inodes)
	return err
}

func NewBlockDevice(Filename string) *entity.BlockDevice {
	bs := &entity.BlockDevice{
		Filename:    Filename,
		BlockSize:   BlockSize,
		NumBlocks:   NumBlocks,
		FreeBlocks:  make([]int, NumBlocks),
		BlockMap:    make(map[int]bool),
		Inodes:      make(map[string]entity.Inode),
		FdTable:     make(map[int]*entity.FileDescriptor),
		FileCounter: 0,
	}

	for i := 0; i < NumBlocks; i++ {
		bs.FreeBlocks[i] = i
	}

	file, err := os.OpenFile(Filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if stat.Size() == 0 {
		// Initialize the file with empty blocks
		file.WriteString(MagicNumber)
		file.Write(make([]byte, BlockSize*NumBlocks))
	}

	return bs
}

func NewBlockService(name string) interfaces.BlockInterface {
	return &blockService{BD: NewBlockDevice(name)}
}
