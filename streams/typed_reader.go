package streams

import (
	"bufio"
	"bytes"
	go_simple_files "github.com/ShindouMihou/go-simple-files/go-simple-files"
)

type TypedReader[T any] struct {
	reader *Reader
}

func NewTypedReader[T any](reader *Reader) *TypedReader[T] {
	return &TypedReader[T]{
		reader: reader,
	}
}

type TypedLineReader[T any] func(t *T)

func (reader *TypedReader[T]) Lines() ([]T, error) {
	var arr []T
	if err := reader.EachLine(func(t *T) {
		arr = append(arr, *t)
	}); err != nil {
		return nil, err
	}
	return arr, nil
}

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
		if err := go_simple_files.Unmarshal(line[:end], &t); err != nil {
			return err
		}

		fn(&t)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
