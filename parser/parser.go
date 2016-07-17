package nndb

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/rtyer/nndb"
)

const (
	calories = "208"
	fat      = "204"
	protein  = "203"
	sugar    = "269"
	fiber    = "291"
)

type scannerParser struct {
	foodDesScanner *bufio.Scanner
	fdGroupScanner *bufio.Scanner
	nutrDefScanner *bufio.Scanner
}

// NewParser returns a parser configured for the passed readers. All readers must be non-nil or an error will be returned.
func NewParser(foodDesReader io.Reader, fdGroupReader io.Reader, nutrDefReader io.Reader) (Parser, error) {
	if foodDesReader == nil || fdGroupReader == nil || nutrDefReader == nil {
		return nil, errors.New("All readers must be valid")
	}
	return scannerParser{
		foodDesScanner: bufio.NewScanner(foodDesReader),
		fdGroupScanner: bufio.NewScanner(fdGroupReader),
		nutrDefScanner: bufio.NewScanner(nutrDefReader),
	}, nil

}

// Parser provides the means to convert from the National Nutrient Database file formats into a []nndb.Food
type Parser interface {
	parseFoodGroups() (map[int]nndb.FoodGroup, error)
	parseNutrients() (map[int]nndb.Nutrients, error)
	Parse() ([]nndb.Food, error)
}

// parseFoodGroups is an internal method that will return a map of food group code/id : nndb.FoodGroup parsed from the scanner
func (parser scannerParser) parseFoodGroups() (map[int]nndb.FoodGroup, error) {
	groups := make(map[int]nndb.FoodGroup)
	for parser.fdGroupScanner.Scan() {
		line := parser.fdGroupScanner.Text()
		tokens := strings.Split(line, "^")
		if len(tokens) != 2 {
			return nil, errors.New("Invalid Format")
		}

		id, err := strconv.Atoi(strings.Trim(tokens[0], "~"))
		if err != nil {
			return nil, err
		}

		groups[id] = nndb.FoodGroup{
			ID:   id,
			Name: strings.Trim(tokens[1], "~"),
		}
	}
	return groups, nil
}

// parseNutrients is an internal method that will return a map of ndbNO/nndb.Food.ID : nndb.Nutrients parsed from the scanner
func (parser scannerParser) parseNutrients() (map[int]nndb.Nutrients, error) {
	nutrientMap := make(map[int]nndb.Nutrients)
	for parser.nutrDefScanner.Scan() {
		line := parser.nutrDefScanner.Text()
		tokens := strings.Split(line, "^")
		if len(tokens) != 18 {
			return nil, errors.New("Invalid Format")
		}

		id, err := strconv.Atoi(strings.Trim(tokens[0], "~"))
		if err != nil {
			return nil, err
		}

		nutrientID := strings.Trim(tokens[1], "~")

		if isValidNutrient(nutrientID) {
			nutrient, _ := nutrientMap[id]

			f, err := strconv.ParseFloat(strings.Trim(tokens[2], "~"), 64)
			if err != nil {
				return nil, fmt.Errorf("Could not parse nutrient value for %v", nutrientID)
			}

			switch nutrientID {
			case calories:
				nutrient.Calories = f
			case fat:
				nutrient.Fat = f
			case sugar:
				nutrient.Sugar = f
			case fiber:
				nutrient.Fiber = f
			case protein:
				nutrient.Protein = f
			}
			nutrientMap[id] = nutrient
		}
	}
	return nutrientMap, nil
}

func (parser scannerParser) Parse() ([]nndb.Food, error) {
	foodGroups, err := parser.parseFoodGroups()
	if err != nil {
		return nil, err
	}
	nutrients, err := parser.parseNutrients()
	if err != nil {
		return nil, err
	}

	food := []nndb.Food{}
	for parser.foodDesScanner.Scan() {
		line := parser.foodDesScanner.Text()
		tokens := strings.Split(line, "^")
		if len(tokens) != 14 {
			return nil, errors.New("Invalid Format")
		}
		id, err := strconv.Atoi(strings.Trim(tokens[0], "~"))
		if err != nil {
			return nil, err
		}
		foodGroupID, err := strconv.Atoi(strings.Trim(tokens[1], "~"))
		if err != nil {
			return nil, err
		}
		food = append(food, nndb.Food{
			ID:            id,
			FoodGroup:     foodGroups[foodGroupID],
			Name:          strings.Trim(tokens[2], "~"),
			AlternateName: strings.Trim(tokens[3], "~"),
			Manufacturer:  strings.Trim(tokens[5], "~"),
			Nutrients:     nutrients[id],
		})
	}
	return food, nil
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
