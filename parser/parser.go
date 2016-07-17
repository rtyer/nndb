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
	foodType
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

func newFoodDescriptionParser(reader io.Reader) (parser, error) {
	return scannerParser{
		scanner:  bufio.NewScanner(reader),
		dataType: foodType,
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
			if len(tokens) != 2 {
				return nil, unknown, errors.New("Invalid Format")
			}
			groups = append(groups, fdGroup{
				code:        strings.Trim(tokens[0], "~"),
				description: strings.Trim(tokens[1], "~"),
			})
		}
		return groups, parser.dataType, nil
	case foodType:
		food := []foodDescription{}
		for parser.scanner.Scan() {
			line := parser.scanner.Text()
			tokens := strings.Split(line, "^")
			if len(tokens) != 14 {
				return nil, unknown, errors.New("Invalid Format")
			}
			food = append(food, foodDescription{
				ndbNo:            strings.Trim(tokens[0], "~"),
				fdGroupCode:      strings.Trim(tokens[1], "~"),
				longDescription:  strings.Trim(tokens[2], "~"),
				shortDescription: strings.Trim(tokens[3], "~"),
				commonName:       strings.Trim(tokens[4], "~"),
				manufacturerName: strings.Trim(tokens[5], "~"),
			})
		}
		return food, parser.dataType, nil
	default:
		return nil, unknown, errors.New("Unsupported dataType")
	}
}

type fdGroup struct {
	code        string
	description string
}

type foodDescription struct {
	ndbNo            string
	fdGroupCode      string
	longDescription  string
	shortDescription string
	commonName       string
	manufacturerName string
}
