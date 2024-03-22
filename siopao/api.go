package siopao

type File struct {
	path  string
	isDir int
}

// Open opens up a new interface with the given file.
//
// siopao.Open will lazily open the file, which means that the file is opened as many times as it is needed and is
// closed immediately after use, unless it is needed by streaming. This prevents unnecessary resources from being
// leaked.
func Open(path string) *File {
	return &File{
		path:  path,
		isDir: -1,
	}
}

// Path gets the path of the file.
func (file *File) Path() string {
	return file.path
}
