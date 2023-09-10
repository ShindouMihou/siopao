package siopao

import (
	"github.com/ShindouMihou/siopao/streaming"
	"os"
	"testing"
)

func TestFile_Overwrite(t *testing.T) {
	file := Open(".tests/write-01.txt")
	if err := file.Overwrite("hello world"); err != nil {
		t.Fatal("failed to write to test text file: ", err)
	}
}

func TestFile_Text(t *testing.T) {
	file := Open(".tests/write-01.txt")

	text, err := file.Text()
	if err != nil {
		t.Fatal("failed to read to test text file: ", err)
	}
	if text != "hello world" {
		t.Fatal("test file does not match expected result, got '", text, "' instead of 'hello world'")
	}
}

type Hello struct {
	World string `json:"world"`
}

func TestFile_Overwrite2(t *testing.T) {
	file := Open(".tests/write-02.json")
	if err := file.Overwrite(Hello{"hello world"}); err != nil {
		t.Fatal("failed to write to test json file: ", err)
	}
}

func TestFile_Json(t *testing.T) {
	file := Open(".tests/write-02.json")

	var hello Hello
	if err := file.Json(&hello); err != nil {
		t.Fatal("failed to read to test json file: ", err)
	}
	if hello.World != "hello world" {
		t.Fatal("test file does not match expected result, got '", hello.World, "' instead of 'hello world'")
	}
}

func BenchmarkFile_Write(b *testing.B) {
	file := Open(".tests/bench-01.txt")
	for i := 0; i < b.N; i++ {
		if err := file.Write("hello world\n"); err != nil {
			b.Fatal("failed to write to test text file: ", err)
		}
	}

	b.Cleanup(func() {
		if err := os.Remove(".tests/bench-01.txt"); err != nil {
			b.Fatal("failed to clean up benchmark file.")
		}
	})
}

func BenchmarkFile_Writer(b *testing.B) {
	file := Open(".tests/bench-01.txt")
	writer, err := file.Writer(true)
	if err != nil {
		b.Fatal("failed to clean up benchmark file.")
	}
	b.ResetTimer()

	defer func(writer *streaming.Writer) {
		err := writer.End()
		if err != nil {
			b.Fatal("failed to close writer")
		}
	}(writer)

	for i := 0; i < b.N; i++ {
		if err := writer.Write("hello world\n"); err != nil {
			b.Fatal("failed to write to test text file: ", err)
		}
	}

	b.Cleanup(func() {
		if err := os.Remove(".tests/bench-01.txt"); err != nil {
			b.Fatal("failed to clean up benchmark file.")
		}
	})
}

func BenchmarkFile_Reader(b *testing.B) {
	file := Open(".tests/bench-01.txt")
	writer, err := file.Writer(true)
	if err != nil {
		b.Fatal("failed to clean up benchmark file.")
	}
	for i := 0; i < b.N; i++ {
		if err := writer.Write("hello world\n"); err != nil {
			b.Fatal("failed to write to test text file: ", err)
		}
	}
	if err := writer.End(); err != nil {
		b.Fatal("failed to close writer")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader, err := file.Reader()
		if err != nil {
			b.Fatal("failed to open reader")
		}
		if err := reader.EachLine(func(line []byte) {}); err != nil {
			b.Fatal("failed to read test file: ", err)
		}
	}

	b.Cleanup(func() {
		if err := os.Remove(".tests/bench-01.txt"); err != nil {
			b.Fatal("failed to clean up benchmark file.")
		}
	})
}

func BenchmarkFile_Bytes2(b *testing.B) {
	file := Open(".tests/bench-01.txt")
	writer, err := file.Writer(true)
	if err != nil {
		b.Fatal("failed to clean up benchmark file.")
	}
	for i := 0; i < b.N; i++ {
		if err := writer.Write("hello world\n"); err != nil {
			b.Fatal("failed to write to test text file: ", err)
		}
	}
	if err := writer.End(); err != nil {
		b.Fatal("failed to close writer")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := file.Bytes(); err != nil {
			b.Fatal("failed to read test file: ", err)
		}
	}

	b.Cleanup(func() {
		if err := os.Remove(".tests/bench-01.txt"); err != nil {
			b.Fatal("failed to clean up benchmark file.")
		}
	})
}

func BenchmarkFile_Overwrite(b *testing.B) {
	file := Open(".tests/write-01.txt")
	for i := 0; i < b.N; i++ {
		if err := file.Overwrite("hello world"); err != nil {
			b.Fatal("failed to write to test text file: ", err)
		}
	}
}

func BenchmarkFile_Text(b *testing.B) {
	file := Open(".tests/write-01.txt")
	for i := 0; i < b.N; i++ {
		if _, err := file.Text(); err != nil {
			b.Fatal("failed to read to test text file: ", err)
		}
	}
}

func BenchmarkFile_Bytes(b *testing.B) {
	file := Open(".tests/write-01.txt")
	for i := 0; i < b.N; i++ {
		if _, err := file.Bytes(); err != nil {
			b.Fatal("failed to read to test text file: ", err)
		}
	}
}
