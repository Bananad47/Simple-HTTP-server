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
	StatusCode  string
	Content     string
}

func CreateResponse(path string, statuscode string, headers map[string]string) *Response {
	resp := &Response{}
	resp.Path = path
	resp.HttpVersion = "HTTP/1.1"
	resp.StatusCode = statuscode
	resp.Headers = headers
	resp.Headers["Content-Length"] = "0"
	return resp
}

func (resp *Response) AddContent(content string) {
	resp.Content = content
	resp.Headers["Content-Length"] = strconv.Itoa(len([]byte(content)) + 4)
}

func (resp *Response) WriteResponse(wr io.Writer) error {
	_, err := fmt.Fprintf(wr, "%s %s\r\n", resp.HttpVersion, resp.StatusCode)
	if err != nil {
		return err
	}
	for key, val := range resp.Headers {
		_, err = fmt.Fprintf(wr, "%s: %s\r\n", key, val)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintf(wr, "%s\r\n\r\n", resp.Content)
	return err
}
