package streaming

type TextReader struct {
	reader *Reader
	cache  *[]string
}

type TextLineReader func(line string)

func (reader *Reader) AsTextReader() *TextReader {
	return &TextReader{reader: reader}
}

func (reader *TextReader) Empty() {
	reader.cache = nil
}

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

func (reader *TextReader) EachLine(fn TextLineReader) error {
	return reader.reader.EachLine(func(line []byte) {
		fn(string(line))
	})
}
