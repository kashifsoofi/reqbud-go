package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type header struct {
	name  string
	value string
}

type headerFlag []header

func (f *headerFlag) String() string {
	return ""
}

func (f *headerFlag) Set(arg string) error {
	name, value, found := strings.Cut(arg, ":")
	if !found {
		return errors.New("invalid header, must be name:value")
	}

	*f = append(*f, header{
		name,
		value,
	})
	return nil
}

type Options struct {
	URL      string
	Headers  headerFlag
	ShowHelp bool
}

func parseArguments() Options {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage: reqbud [options...] <url>\n")
		// flag.PrintDefaults()
	}

	var defaultOptions Options
	var options = defaultOptions

	// flag.BoolVar(&options.ShowHelp, "help", defaultOptions.ShowHelp, "This help text")
	flag.BoolVar(&options.ShowHelp, "h", defaultOptions.ShowHelp, "This help text")

	headers := headerFlag{}
	// flag.Var(&headers, "header", "Pass custom header(s) to server")
	flag.Var(&headers, "H", "Pass custom header(s) to server")

	flag.Parse()
	options.Headers = headers

	if options.ShowHelp {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage: reqbud [options...] <url>\n")
		flag.PrintDefaults()
	}

	url := flag.Arg(0)
	if len(url) == 0 {
		fmt.Println("type 'reqbud -h' for more information")
		os.Exit(0)
	}
	options.URL = url

	return options
}

func main() {
	opts := parseArguments()

	resp, err := http.Get(opts.URL)
	if err != nil {
		log.Fatalf("http.get error %v\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("http.get error %v\n", err)
	}

	log.Println(string(body))

	for name, values := range resp.Header {
		log.Printf("%s: %v\n", name, values)
	}
}
