package siopao

import (
	"errors"
	"github.com/ShindouMihou/siopao/paopao"
	"github.com/ShindouMihou/siopao/streaming"
	"io"
	"os"
	"reflect"
)

type File struct {
	path string
	file *os.File
}

func Open(path string) *File {
	return &File{
		path: path,
		file: nil,
	}
}

func (file *File) close() {
	// ignore the error, it's likely that it just already called
	_ = file.file.Close()
}

func (file *File) Text() (string, error) {
	bytes, err := file.Bytes()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (file *File) Bytes() ([]byte, error) {
	bytes, err := read(file, func() (*[]byte, error) {
		bytes, err := io.ReadAll(file.file)
		if err != nil {
			return nil, nil
		}
		return &bytes, nil
	})
	if err != nil {
		return nil, err
	}
	return *bytes, nil
}

func (file *File) Reader() (*streaming.Reader, error) {
	err := file.openRead()
	if err != nil {
		return nil, err
	}
	return streaming.NewReader(file.file), nil
}

func (file *File) TextReader() (*streaming.TextReader, error) {
	reader, err := file.Reader()
	if err != nil {
		return nil, err
	}
	return reader.AsTextReader(), nil
}

func (file *File) Unmarshal(unmarshal paopao.Unmarshaler, t interface{}) error {
	if reflect.TypeOf(t).Kind() != reflect.Pointer {
		return errors.New("non-pointer kind for value")
	}

	if _, err := read[any](file, func() (*any, error) {
		bytes, err := io.ReadAll(file.file)
		if err != nil {
			return nil, err
		}
		if err := unmarshal(bytes, t); err != nil {
			return nil, err
		}
		return nil, nil
	}); err != nil {
		return err
	}
	return nil
}

func (file *File) Json(t interface{}) error {
	return file.Unmarshal(paopao.Unmarshal, t)
}

func (file *File) WriterSize(overwrite bool, size int) (*streaming.Writer, error) {
	err := file.openWrite(overwrite)
	if err != nil {
		return nil, err
	}
	return streaming.NewWriterSize(file.file, size), nil
}

func (file *File) Writer(overwrite bool) (*streaming.Writer, error) {
	err := file.openWrite(overwrite)
	if err != nil {
		return nil, err
	}
	return streaming.NewWriter(file.file), nil
}

func (file *File) Write(t any) error {
	return file.wrtany(false, t)
}

func (file *File) Overwrite(t any) error {
	return file.wrtany(true, t)
}

func (file *File) WriteMarshal(marshal paopao.Marshaller, t any) error {
	return file.wrtmarshal(marshal, false, t)
}

func (file *File) OverwriteMarshal(marshal paopao.Marshaller, t any) error {
	return file.wrtmarshal(marshal, true, t)
}
