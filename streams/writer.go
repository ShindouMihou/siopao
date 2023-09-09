package streams

import (
	"bufio"
	"encoding/json"
	"os"
)

type Writer struct {
	file          *os.File
	writer        *bufio.Writer
	appendNewLine bool
}

func NewWriter(file *os.File) *Writer {
	return NewWriterSize(file, 4096)
}

func NewWriterSize(file *os.File, size int) *Writer {
	return &Writer{
		file:          file,
		writer:        bufio.NewWriterSize(file, size),
		appendNewLine: false,
	}
}

func (writer *Writer) AlwaysAppendNewLine() *Writer {
	writer.appendNewLine = true
	return writer
}

func (writer *Writer) Write(t any) error {
	switch t.(type) {
	case string:
		return writer.write([]byte(t.(string)))
	case []byte:
		return writer.write(t.([]byte))
	default:
		bytes, err := json.Marshal(t)
		if err != nil {
			return err
		}
		return writer.write(bytes)
	}
}

func (writer *Writer) Flush() error {
	return writer.writer.Flush()
}

func (writer *Writer) Close() {
	_ = writer.file.Close()
}

func (writer *Writer) End() error {
	defer writer.Close()
	return writer.Flush()
}

func (writer *Writer) Reset() {
	writer.writer.Reset(writer.file)
}

func (writer *Writer) write(t []byte) error {
	if _, err := writer.writer.Write(t); err != nil {
		return err
	}
	if writer.appendNewLine {
		return writer.write([]byte{'\n'})
	}
	return nil
}
