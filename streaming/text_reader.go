package streaming

type TextReader struct {
	reader *Reader
	cache  *[]string
}

type TextLineReader func(line string)

// AsTextReader converts a Reader into a TextReader.
func (reader *Reader) AsTextReader() *TextReader {
	return &TextReader{reader: reader}
}

// Empty dereferences the cache of the reader, if any. A cache will be added when methods such as Count or Lines
// are used as it empties the underlying io.Reader, therefore, if you don't want the cache then it is recommended
// to dereference it.
func (reader *TextReader) Empty() {
	reader.cache = nil
}

// Lines returns all the lines of the file, this references the cache if there is any, otherwise loads the file
// and saves the array into the cache. To dereference the cache, simply use the Empty method. Note that this will exhaust
// the underlying io.Reader which means that the reader will no longer be usable.
func (reader *TextReader) Lines() ([]string, error) {
	if reader.cache != nil {
		return *reader.cache, nil
	}
	var lines []string
	err := reader.EachLine(func(line string) {
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
func (reader *TextReader) Count() (int, error) {
	if reader.cache == nil {
		arr, err := reader.Lines()
		if err != nil {
			return 0, err
		}
		return len(arr), nil
	}

	return len(*reader.cache), nil
}

// EachLine reads each line of the file as a string.
func (reader *TextReader) EachLine(fn TextLineReader) error {
	return reader.reader.EachLine(func(line []byte) {
		fn(string(line))
	})
}
