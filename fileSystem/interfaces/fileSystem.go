package interfaces

type Filesystem interface {
	Ls()
	Add(filesystem Filesystem)
}
