package siopao

import "os"

// Delete deletes the file, or an empty directory. If you need to delete a directory that isn't empty, then use
// DeleteRecursively instead.
func (file *File) Delete() error {
	return os.Remove(file.path)
}

// DeleteRecursively deletes the file or directory and its children, if there are any, simply a short-hand of os.RemoveAll.
func (file *File) DeleteRecursively() error {
	return os.RemoveAll(file.path)
}
