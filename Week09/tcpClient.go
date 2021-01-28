package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:1000")
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}

	defer client.Close()
	reqReader := bufio.NewReader(os.Stdin)
	for {
		req, err := reqReader.ReadString('\n')
		if err != nil {
			log.Fatalf("read message from stdin error: %v\n", err)
		}
		client.Write([]byte(req))
		resp, err := ioutil.ReadAll(client)

		fmt.Printf("Recieve date from server: %v\n", resp)
	}
}
