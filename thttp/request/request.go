package request

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"time"
)

//const maxCapacity = 1024 * 1024 // your required line length
type Request struct {
	Method  string
	URI     string
	Headers map[string]string
	Payload string
}

type RequestData struct {
	User map[string]string
	Now  time.Time
}

func NewRequest(filePath string) (*Request, error) {
	newRequest := Request{
		Headers: make(map[string]string),
	}
	requestFile, err := readRequestFile(filePath)
	if err != nil {
		return nil, err
	}

	if requestFile != nil {
		defer requestFile.Close()
	}

	request, err := buildRequestTemplate(requestFile)
	if err != nil {
		return nil, err
	}
	err = parseRequest(request, &newRequest)
	if err != nil {
		return nil, err
	}

	return &newRequest, nil
}

func parseRequest(requestString string, request *Request) error {
	templateReader := strings.NewReader(requestString)

	//buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(templateReader)
	//scanner.Buffer(buf, maxCapacity)
	method, uri, err := readFirst(scanner)
	if err != nil {
		return err
	}
	err = readHeaders(scanner, request)
	if err != nil {
		return err
	}

	request.URI = uri
	request.Method = method
	return nil
}
func readHeaders(scanner *bufio.Scanner, request *Request) error {
	for scanner.Scan() {
		line := strings.SplitN(scanner.Text(), ":", 2)
		if len(line) == 1 {
			break
		}
		request.Headers[line[0]] = line[1]
	}
	return nil
}

func readFirst(scanner *bufio.Scanner) (string, string, error) {
	if scanner.Scan() {
		firstLine := strings.Split(scanner.Text(), " ")
		if len(firstLine) != 2 {
			return "", "", errors.New("method or URI missing on the first line")
		} else {
			return firstLine[0], firstLine[1], nil
		}
	} else {
		return "", "", errors.New("unable to read first line")
	}
}

func buildRequestTemplate(requestTemplate *os.File) (string, error) {
	request, err := ioutil.ReadAll(requestTemplate)
	if err != nil {
		return "", err
	}
	tmpl, err := template.New("HTTPRequest").Parse(string(request))
	if err != nil {
		return "", err
	}

	templateResult := bytes.NewBufferString("")
	templateData := RequestData{
		Now: time.Now(),
	}
	err = tmpl.Execute(templateResult, templateData)
	if err != nil {
		return "", err
	}

	return templateResult.String(), nil
}

func readRequestFile(filePath string) (*os.File, error) {
	requestFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return requestFile, nil
}
