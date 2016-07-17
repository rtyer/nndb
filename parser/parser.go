package nndb

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type dataType int

const (
	fdGroupType dataType = iota
	foodType
	nutrientType
	unknown
)

const (
	calories = "208"
	fat      = "204"
	protein  = "203"
	sugar    = "269"
	fiber    = "291"
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

func newNutrientParser(reader io.Reader) (parser, error) {
	return scannerParser{
		scanner:  bufio.NewScanner(reader),
		dataType: nutrientType,
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
	case nutrientType:
		// get first field, look up in map
		// if not present, add new empty
		// set the value for the field if it is one we care about
		// when done with all lines, create a slice of values
		nutrientMap := make(map[string]nutrientDescription)
		for parser.scanner.Scan() {
			line := parser.scanner.Text()
			tokens := strings.Split(line, "^")
			if len(tokens) != 18 {
				return nil, unknown, errors.New("Invalid Format")
			}
			ndbNo := strings.Trim(tokens[0], "~")
			nutrientID := strings.Trim(tokens[1], "~")
			if isValidNutrient(nutrientID) {
				nutrient, _ := nutrientMap[ndbNo]

				f, err := strconv.ParseFloat(strings.Trim(tokens[2], "~"), 64)
				if err != nil {
					return nil, parser.dataType, fmt.Errorf("Could not parse nutrient value for %v", nutrientID)
				}
				switch nutrientID {
				case calories:
					nutrient.calories = f
				case fat:
					nutrient.fat = f
				case sugar:
					nutrient.sugar = f
				case fiber:
					nutrient.fiber = f
				case protein:
					nutrient.protein = f
				}
				nutrientMap[ndbNo] = nutrient
			}
		}
		var nutrients []nutrientDescription
		for k := range nutrientMap {
			nutrients = append(nutrients, nutrientMap[k])
		}
		return nutrients, parser.dataType, nil

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

type nutrientDescription struct {
	calories float64
	fat      float64
	sugar    float64
	protein  float64
	fiber    float64
}

func isValidNutrient(nutrientID string) bool {
	switch nutrientID {
	case
		calories,
		fat,
		protein,
		fiber,
		sugar:
		return true
	}
	return false
}
