package streaming

import (
	"github.com/ShindouMihou/siopao/internal/buffer"
	"io"
)

func (writer *Writer) wrtbuffer(buf io.Reader) error {
	return buffer.Read(buf, 4_096, func(bytes []byte) error {
		if _, err := writer.file.Write(bytes); err != nil {
			return err
		}
		return nil
	})
}
