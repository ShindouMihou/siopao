package main

import (
	"fmt"
	"github.com/ShindouMihou/siopao/siopao"
	"github.com/ShindouMihou/siopao/streaming"
	"log"
)

type Hello struct {
	World string `json:"world"`
}

func main() {
	file := siopao.Open("examples/typed_reader/big_test.json")
	reader, err := file.Reader()
	if err != nil {
		log.Fatalln(err)
	}
	hellos, err := streaming.NewTypedReader[Hello](reader).Lines()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("A total of", len(hellos), "hellos!")
}
