package main

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/pkg/constants"
	"github.com/codecrafters-io/http-server-starter-go/pkg/request"
	"github.com/codecrafters-io/http-server-starter-go/pkg/response"
	"github.com/codecrafters-io/http-server-starter-go/pkg/router"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {

	rt := router.NewRouter()

	rt.GET("/echo/(.+)", func(r *request.Request) *response.Response {
		resp := response.CreateResponse(r.Path, http.StatusOK)
		re, _ := regexp.Compile("/echo/(.+)")
		m := re.FindStringSubmatch(r.Path)
		resp.AddContent(m[1])
		return resp
	})

	rt.GET("/user-agent", func(r *request.Request) *response.Response {
		resp := response.CreateResponse(r.Path, http.StatusOK)
		resp.AddContent(r.Headers["User-Agent"])
		return resp
	})

	rt.GET("/files/(.+)", func(r *request.Request) *response.Response {

		fmt.Println(os.Getwd())

		resp := response.CreateResponse(r.Path, http.StatusOK)
		resp.AddHeader("Content-Type", "application/octet-stream")
		re, _ := regexp.Compile("/files/(.+)")
		m := re.FindStringSubmatch(r.Path)
		file, err := os.Open(m[1])
		if err != nil {
			resp.StatusCode = http.StatusNotFound
			return resp
		}
		defer file.Close()
		content, _ := io.ReadAll(file)
		resp.AddContent(string(content))
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
