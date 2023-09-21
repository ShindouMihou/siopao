package siopao

type File struct {
	path string
}

// Open opens up a new interface with the given file.
//
// siopao.Open will lazily open the file, which means that the file is opened as many times as it is needed and is
// closed immediately after use, unless it is needed by streaming. This prevents unnecessary resources from being
// leaked.
func Open(path string) *File {
	return &File{
		path: path,
	}
}
