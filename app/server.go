package main

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/internal/request"
	"log"
	"net"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	//fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Fatalln("Failed to bind to port 4221")
	}

	connection, err := listener.Accept()
	if err != nil {
		log.Fatalln("Error accepting connection: ", err.Error())
	}
	log.Println("OK in main")
	req, err := request.ParseRequest(connection)
	fmt.Println(req.Path)
	if err != nil || req.Path != "/" {
		connection.Write([]byte("HTTP/1.1 400 Not\r\n\r\n"))
	} else {
		connection.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	}
}
