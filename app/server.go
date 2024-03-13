package main

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/pkg/constants"
	"github.com/codecrafters-io/http-server-starter-go/pkg/request"
	"github.com/codecrafters-io/http-server-starter-go/pkg/response"
	"github.com/codecrafters-io/http-server-starter-go/pkg/router"
	"log"
	"net/http"
	"regexp"
)

func main() {

	rt := router.NewRouter()

	rt.GET("/echo/(.+)", func(r *request.Request) *response.Response {
		resp := response.CreateResponse(r.Path, http.StatusOK)
		t, _ := regexp.Compile("/echo/(.+)")
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
	log.Fatal(rt.Launch("0.0.0.0:4221"))
}
