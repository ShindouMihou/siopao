package siopao

import "os"

func (file *File) close(f *os.File) {
	// ignore the error, it's likely that it just already called
	_ = f.Close()
}
