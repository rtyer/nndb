package nndb

import (
	"strings"
	"testing"
)

func TestParseFoodDescription(t *testing.T) {
	input := "~01116~^~0100~^~Yogurt, plain, whole milk, 8 grams protein per 8 ounce~^~YOGURT,PLN,WHL MILK,8 GRAMS PROT PER 8 OZ~^~~^~~^~Y~^~~^0^~~^6.38^4.27^8.79^3.87"
	parser, error := newFoodDescriptionParser(strings.NewReader(input))
	if error != nil {
		t.Errorf(`newReaderParser returned an error %v`, error)
	}
	if parser == nil {
		t.Error(`newReaderParser returned nil parser`)
	}

	result, dataType, error := parser.parse()

	if error != nil {
		t.Errorf(`parse() returned an error %v`, error)
	}
	if result == nil {
		t.Error(`parse returned nil`)
	}
	groups := []foodDescription{}
	switch dataType {
	case foodType:
		groups = result.([]foodDescription)
	default:
		t.Error(`Wrong type`)
	}
	if groups[0].ndbNo != "01116" {
		t.Error(`incorect value for foodGroups[0].ndbNo`)
	}
	if groups[0].fdGroupCode != "0100" {
		t.Error(`incorect value for foodGroups[0].fdGroupCode`)
	}
	if groups[0].longDescription != "Yogurt, plain, whole milk, 8 grams protein per 8 ounce" {
		t.Error(`incorect value for foodGroups[0].longDescription`)
	}
}

func TestParseFoodGroup(t *testing.T) {
	input := "~0100~^~Dairy and Egg Products~\n~0200~^~Spices and Herbs~\n"
	parser, error := newFdGroupParser(strings.NewReader(input))

	if error != nil {
		t.Errorf(`newReaderParser returned an error %v`, error)
	}
	if parser == nil {
		t.Error(`newReaderParser returned nil parser`)
	}

	result, dataType, error := parser.parse()

	if error != nil {
		t.Errorf(`parse() returned an error %v`, error)
	}
	if result == nil {
		t.Error(`parse returned nil`)
	}

	groups := []fdGroup{}
	switch dataType {
	case fdGroupType:
		groups = result.([]fdGroup)
	default:
		t.Error(`Wrong type`)
	}

	if groups[0].code != "0100" {
		t.Error(`incorect value for foodGroups[0].code`)
	}
	if groups[0].description != "Dairy and Egg Products" {
		t.Error(`incorect value for foodGroups[0].code`)
	}

	if groups[1].code != "0200" {
		t.Error(`incorect value for foodGroups[0].code`)
	}
	if groups[1].description != "Spices and Herbs" {
		t.Error(`incorect value for foodGroups[0].code`)
	}
}

func TestParseWrongFileFormat(t *testing.T) {
	input := "blahblah"
	parser, _ := newFdGroupParser(strings.NewReader(input))
	_, _, error := parser.parse()

	if error == nil {
		t.Error(`invalid data should cause error`)
	}
	parser, _ = newFoodDescriptionParser(strings.NewReader(input))
	_, _, error = parser.parse()

	if error == nil {
		t.Error(`invalid data should cause error`)
	}
}

func TestParseEmptyDataOK(t *testing.T) {
	input := ""
	parser, _ := newFdGroupParser(strings.NewReader(input))
	_, _, error := parser.parse()

	if error != nil {
		t.Error(`empty data should not cause error`)
	}

	parser, _ = newFoodDescriptionParser(strings.NewReader(input))
	_, _, error = parser.parse()

	if error != nil {
		t.Error(`empty data should not cause error`)
	}
}
