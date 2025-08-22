package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for {
		data := make([]byte, 8)
		count, err := f.Read(data)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		fmt.Printf("read: %s\n", string(data[:count]))
	}
}
