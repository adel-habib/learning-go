package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		os.Exit(100)
	}
	addr := fmt.Sprintf("localhost:%s", arguments[1])
	fmt.Println(addr)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Failed to create socket", err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleConnection(conn)

	}
}

func handleConnection(c net.Conn) {
	addr := c.RemoteAddr()
	c.SetReadDeadline(time.Now().Add(time.Minute))
	defer c.Close()
	for {
		buffer := make([]byte, 1024)
		nbytes, err := c.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Println("Client closed connection")
			}
			log.Println(err)
			return
		}
		data := buffer[0:nbytes]
		fmt.Printf("%s->: %v\n", addr.String(), string(data))
		_, err = c.Write(data)
		if err != nil {
			fmt.Println(err)
		}
		msg := strings.TrimSpace(string(data))
		if msg == "stop" {
			return
		}
	}
}
