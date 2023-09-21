package siopao

import (
	"errors"
	"github.com/ShindouMihou/siopao/paopao"
	"io"
	"os"
	"reflect"
)

// Text reads the file directly as a string, this is not recommended to use when handling big
// files, we recommend using TextReader to stream big files instead.
func (file *File) Text() (string, error) {
	bytes, err := file.Bytes()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Bytes reads the file directly as a byte array, this is not recommend to use when handling big
// files, we recommend using Reader to stream big files instead.
func (file *File) Bytes() ([]byte, error) {
	bytes, err := read(file, func(f *os.File) (*[]byte, error) {
		bytes, err := io.ReadAll(f)
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

// Unmarshal unmarshals the given contents of the file with the given unmarshaler.
func (file *File) Unmarshal(unmarshal paopao.Unmarshaler, t interface{}) error {
	if reflect.TypeOf(t).Kind() != reflect.Pointer {
		return errors.New("non-pointer kind for value")
	}

	if _, err := read[any](file, func(f *os.File) (*any, error) {
		bytes, err := io.ReadAll(f)
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

// Json unmarshals the contents of the file into Json using the paopao.Unmarshal.
func (file *File) Json(t interface{}) error {
	return file.Unmarshal(paopao.Unmarshal, t)
}
