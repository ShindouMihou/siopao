package streaming

import (
	"bufio"
	"bytes"
	"github.com/ShindouMihou/siopao/paopao"
)

type TypedReader[T any] struct {
	reader    *Reader
	unmarshal paopao.Unmarshaler
}

// NewTypedReader creates a TypedReader from a Reader instance, this uses the paopao.Unmarshal as its unmarshaler,
// to change the unmarshaler, use WithUnmarshaler.
func NewTypedReader[T any](reader *Reader) *TypedReader[T] {
	return &TypedReader[T]{
		reader:    reader,
		unmarshal: paopao.Unmarshal,
	}
}

// WithUnmarshaler changes the unmarshaler of the typed reader, allowing you to change it to whichever other
// format that you prefer, or using an even faster unmarshaler.
func (reader *TypedReader[T]) WithUnmarshaler(unmarshaler paopao.Unmarshaler) {
	reader.unmarshal = unmarshaler
}

type TypedLineReader[T any] func(t *T)

// Lines will read each line and unmarshals it into the given type. Note that this will exhaust the underlying
// io.Reader which means that the reader becomes unusable after using this method.
func (reader *TypedReader[T]) Lines() ([]T, error) {
	var arr []T
	if err := reader.EachLine(func(t *T) {
		arr = append(arr, *t)
	}); err != nil {
		return nil, err
	}
	return arr, nil
}

// EachLine reads each line and unmarshals it into the given type before performing the given function. Note that this
// will exhaust the underlying io.Reader which means that the reader becomes unusable after using this method.
func (reader *TypedReader[T]) EachLine(fn TypedLineReader[T]) error {
	defer reader.reader.Close()
	scanner := bufio.NewScanner(reader.reader.file)
	for scanner.Scan() {
		line := scanner.Bytes()

		if len(line) < 2 {
			continue
		}

		if bytes.EqualFold(line, []byte{'['}) || bytes.EqualFold(line, []byte{']'}) {
			continue
		}

		end := len(line)
		if bytes.HasSuffix(line, []byte{','}) {
			end = end - 1
		}

		var t T
		if err := reader.unmarshal(line[:end], &t); err != nil {
			return err
		}

		fn(&t)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
