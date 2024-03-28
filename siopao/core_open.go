package siopao

import (
	"os"
	"path/filepath"
	"strings"
)

func (file *File) openRead() (*os.File, error) {
	f, err := os.Open(file.path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (file *File) openWrite(trunc bool) (*os.File, error) {
	if err := file.MkdirParent(); err != nil {
		return nil, err
	}

	var f *os.File
	var err error

	if trunc {
		f, err = os.Create(file.path)
	} else {
		f, err = os.OpenFile(file.path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	}

	if err != nil {
		return nil, err
	}

	if trunc {
		if err := file.clear(f); err != nil {
			return nil, err
		}
	}
	return f, nil
}

func (file *File) clear(f *os.File) error {
	if err := f.Truncate(0); err != nil {
		return err
	}
	if _, err := f.Seek(0, 0); err != nil {
		return err
	}
	return nil
}

// MkdirParent creates the parent folders of the path, this also includes the current
// path if it is a directory already.
func (file *File) MkdirParent() error {
	return mkparent(file.path)
}

func mkparent(path string) error {
	if strings.Contains(path, "\\") || strings.Contains(path, "/") {
		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
