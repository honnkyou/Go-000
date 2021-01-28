package main

import (
	"log"
	"net"
)

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:1000")
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}

	defer client.Close()

}
