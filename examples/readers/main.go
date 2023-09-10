package main

import (
	"fmt"
	"github.com/ShindouMihou/siopao/siopao"
	"log"
)

func main() {
	file := siopao.Open("examples/readers/big_test.txt")
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
