package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filename := "words.txt"

	fileContexts, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	words := strings.Fields(string(fileContexts))

	fmt.Println("Found", len(words), "words")
}
