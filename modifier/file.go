package modifier

import (
	"log"
	"os"
)

//FileManipulator interface
type FileManipulator interface {
	MkDir()
	RemoveDir()
	RemoveFile()
}

//Paths is a file string
type Paths struct {
	Path string
}

//RemovePath is a function for removing a path
func RemovePath(path string) {
	_, err := os.Stat(path)

	if err == nil {
		err2 := os.RemoveAll(path)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
}

//MkDir is a method for creating a Directory
func (p Paths) MkDir() {
	RemovePath(p.Path)
	os.Mkdir(p.Path, 0755)
}

//RemoveDir is a method for creating a Path
func (p Paths) RemoveDir() {
	RemovePath(p.Path)
}

//RemoveFile removes file
func (p Paths) RemoveFile() {
	_, err := os.Stat(p.Path)

	if err == nil {
		err2 := os.Remove(p.Path)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
}
