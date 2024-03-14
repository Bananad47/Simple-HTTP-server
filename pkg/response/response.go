package response

import (
	"fmt"
	"io"
	"strconv"
)

type Response struct {
	Headers     map[string]string
	Path        string
	HttpVersion string
	StatusCode  int
	Content     string
}

func CreateResponse(path string, statuscode int) *Response {
	resp := &Response{}
	resp.Path = path
	resp.HttpVersion = "HTTP/1.1"
	resp.StatusCode = statuscode
	resp.Headers = map[string]string{"Content-Type": "text/plain"}
	resp.Headers["Content-Length"] = "0"
	return resp
}

func (resp *Response) AddHeader(key, val string) {
	resp.Headers[key] = val
}

func (resp *Response) AddContent(content string) {
	resp.Content = content
	resp.Headers["Content-Length"] = strconv.Itoa(len([]byte(content)))
}

func (resp *Response) WriteResponse(wr io.Writer) error {
	_, err := fmt.Fprintf(wr, "%s %d\r\n", resp.HttpVersion, resp.StatusCode)
	if err != nil {
		return err
	}
	for key, val := range resp.Headers {
		_, err = fmt.Fprintf(wr, "%s: %s\r\n", key, val)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintf(wr, "\r\n%s\r\n\r\n", resp.Content)
	return err
}
