package siopao

import (
	"os"
	"path/filepath"
)

// Move renames, or moves the file to another path. This is a more direct approach, and will be able to
// move the file to another folder. If you want to simply rename the file's name, use Rename instead, otherwise,
// if you want to keep the name, but move the folder, use MoveTo instead.
func (file *File) Move(dest string) error {
	if err := mkparent(dest); err != nil {
		return err
	}
	return os.Rename(file.path, dest)
}

// Rename renames the file while keeping the source folder, this is useful when you simply want to rename the
// name of the file, change the extension or something similar.
//
// If you want to move the file into an entirely new folder, use Move instead.
// You can also use MoveTo if you want to move to another folder, but still keep the name.
func (file *File) Rename(name string) error {
	dir := filepath.Dir(file.path)
	return os.Rename(file.path, filepath.Join(dir, name))
}

// MoveTo moves the file to another folder while keeping its name, this is useful when you just want to change
// the folder of the file.
//
// If you want to move the file into an entirely new folder, use Move instead.
// You can also use Rename if you want to rename the file's name.
func (file *File) MoveTo(dir string) error {
	base := filepath.Base(file.path)
	dest := filepath.Join(dir, base)
	if err := mkparent(dest); err != nil {
		return err
	}
	return os.Rename(file.path, dest)
}
