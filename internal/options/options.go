package options

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	URL      string
	ShowHelp bool
	Headers  HeaderFlag
	Request  string
	Data     string
}

func ParseArguments() Options {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage: reqbud [options...] <url>\n")
	}

	var defaultOptions Options
	var options = defaultOptions

	// flag.BoolVar(&options.ShowHelp, "help", defaultOptions.ShowHelp, "This help text")
	flag.BoolVar(&options.ShowHelp, "h", defaultOptions.ShowHelp, "This help text")

	// flag.StringVar(&options.Request, "data", defaultOptions.Request, "HTTP POST data")
	flag.StringVar(&options.Data, "d", defaultOptions.Request, "HTTP POST data")

	// flag.StringVar(&options.Request, "request", defaultOptions.Request, "Specify request command to use")
	flag.StringVar(&options.Request, "X", defaultOptions.Request, "Specify request command to use")

	headers := HeaderFlag{}
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
