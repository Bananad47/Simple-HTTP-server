package main

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/pkg/constants"
	"github.com/codecrafters-io/http-server-starter-go/pkg/request"
	"github.com/codecrafters-io/http-server-starter-go/pkg/response"
	"github.com/codecrafters-io/http-server-starter-go/pkg/router"
	"log"
	"net"
	"net/http"
	"regexp"
)

func main() {

	rt := router.NewRouter()

	rt.GET("/echo/([\\w/]+)", func(r *request.Request) *response.Response {
		resp := response.CreateResponse(r.Path, http.StatusOK)
		t, _ := regexp.Compile("/echo/([\\w/]+)")
		m := t.FindStringSubmatch(r.Path)
		resp.AddContent(m[1])
		return resp
	})

	rt.GET("/user-agent", func(r *request.Request) *response.Response {
		resp := response.CreateResponse(r.Path, http.StatusOK)
		resp.AddContent(r.Headers["User-Agent"])
		return resp
	})

	rt.GET("[/]$", func(r *request.Request) *response.Response {
		resp := response.CreateResponse(r.Path, http.StatusOK)
		resp.AddContent(constants.StatusMessages[http.StatusOK])
		return resp
	})

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
	defer func() {
		if err := connection.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	err = rt.ProcessConnection(connection)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Server closed")
}
