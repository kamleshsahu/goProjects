package interfaces

type BlockInterface interface {
	FRead(fd int, size int) ([]byte, error)
	FWrite(fd int, data []byte) error
	FOpen(filename string, mode string) (int, error)
	FClose(fd int) error
	Remove(Filename string) error
	Rename(oldName, newName string) error
}
