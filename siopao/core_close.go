package siopao

func (file *File) close() {
	// ignore the error, it's likely that it just already called
	_ = file.file.Close()
}
