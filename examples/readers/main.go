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
	file := files.Of("examples/readers/big_test.txt")
	reader, err := file.TextReader()
	if err != nil {
		log.Fatalln(err)
	}
	size, err := reader.Count()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("A total of", size, "hellos!")
}
