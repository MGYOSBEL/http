package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal("Error resolving ", err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Error Dialing ", err)
	}
	defer conn.Close()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(">")
		str, err := r.ReadString('\n')
		if err != nil {
			fmt.Printf("%s", err)
		}

		_, err = conn.Write([]byte(str))
		if err != nil {
			fmt.Printf("%s", err)
		}

	}
}
