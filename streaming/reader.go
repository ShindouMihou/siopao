package streaming

import (
	"bufio"
	"os"
)

type Reader struct {
	file  *os.File
	cache *[][]byte
}

// NewReader creates a streaming reader for the given file.
func NewReader(file *os.File) *Reader {
	return &Reader{file: file}
}

type LineReader func(line []byte)

// EachLine reads each line of the file as bytes. Unlike EachImmutableLine, the byte array is reused which means
// it will be overridden each next line, therefore, it is not recommended to store the byte array elsewhere without
// copying.
func (reader *Reader) EachLine(fn LineReader) error {
	return reader.eachline(false, fn)
}

// EachImmutableLine reads each line of the file as bytes. Unlike EachLine, there is copying involved which makes this
// slower than the other, but the byte array here won't be overridden each line, allowing you to store the byte array
// elsewhere without extra copying.
func (reader *Reader) EachImmutableLine(fn LineReader) error {
	return reader.eachline(true, fn)
}

// File gets the underlying os.File of the Reader.
func (reader *Reader) File() *os.File {
	return reader.file
}

// Empty dereferences the cache of the reader, if any. A cache will be added when methods such as Count or Lines
// are used as it empties the underlying io.Reader, therefore, if you don't want the cache then it is recommended
// to dereference it.
func (reader *Reader) Empty() {
	reader.cache = nil
}

// Lines returns all the lines of the file, this references the cache if there is any, otherwise loads the file
// and saves the array into the cache. To dereference the cache, simply use the Empty method. Note that this will exhaust
// the underlying io.Reader which means that the reader will no longer be usable.
func (reader *Reader) Lines() ([][]byte, error) {
	if reader.cache != nil {
		return *reader.cache, nil
	}
	var lines [][]byte
	err := reader.EachImmutableLine(func(line []byte) {
		lines = append(lines, line)
	})
	if err != nil {
		return nil, err
	}
	reader.cache = &lines
	return lines, nil
}

// Count will count all the lines of the file. Internally, this uses Lines and will exhaust the underlying io.Reader
// which means that the reader will no longer be usable, but unless the cache is dereferenced, you can retrieve all the
// lines using the Lines method afterward.
func (reader *Reader) Count() (int, error) {
	if reader.cache == nil {
		arr, err := reader.Lines()
		if err != nil {
			return 0, err
		}
		return len(arr), nil
	}

	return len(*reader.cache), nil
}

// Close will abruptly close the underlying io.Reader, this is not needed in most cases as all the methods in the Reader
// will close the io.Reader upon completion of the action.
func (reader *Reader) Close() {
	_ = reader.file.Close()
}

func (reader *Reader) eachline(immutable bool, fn LineReader) error {
	defer reader.Close()

	scanner := bufio.NewScanner(reader.file)
	for scanner.Scan() {
		line := scanner.Bytes()
		if immutable {
			cpy := make([]byte, len(line))
			copy(cpy, line)
			line = cpy
		}
		fn(line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
