package main

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/internal/request"
	"github.com/codecrafters-io/http-server-starter-go/internal/response"
	"log"
	"net"
	"regexp"
	"strings"
)

var headers = map[string]string{
	"Content-Type": "text/plain",
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	//fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	fmt.Println("Server Started")
	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Fatalln("Failed to bind to port 4221")
	}

	connection, err := listener.Accept()
	if err != nil {
		log.Fatalln("Error accepting connection: ", err.Error())
	}

	var test strings.Builder

	req, err := request.ParseRequest(connection)
	if err != nil {
		resp := response.CreateResponse(req.Path, "404 Not Found", headers)
		resp.WriteResponse(connection)
	} else {
		resp := response.CreateResponse(req.Path, "200 OK", headers)
		t, _ := regexp.Compile("/echo/([\\w]+)")
		m := t.FindStringSubmatch(req.Path)
		if len(m) != 0 {
			resp.AddContent(m[1])
		}
		resp.WriteResponse(connection)
		resp.WriteResponse(&test)
	}
	fmt.Println(test.String())
	connection.Close()
}
