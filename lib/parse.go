package nndb

import (
	"bufio"
	"io"
)

type dataType int

const (
	fdGroupType dataType = iota
)

type readerParser struct {
	scanner  *bufio.Scanner
	dataType dataType
}

func newReaderParser(reader io.Reader, dataType dataType) (parser, error) {
	return readerParser{
		scanner:  bufio.NewScanner(reader),
		dataType: dataType,
	}, nil
}

type parser interface {
	parse() (interface{}, error)
}

func (parser readerParser) parse() (interface{}, error) {
	// parse from info.reader into data
	// parse strategy depends on format
	return fdGroup{
		code:        "12345",
		description: "it's a thing",
	}, nil
}

type fdGroup struct {
	code        string
	description string
}
