package request

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Request struct {
	Headers     map[string]string
	Method      string
	Path        string
	HttpVersion string
}

var IncorrectRequestError = errors.New("Incorrect request")

func ParseRequest(rawreq io.Reader) (*Request, error) {
	rd := bufio.NewReader(rawreq)
	lines := []string{}
	var prev byte = 0
	line := []byte{}
	var err error
	for {
		cur, err := rd.ReadByte()
		fmt.Println(cur, err, err == io.EOF, lines)
		if err != nil {
			lines = append(lines, string(line))
			break
		}
		fmt.Println(prev, cur, '\r', '\n')
		if prev == '\r' && cur == '\n' {
			lines = append(lines, string(line))
			line = []byte{}
		} else if cur != '\n' && cur != '\r' {
			line = append(line, cur)
		}
		prev = cur
	}
	fmt.Println("OKKKKK", lines)
	if err != io.EOF {
		return nil, err
	}
	if len(lines) < 1 {
		return nil, IncorrectRequestError
	}
	head := strings.Split(lines[0], " ")
	if len(head) != 3 {
		return nil, IncorrectRequestError
	}
	req := &Request{}
	req.Method = head[0]
	req.Path = head[1]
	req.HttpVersion = head[2]
	for _, headerline := range lines {
		temp := strings.SplitN(headerline, ":", 2)
		if len(temp) != 2 {
			return nil, IncorrectRequestError
		}
		req.Headers[temp[0]] = temp[1]
	}
	return req, nil
}
