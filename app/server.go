package main

import (
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
	_, err = connection.Write([]byte("HTTP/1.1 200 OK\\r\\n\\r\\n"))
	if err != nil {
		log.Fatalln("Error writing answer: ", err.Error())
	}
}
