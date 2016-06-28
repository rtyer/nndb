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
	unknown
)

type scannerParser struct {
	scanner  *bufio.Scanner
	dataType dataType
}

func newFdGroupParser(reader io.Reader) (parser, error) {
	return scannerParser{
		scanner:  bufio.NewScanner(reader),
		dataType: fdGroupType,
	}, nil
}

type parser interface {
	parse() (interface{}, dataType, error)
}

func (parser scannerParser) parse() (interface{}, dataType, error) {
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
		return groups, parser.dataType, nil
	default:
		return nil, unknown, errors.New("Unsupported dataType")
	}
}

type fdGroup struct {
	code        string
	description string
}
