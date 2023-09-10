package siopao

import (
	"os"
	"path/filepath"
	"strings"
)

func (file *File) openRead() error {
	f, err := os.Open(file.path)
	if err != nil {
		return err
	}
	file.file = f
	return nil
}

func (file *File) openWrite(trunc bool) error {
	if err := file.mkparent(); err != nil {
		return err
	}

	var f *os.File
	var err error

	if trunc {
		f, err = os.Create(file.path)
	} else {
		f, err = os.OpenFile(file.path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	}

	if err != nil {
		return err
	}

	file.file = f
	return nil
}

func (file *File) clear() error {
	if err := file.file.Truncate(0); err != nil {
		return err
	}
	if _, err := file.file.Seek(0, 0); err != nil {
		return err
	}
	return nil
}

// MkdirParent creates the parent folders of the path.
func (file *File) mkparent() error {
	if strings.Contains(file.path, "\\") || strings.Contains(file.path, "/") {
		if err := os.MkdirAll(filepath.Dir(file.path), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
