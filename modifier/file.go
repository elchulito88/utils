package modifier

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"log"

	l "github.com/elchulito88/utils/logging"
	s "github.com/elchulito88/utils/ssh"
)

//FileManipulator interface
type FileManipulator interface {
	MkDir()
	RemoveDir()
	RemoveFile()
	MkFile(obj string)
	MvFile(obj string)
	CopyPath(obj string) (int64, error)
	CreateSSHKey(obj string)
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
func (p Paths) MkFile(obj string) {
	err := ioutil.WriteFile(p.Path, []byte(obj), os.ModePerm)
	l.Log(err)
}

//MvFile moves files across different locations
func (p Paths) MvFile(obj string) {
	err := os.Rename(p.Path, obj)
	l.Log(err)
}

//CopyPath is used to copy files from one location to the next
func (p Paths) CopyPath(dst string) (int64, error) {
	sourceFileStat, err := os.Stat(p.Path)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", p.Path)
	}

	source, err := os.Open(p.Path)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nbytes, err := io.Copy(destination, source)
	return nbytes, err
}

// CreateSSHKey is used to create SSH Keys
func (p Paths) CreateSSHKey(src string) {

	savePrivateFileTo := "./" + p.Path
	savePublicFileTo := "./" + p.Path + ".pub"

	bitSize := 4096

	privateKey, err := s.GeneratePrivateKey(bitSize)
	if err != nil {
		log.Fatal(err.Error())
	}

	publicKeyBytes, err := s.GeneratePublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	privateKeyBytes := s.EncodePrivateKeyToPEM(privateKey)

	err = s.WriteKeyToFile(privateKeyBytes, savePrivateFileTo)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = s.WriteKeyToFile([]byte(publicKeyBytes), savePublicFileTo)
	if err != nil {
		log.Fatal(err.Error())
	}
}