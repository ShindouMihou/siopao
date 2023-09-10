package streaming

import (
	"bufio"
	"os"
)

type Reader struct {
	file  *os.File
	cache *[][]byte
}

func NewReader(file *os.File) *Reader {
	return &Reader{file: file}
}

type LineReader func(line []byte)

func (reader *Reader) EachLine(fn LineReader) error {
	return reader.eachline(false, fn)
}

func (reader *Reader) EachImmutableLine(fn LineReader) error {
	return reader.eachline(true, fn)
}

func (reader *Reader) File() *os.File {
	return reader.file
}

func (reader *Reader) Empty() {
	reader.cache = nil
}

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
