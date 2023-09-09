package files

import (
	"encoding/json"
	"errors"
	"github.com/ShindouMihou/go-simple-files/streams"
	"io"
	"os"
	"reflect"
)

type File struct {
	path string
	file *os.File
}

func Of(path string) *File {
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

func (file *File) Reader() (*streams.Reader, error) {
	err := file.openRead()
	if err != nil {
		return nil, err
	}
	return streams.NewReader(file.file), nil
}

func (file *File) TextReader() (*streams.TextReader, error) {
	reader, err := file.Reader()
	if err != nil {
		return nil, err
	}
	return reader.AsTextReader(), nil
}

func (file *File) Json(t interface{}) error {
	if reflect.TypeOf(t).Kind() != reflect.Pointer {
		return errors.New("non-pointer kind for value")
	}

	if _, err := read[any](file, func() (*any, error) {
		bytes, err := io.ReadAll(file.file)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(bytes, t); err != nil {
			return nil, err
		}
		return nil, nil
	}); err != nil {
		return err
	}
	return nil
}

func (file *File) WriterSize(overwrite bool, size int) (*streams.Writer, error) {
	err := file.openWrite(overwrite)
	if err != nil {
		return nil, err
	}
	return streams.NewWriterSize(file.file, size), nil
}

func (file *File) Writer(overwrite bool) (*streams.Writer, error) {
	err := file.openWrite(overwrite)
	if err != nil {
		return nil, err
	}
	return streams.NewWriter(file.file), nil
}

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

func (file *File) wrtjson(trunc bool, t interface{}) error {
	bytes, err := json.Marshal(t)
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
	default:
		return file.wrtjson(true, t)
	}
}

func (file *File) Write(t any) error {
	return file.wrtany(false, t)
}

func (file *File) Overwrite(t any) error {
	return file.wrtany(true, t)
}
