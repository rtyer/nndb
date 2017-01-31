package nndb

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	calories = "208"
	fat      = "204"
	protein  = "203"
	sugar    = "269"
	fiber    = "291"
)

const (
	//FoodGroupFile is the location for the food group file
	FoodGroupFile = "FD_GROUP.txt"
	//FoodDesFile is the location for the food description file
	FoodDesFile = "FOOD_DES.txt"
	//NutrDefFile is the location for the nutrient description file
	NutrDefFile = "NUT_DATA.txt"
	//WeightFile is the location for the nutrient weights file
	WeightFile = "WEIGHT.txt"
)

type scannerParser struct {
	foodDesScanner *bufio.Scanner
	fdGroupScanner *bufio.Scanner
	nutrDefScanner *bufio.Scanner
	weightScanner  *bufio.Scanner
}

// NewParser returns a parser configured for the passed readers. All readers must be non-nil or an error will be returned.
func NewParser(foodDesReader io.Reader, fdGroupReader io.Reader, nutrDefReader io.Reader, weightScanner io.Reader) (Parser, error) {
	if foodDesReader == nil || fdGroupReader == nil || nutrDefReader == nil {
		return nil, errors.New("All readers must be valid")
	}
	return scannerParser{
		foodDesScanner: bufio.NewScanner(foodDesReader),
		fdGroupScanner: bufio.NewScanner(fdGroupReader),
		nutrDefScanner: bufio.NewScanner(nutrDefReader),
		weightScanner:  bufio.NewScanner(weightScanner),
	}, nil

}

// Parser provides the means to convert from the National Nutrient Database file formats into a []Food
type Parser interface {
	parseFoodGroups() (map[int]FoodGroup, error)
	parseNutrients() (map[int]Nutrients, error)
	parseWeights() (map[int][]Measurement, error)
	Parse() ([]Food, error)
}

// parseFoodGroups is an internal method that will return a map of food group code/id : FoodGroup parsed from the scanner
func (parser scannerParser) parseFoodGroups() (map[int]FoodGroup, error) {
	groups := make(map[int]FoodGroup)
	for parser.fdGroupScanner.Scan() {
		line := parser.fdGroupScanner.Text()
		tokens := strings.Split(line, "^")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("Invalid Format.  Expected 2 tokens and saw %d", len(tokens))
		}

		id, err := strconv.Atoi(strings.Trim(tokens[0], "~"))
		if err != nil {
			return nil, err
		}

		groups[id] = FoodGroup{
			ID:   id,
			Name: strings.Trim(tokens[1], "~"),
		}
	}
	return groups, nil
}

// parseNutrients is an internal method that will return a map of ndbNO/Food.ID : Nutrients parsed from the scanner
func (parser scannerParser) parseNutrients() (map[int]Nutrients, error) {
	nutrientMap := make(map[int]Nutrients)
	for parser.nutrDefScanner.Scan() {
		line := parser.nutrDefScanner.Text()

		tokens := strings.Split(line, "^")
		if len(tokens) != 18 {
			return nil, fmt.Errorf("Invalid Format.  Expected 18 tokens and saw %d", len(tokens))
		}

		foodID, err := strconv.Atoi(strings.Trim(tokens[0], "~"))
		if err != nil {
			return nil, err
		}

		nutrientID := strings.Trim(tokens[1], "~")

		err = extractNutrientValue(nutrientID, foodID, nutrientMap, tokens)
		if err != nil {
			return nil, err
		}
	}
	return nutrientMap, nil
}

func extractNutrientValue(nutrientID string, foodID int, nutrientMap map[int]Nutrients, tokens []string) error {
	if isValidNutrient(nutrientID) {
		nutrient, _ := nutrientMap[foodID]

		f, err := strconv.ParseFloat(strings.Trim(tokens[2], "~"), 64)
		if err != nil {
			return fmt.Errorf("Could not parse nutrient value for %v", nutrientID)
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
		nutrientMap[foodID] = nutrient
	}
	return nil
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

func (parser scannerParser) parseWeights() (map[int][]Measurement, error) {
	weightsMap := make(map[int][]Measurement)
	for parser.weightScanner.Scan() {
		line := parser.weightScanner.Text()
		tokens := strings.Split(line, "^")
		if len(tokens) != 7 {
			return nil, fmt.Errorf("Invalid Format.  Expected 6 tokens and saw %d", len(tokens))
		}

		foodID, err := strconv.Atoi(strings.Trim(tokens[0], "~"))
		if err != nil {
			return nil, err
		}

		amount, err := strconv.ParseFloat(strings.Trim(tokens[2], "~"), 64)
		if err != nil {
			return nil, err
		}

		unit := strings.Trim(tokens[3], "~")

		weight, err := strconv.ParseFloat(strings.Trim(tokens[4], "~"), 64)
		if err != nil {
			return nil, err
		}

		weightsMap[foodID] = append(weightsMap[foodID], Measurement{
			Amount: amount,
			Unit:   unit,
			Weight: weight,
		})
	}
	return weightsMap, nil
}

func (parser scannerParser) Parse() ([]Food, error) {
	foodGroups, err := parser.parseFoodGroups()
	if err != nil {
		return nil, err
	}
	nutrients, err := parser.parseNutrients()
	if err != nil {
		return nil, err
	}

	weights, err := parser.parseWeights()
	if err != nil {
		return nil, err
	}

	food := []Food{}
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
		food = append(food, Food{
			ID:            id,
			FoodGroup:     foodGroups[foodGroupID],
			Name:          strings.Trim(tokens[2], "~"),
			AlternateName: strings.Trim(tokens[3], "~"),
			Manufacturer:  strings.Trim(tokens[5], "~"),
			Nutrients:     nutrients[id],
			Measurements:  weights[id],
		})
	}
	return food, nil
}
