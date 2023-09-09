package main

import (
	"fmt"
	"github.com/ShindouMihou/go-simple-files/files"
	"github.com/ShindouMihou/go-simple-files/streams"
	"log"
)

type Hello struct {
	World string `json:"world"`
}

func main() {
	file := files.Of("examples/typed_reader/big_test.json")
	reader, err := file.Reader()
	if err != nil {
		log.Fatalln(err)
	}
	hellos, err := streams.NewTypedReader[Hello](reader).Lines()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("A total of", len(hellos), "hellos!")
}
