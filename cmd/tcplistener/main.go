package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const BUFFER_LENGTH = 8

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		fmt.Println("A connection has been accepted")
		if err != nil {
			log.Fatal(err)
		}
		ch := getLinesChannel(conn)
		for line := range ch {
			fmt.Printf("%s\n", line)
		}
		fmt.Println("The connection has been closed")
	}
}

func getLinesChannel(r io.ReadCloser) <-chan string {
	ch := make(chan string)

	var currentLine string
	go func() {
		defer r.Close()
		defer close(ch)
		for {
			buf := make([]byte, 8)
			bytes, err := r.Read(buf)
			if err != nil {
				if currentLine != "" {
					ch <- currentLine
				}
				if err == io.EOF {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				return
			}
			parts := strings.Split(string(buf[:bytes]), "\n")
			for i := 0; i < len(parts)-1; i++ {
				currentLine = fmt.Sprintf("%s%s", currentLine, parts[i])
				ch <- currentLine
				currentLine = ""
			}
			currentLine += parts[len(parts)-1]
		}
	}()
	return ch
}
