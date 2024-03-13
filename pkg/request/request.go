package request

import (
	"bufio"
	"errors"
	"github.com/codecrafters-io/http-server-starter-go/pkg/constants"
	"io"
	"strings"
)

type Request struct {
	Headers     map[string]string
	Method      constants.Method
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
	cntn := 0
	cntr := 0
	for {
		cur, err := rd.ReadByte()
		if err != nil || (cntr == 2 && cntn == 2) {
			break
		}
		if prev == '\r' && cur == '\n' && len(line) > 0 {
			lines = append(lines, string(line))
			line = []byte{}
		} else if cur != '\n' && cur != '\r' {
			line = append(line, cur)
			cntn = 0
			cntr = 0
		} else {
			cntr++
			cntn++
		}
		prev = cur
	}
	if err != nil && err != io.EOF {
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
	req.Headers = map[string]string{}
	req.Method = constants.ParseMethod(head[0])
	req.Path = head[1]
	req.HttpVersion = head[2]
	for _, headerline := range lines[1:] {
		temp := strings.SplitN(headerline, ":", 2)
		if len(temp) != 2 {
			return nil, IncorrectRequestError
		}
		req.Headers[temp[0]] = strings.Trim(temp[1], " ")
	}
	return req, nil
}
