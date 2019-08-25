package modifier

import (
	"io/ioutil"
	"os"

	l "github.com/elchulito88/utils/logging"
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
		l.Log(err2)
	}
}

//MkDir is a method for creating a Directory
func (p Paths) MkDir() {
	RemovePath(p.Path)
	err := os.Mkdir(p.Path, os.ModePerm)
	l.Log(err)
}

//RemoveDir is a method for creating a Path
func (p Paths) RemoveDir() {
	RemovePath(p.Path)
}

//RemoveFile removes file
func (p Paths) RemoveFile() {
	_, err := os.Stat(p.Path)

	if err == nil {
		errR := os.Remove(p.Path)
		l.Log(errR)
	}
}

//MkFile is used to make and save a file
func (p Paths) MkFile(obj []byte) {
	err := ioutil.WriteFile(p.Path, []byte(obj), os.ModePerm)
	l.Log(err)
}
