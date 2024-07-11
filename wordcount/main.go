package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("need to provide filename!")
		os.Exit(1)
	}

	// Loop over all the arguments from index 1 onwards.
	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)

		// If there's an error, print it and skip to the next file.
		if err != nil {
			log.Printf("%s: %v", filename, err)
			continue
		}

		// Create a scanner to read from the file.
		scanner := bufio.NewScanner(file)

		// Change the split function to split on words instead of lines.
		scanner.Split(bufio.ScanWords)

		// Counter to track the running total.
		var wordCount int

		// Add 1 to the count for each word
		for scanner.Scan() {
			wordCount++
		}

		if scanner.Err() != nil {
			log.Printf("scan error %s: %v", filename, scanner.Err())
			file.Close()
			continue
		}

		file.Close()
		fmt.Printf("%s: %d\n", filename, wordCount)
	}
}
