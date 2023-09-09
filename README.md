# go-simple-files

*Simplified file operations for Golang.* Go-Simple-Files is a library built from the experimental joke project 
[`go-bun`](https://github.com/ShindouMihou/go-bun) in order to bring some simpler abstractions over common, simple 
file operations. As its roots, go-simple-files was inspired [`bun.sh`](https://bun.sh)'s incredible developer-experience 
for file operations.

## demo
```go
package main

import (
	"fmt"
	"github.com/ShindouMihou/go-simple-files/files"
	"log"
)

type Hello struct {
	World string `json:"world"`
}

func main() {
	file := files.Of("test.json")
	if err := file.Overwrite(Hello{World: "hello world"}); err != nil {
		log.Fatalln(err)
	}
	var hello Hello
	if err := file.Json(&hello); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(hello.World)
}

```

## file io
go-simple-files supports the following file methods:
- [x] `File.Write(any)`: writes to file, appends if it exists. anything other than string and byte array is translated to json.
- [x] `File.Overwrite(any)`: overwrites the file with the new contents, similar to the above and translates anything else to json.
- [x] `File.Text`: reads the file contents and into a string.
- [x] `File.Bytes`: reads the file contents and into a byte array.
- [x] `File.Reader`: returns a [`Reader`](#reader) of the file.
- [x] `File.TextReader`: returns a [`TextReader`](#textreader) of the file.
- [x] `File.Writer(overwrite)`: returns a [`Writer`](#write-streams) of the file, creates the file if needed.
- [x] `File.WriterSize(overwrite, buffer_size)`: returns a [`Writer`](#write-streams) with a specified buffer size of the file, creates the file if needed.


## read streams

go-simple-files also has simplified streaming that helps with stream reading.

> **Warning**
> 
> All methods in the readers will close the File, which means that these are not reusable. Although, you can create 
> the reader again through the same way using the [`File`](#file-io) interface as the `Reader` creation methods will open the file 
> once again.

### typedreader
a streaming reader that is intended to be used for json arrays with each line being a one-line json object.
can be created using `streams.NewTypedReader[T any](reader)`.
- [x] `Lines`: reads each line and transform it into the type before adding them to an array.

### reader
the base streaming reader that handles with bytes.
- [x] `Lines`: reads each line and creates an array of `[]bytes`. this also caches the array into the reader, you can empty it using `empty`
- [x] `Count`: counts all the lines in the file, this calls `Lines` and counts the cache if there is one already.
- [x] `EachLine`: reads each line and performs an action upon that line, **the line's byte array will be overridden on each next line**
- [x] `EachImmutableLine`: reads each line and performs an action upon that line, slower than the prior method, but the line's value is never overridden on each next line.
- [x] `Empty`: dereferences the cache if there is any.

### textreader
a simple streaming reader that handles with strings. it wraps around [`reader`](#reader).
- [x] `Lines`: reads each line and creates an array of string. this also caches the array into the reader, you can empty it using `empty`
- [x] `Count`: counts all the lines in the file, this calls `Lines` and counts the cache if there is one already.
- [x] `EachLine`: reads each line and performs an action upon that line.
- [x] `Empty`: dereferences the cache if there is any.

## write streams

go-simple-files also has simplified streaming that helps with streamwriting.

> **Warning**
> 
> It is your responsibility to close the buffer when it comes to writing, when possible, use 
> `Close` to flush the buffer and close the file to prevent anything crazy happening.

- `Writer`: the all-around streaming writer, defaults to json for anything other than bytes and string.
  - [x] `AlwaysAppendNewLine`: sets the writer to always append a new line on each new write.
  - [x] `Write(any)`: similar to the [`File.write`](#file-io) but pushes to the buffer, this translates anything other than bytes and string to json.
  - [x] `Flush`: flushes the buffer.
  - [x] `End`: flushes the buffer and closes the file. similar to bun's `FileSink.end`.
  - [x] `Close`: closes the file, but does not flush the buffer, this is risky.
  - [x] `Reset`: whatever the heck `bufio.Writer.Reset` does.
