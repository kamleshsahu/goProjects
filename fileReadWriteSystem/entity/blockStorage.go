package entity

import "sync"

type Inode struct {
	Filename   string
	Size       int
	StartBlock int
	BlockCount int
}

type FileDescriptor struct {
	Inode    *Inode
	Position int
	Mode     string
}

type BlockDevice struct {
	Filename    string
	BlockSize   int
	NumBlocks   int
	FreeBlocks  []int
	BlockMap    map[int]bool
	Inodes      map[string]Inode
	FdTable     map[int]*FileDescriptor
	FileCounter int
	Lock        sync.Mutex
}
