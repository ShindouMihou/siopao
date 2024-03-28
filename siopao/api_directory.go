package siopao

import (
	"fmt"
	"os"
	"path/filepath"
)

// IsDir checks whether the file is a directory, when the File comes from a Recurse call, or another call previously
// used `IsDir` then that value will be cached. To not use the cached value, use the UncachedIsDir method instead.
func (file *File) IsDir() (bool, error) {
	if file.isDir != -1 {
		return file.isDir == 1, nil
	}

	_, err := file.UncachedIsDir()
	if err != nil {
		return false, err
	}

	return file.isDir == 1, nil
}

// UncachedIsDir checks whether the file is a directory without passing through the cache. This is recommended
// to use when the file is frequently changing between a directory, or a file.
func (file *File) UncachedIsDir() (bool, error) {
	fileInfo, err := os.Stat(file.path)
	if err != nil {
		return false, err
	}

	if fileInfo.IsDir() {
		file.isDir = 1
	} else {
		file.isDir = 0
	}

	return file.isDir == 1, nil
}

// Recurse recurses through the directory if it's a directory. You can specify whether to recurse
// deep into the directory by setting the nested option to true.
func (file *File) Recurse(nested bool, fn func(file *File)) error {
	isDirectory, err := file.IsDir()
	if err != nil {
		return err
	}
	if !isDirectory {
		return fmt.Errorf("%s is not a directory", file.path)
	}
	return file.recurse(nested, fn)
}

// MkdirParent creates the parent folders of the path, this also includes the current
// path if it is a directory already.
func (file *File) MkdirParent() error {
	return mkparent(file.path)
}

func (file *File) recurse(nested bool, fn func(file *File)) error {
	files, err := os.ReadDir(file.path)
	if err != nil {
		return err
	}
	for _, f := range files {
		path := filepath.Join(file.path, f.Name())
		child := Open(path)

		if f.IsDir() {
			child.isDir = 1
		} else {
			child.isDir = 0
		}

		fn(child)
		if f.IsDir() && nested {
			err := child.recurse(nested, fn)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
