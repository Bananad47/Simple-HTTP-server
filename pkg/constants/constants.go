package constants

import "net/http"

type Method int

const (
	GET Method = iota
	POST
	DELETE
	HEAD
	IDK
)

var StatusMessages = map[int]string{
	http.StatusOK:       "OK",
	http.StatusNotFound: "Not Found",
}

func ParseMethod(m string) Method {
	if m == "GET" {
		return GET
	}
	if m == "POST" {
		return POST
	}
	if m == "DELETE" {
		return DELETE
	}
	if m == "HEAD" {
		return HEAD
	}
	return IDK
}
