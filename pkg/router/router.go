package router

import (
	"github.com/codecrafters-io/http-server-starter-go/pkg/constants"
	"github.com/codecrafters-io/http-server-starter-go/pkg/request"
	"github.com/codecrafters-io/http-server-starter-go/pkg/response"
	"net"
	"net/http"
	"regexp"
)

type Handler func(*request.Request) *response.Response

type Pair[T1, T2 any] struct {
	First  T1
	Second T2
}

type Router struct {
	paths map[constants.Method][]Pair[*regexp.Regexp, Handler]
}

func NewRouter() *Router {
	return &Router{
		paths: map[constants.Method][]Pair[*regexp.Regexp, Handler]{},
	}
}

func (r *Router) GET(path string, hd Handler) error {
	return r._AddPath(constants.GET, path, hd)
}
func (r *Router) POST(path string, hd Handler) error {
	return r._AddPath(constants.POST, path, hd)
}
func (r *Router) DELETE(path string, hd Handler) error {
	return r._AddPath(constants.DELETE, path, hd)
}
func (r *Router) HEAD(path string, hd Handler) error {
	return r._AddPath(constants.HEAD, path, hd)
}

func (r *Router) _AddPath(method constants.Method, path string, hd Handler) error {
	re, err := regexp.Compile(path)
	if err != nil {
		return err
	}
	r.paths[method] = append(r.paths[method], Pair[*regexp.Regexp, Handler]{re, hd})
	return nil
}

func (r *Router) ProcessRequest(req *request.Request) *response.Response {
	for _, path := range r.paths[req.Method] {
		if path.First.Match([]byte(req.Path)) {
			return path.Second(req)
		}
	}
	return NotFoundHandler(req)
}

func NotFoundHandler(r *request.Request) *response.Response {
	resp := response.CreateResponse(r.Path, http.StatusNotFound)
	resp.AddContent(constants.StatusMessages[http.StatusNotFound])
	return resp
}

func (r *Router) ProcessConnection(con net.Conn) error {
	req, err := request.ParseRequest(con)
	if err != nil {
		return err
	}
	resp := r.ProcessRequest(req)
	err = resp.WriteResponse(con)
	return err
}

func (r *Router) Launch(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go func() {
			r.ProcessConnection(conn)
			conn.Close()
		}()
	}
}
