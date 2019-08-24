package main

import m "github.com/elchulito88/utils/modifier"

func main() {
	path := m.Paths{"test"}
	var p m.FileManipulator
	p = path
	p.MkDir()
	p.RemoveDir()

}
