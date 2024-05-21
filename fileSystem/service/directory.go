package service

import (
	"awesomeProject/fileSystem/interfaces"
	"fmt"
)

type directory struct {
	name  string
	files []interfaces.Filesystem
}

func (d *directory) Ls() {
	fmt.Println("Directory name: ", d.name)

	for _, file := range d.files {
		file.Ls()
	}
}

func (d *directory) Add(filesystem interfaces.Filesystem) {
	d.files = append(d.files, filesystem)
}

func NewDirectory(name string) interfaces.Filesystem {
	return &directory{name: name, files: make([]interfaces.Filesystem, 0)}
}
