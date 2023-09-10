package siopao

import "github.com/ShindouMihou/siopao/streaming"

// Reader opens a stream to the file, allowing you to handle big file streaming easily.
//
// This causes the file to be opened, therefore, we recommend using the returned streaming.Reader immediately
// to prevent unnecessary leaking of resources.
func (file *File) Reader() (*streaming.Reader, error) {
	err := file.openRead()
	if err != nil {
		return nil, err
	}
	return streaming.NewReader(file.file), nil
}

// TextReader opens a string stream to the file, this is an abstraction over the streaming.Reader to handle
// text (string) instead of bytes.
//
// This causes the file to be opened, therefore, we recommend using the returned streaming.Reader immediately
// to prevent unnecessary leaking of resources.
func (file *File) TextReader() (*streaming.TextReader, error) {
	reader, err := file.Reader()
	if err != nil {
		return nil, err
	}
	return reader.AsTextReader(), nil
}

// WriterSize opens a write stream, allowing easier stream writing to the file. Unlike Writer, this opens a writing stream
// with the provided buffer size, although it's more recommended to use Writer unless you need to use a higher buffer size.
//
// This causes the file to be opened, it is up to you to close the streaming.Writer using the methods provided.
// We recommend using streaming.Writer's End method to close the writer as it flushes and closes the file.
func (file *File) WriterSize(overwrite bool, size int) (*streaming.Writer, error) {
	err := file.openWrite(overwrite)
	if err != nil {
		return nil, err
	}
	return streaming.NewWriterSize(file.file, size), nil
}

// Writer opens a write stream, allowing easier stream writing to the file. Unlike WriterSize, this opens a writing stream
// with a buffer size of 4,096 bytes, if you need to customize the buffer size, then use WriterSize instead.
//
// This causes the file to be opened, it is up to you to close the streaming.Writer using the methods provided.
// We recommend using streaming.Writer's End method to close the writer as it flushes and closes the file.
func (file *File) Writer(overwrite bool) (*streaming.Writer, error) {
	err := file.openWrite(overwrite)
	if err != nil {
		return nil, err
	}
	return streaming.NewWriter(file.file), nil
}
