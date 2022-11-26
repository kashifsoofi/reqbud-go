package options

import (
	"errors"
	"strings"
)

type Header struct {
	Name  string
	Value string
}

type HeaderFlag []Header

func (f *HeaderFlag) String() string {
	return ""
}

func (f *HeaderFlag) Set(arg string) error {
	name, value, found := strings.Cut(arg, ":")
	if !found {
		return errors.New("invalid header, must be name:value")
	}

	*f = append(*f, Header{
		name,
		value,
	})
	return nil
}
