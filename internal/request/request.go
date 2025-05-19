package request

import (
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	reqStr := string(b)
	lines := strings.Split(reqStr, "\r\n")
	reqLine, err := parseRequestLine(lines[0])
	if err != nil {
		return nil, err
	}

	req := Request{
		RequestLine: reqLine,
	}
	return &req, nil
}

func parseRequestLine(line string) (RequestLine, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return RequestLine{}, fmt.Errorf("invalid number of parts")
	}
	method := parts[0]
	target := parts[1]
	version := parts[2]

	if method != strings.ToUpper(method) {
		return RequestLine{}, fmt.Errorf("not uppercase method")
	}

	parsedVersion := strings.TrimPrefix(version, "HTTP/")
	if parsedVersion != "1.1" {
		return RequestLine{}, fmt.Errorf("invalid version")
	}

	reqLine := RequestLine{
		Method:        method,
		RequestTarget: target,
		HttpVersion:   parsedVersion,
	}
	return reqLine, nil
}
