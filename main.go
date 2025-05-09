package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const BUFFER_LENGTH = 8

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	ch := getLinesChannel(file)
	for line := range ch {
		fmt.Printf("read: %s\n", line)
	}
}

func getLinesChannel(r io.ReadCloser) <-chan string {
	ch := make(chan string)

	var line string
	go func() {
		for {
			buf := make([]byte, 8)
			bytes, err := r.Read(buf)
			parts := strings.Split(string(buf[:bytes]), "\n")
			line = fmt.Sprintf("%s%s", line, parts[0])
			if len(parts) > 1 {
				ch <- line
				line = parts[1]
			}
			if err == io.EOF {
				r.Close()
				close(ch)
				return
			}
		}
	}()
	return ch
}
