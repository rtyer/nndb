package nndb

import (
	"bufio"
	"errors"
	"io"
	"strings"
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
	switch parser.dataType {
	case fdGroupType:
		groups := []fdGroup{}
		for parser.scanner.Scan() {
			line := parser.scanner.Text()
			tokens := strings.Split(line, "^")
			groups = append(groups, fdGroup{
				code:        strings.Trim(tokens[0], "~"),
				description: strings.Trim(tokens[1], "~"),
			})
		}
		return groups, nil
	default:
		return nil, errors.New("Unsupported dataType")
	}
}

type fdGroup struct {
	code        string
	description string
}
