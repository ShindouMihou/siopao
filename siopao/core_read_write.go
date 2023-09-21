package siopao

import "os"

func read[T any](file *File, fn func(f *os.File) (*T, error)) (*T, error) {
	f, err := file.openRead()
	if err != nil {
		return nil, err
	}
	defer file.close(f)
	return fn(f)
}

func write[T any](file *File, trunc bool, fn func(f *os.File) (*T, error)) (*T, error) {
	f, err := file.openWrite(trunc)
	if err != nil {
		return nil, err
	}
	defer file.close(f)
	return fn(f)
}
