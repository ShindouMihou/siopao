package buffer

import (
	"errors"
	"io"
)

func Read(buf io.Reader, size uint16, fn func(bytes []byte) error) error {
	buffer := make([]byte, 0, size)
	for {
		n, err := io.ReadFull(buf, buffer[:cap(buffer)])
		buffer = buffer[:n]
		if err != nil {
			if err == io.EOF {
				break
			}
			if !errors.Is(err, io.ErrUnexpectedEOF) {
				return err
			}
		}
		if err := fn(buffer); err != nil {
			return err
		}
	}
	return nil
}
