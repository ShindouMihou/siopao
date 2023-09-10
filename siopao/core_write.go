package siopao

import (
	"bufio"
	buffer2 "github.com/ShindouMihou/siopao/internal/buffer"
	"github.com/ShindouMihou/siopao/paopao"
	"io"
)

func (file *File) wrt(trunc bool, bytes []byte) error {
	if _, err := write(file, trunc, func() (*any, error) {
		if _, err := file.file.Write(bytes); err != nil {
			return nil, err
		}
		return nil, nil
	}); err != nil {
		return err
	}
	return nil
}

func (file *File) wrtbuffer(trunc bool, buffer io.Reader) error {
	if _, err := write(file, trunc, func() (*any, error) {
		if err := buffer2.Read(buffer, 4_096, func(bytes []byte) error {
			if _, err := file.file.Write(bytes); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return nil, err
		}
		return nil, nil
	}); err != nil {
		return err
	}
	return nil
}

func (file *File) wrtjson(trunc bool, t interface{}) error {
	return file.wrtmarshal(paopao.Marshal, trunc, t)
}

func (file *File) wrtmarshal(marshal paopao.Marshaller, trunc bool, t interface{}) error {
	bytes, err := marshal(t)
	if err != nil {
		return err
	}
	return file.wrt(trunc, bytes)
}

func (file *File) wrtany(trunc bool, t any) error {
	switch t.(type) {
	case string:
		return file.wrt(trunc, []byte(t.(string)))
	case []byte:
		return file.wrt(trunc, t.([]byte))
	case *bufio.Reader:
		return file.wrtbuffer(trunc, t.(*bufio.Reader))
	case bufio.Reader:
		buffer := t.(bufio.Reader)
		return file.wrtbuffer(trunc, &buffer)
	case io.Reader:
		return file.wrtbuffer(trunc, t.(io.Reader))
	default:
		return file.wrtjson(true, t)
	}
}
