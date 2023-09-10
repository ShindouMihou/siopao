package streaming

import (
	"bufio"
	"github.com/ShindouMihou/siopao/paopao"
	"os"
)

type Writer struct {
	file          *os.File
	writer        *bufio.Writer
	appendNewLine bool
}

// NewWriter creates a new Writer from the given os.File, this creates a Writer with a buffer size of
// 4,096 bytes. If you want to create one with a different buffer size, use the NewWriterSize method instead.
func NewWriter(file *os.File) *Writer {
	return NewWriterSize(file, 4096)
}

// NewWriterSize creates a new Writer from the given os.File with a given buffer size.
func NewWriterSize(file *os.File, size int) *Writer {
	return &Writer{
		file:          file,
		writer:        bufio.NewWriterSize(file, size),
		appendNewLine: false,
	}
}

// AlwaysAppendNewLine will set the Writer to always append a new line for each write.
func (writer *Writer) AlwaysAppendNewLine() *Writer {
	writer.appendNewLine = true
	return writer
}

// Write writes the content into the file, note that this does not append a new line for each write
// unless the Writer uses AlwaysAppendNewLine. This marshals anything other than  string and byte array into the
// paopao.Marshal which is Json by default.
func (writer *Writer) Write(t any) error {
	switch t.(type) {
	case string:
		return writer.write([]byte(t.(string)))
	case []byte:
		return writer.write(t.([]byte))
	default:
		bytes, err := paopao.Marshal(t)
		if err != nil {
			return err
		}
		return writer.write(bytes)
	}
}

// WriteMarshal marshals the content with the given marshaller. Note that string and byte array are also marshalled
// with the given marshaller.
func (writer *Writer) WriteMarshal(marshaller paopao.Marshaller, t any) error {
	bytes, err := marshaller(t)
	if err != nil {
		return err
	}
	return writer.write(bytes)
}

// Flush will flush all the buffered contents into the file. It is recommended to use this only when you want
// to push the contents of the file immediately, otherwise use End instead to flush and close the Writer.
func (writer *Writer) Flush() error {
	return writer.writer.Flush()
}

// Close will abruptly close the underlying io.Writer of the Writer. IT IS NOT RECOMMENDED TO USE THIS, PLEASE USE
// End INSTEAD TO FLUSH AND CLOSE THE Writer.
func (writer *Writer) Close() {
	_ = writer.file.Close()
}

// End flushes the contents into the file before closing the underlying io.Writer.
func (writer *Writer) End() error {
	defer writer.Close()
	return writer.Flush()
}

// Reset discards any unflushed buffered data, clears any error, and resets buffer to write its output to File.
// i.e. whatever the heck bufio.Writer's Reset method does.
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
