package service

import (
	error2 "awesomeProject/error"
	"awesomeProject/fileSystem/interfaces"
	"fmt"
)

type file struct {
	name string
}

func (f *file) Add(filesystem interfaces.Filesystem) {
	//TODO implement me
	panic("implement me")
}

func (f *file) Ls() {
	fmt.Println("File name: ", f.name)
}

func NewFile(name string) (interfaces.Filesystem, error) {
	return nil, &error2.CustomError{Message: "fileNotFound"}
}
