package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/kashifsoofi/reqbud/internal/options"
)

func createRequest(method, url string, headers options.HeaderFlag, data string) (*http.Request, error) {
	var body io.Reader
	if len(data) > 0 {
		body = strings.NewReader(data)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for _, h := range headers {
		req.Header.Add(h.Name, h.Value)
	}

	return req, nil
}

func main() {
	opts := options.ParseArguments()

	req, err := createRequest(opts.Request, opts.URL, opts.Headers, opts.Data)
	if err != nil {
		log.Fatalf(err.Error())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(resp.StatusCode)
		log.Fatalf("http.get error %v\n", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("http.get error %v\n", err)
	}

	log.Println(string(body))

	// for name, values := range resp.Header {
	// 	log.Printf("%s: %v\n", name, values)
	// }
}
