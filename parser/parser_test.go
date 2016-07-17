package nndb

import (
	"strings"
	"testing"
)

func TestParseNutrientDescription(t *testing.T) {
	input := `~01001~^~203~^0.85^16^0.074^~1~^~~^~~^~~^^^^^^^~~^11/1976^
~01001~^~204~^81.11^580^0.065^~1~^~~^~~^~~^^^^^^^~~^11/1976^
~01001~^~205~^0.06^0^^~4~^~NC~^~~^~~^^^^^^^~~^11/1976^
~01001~^~207~^2.11^35^0.054^~1~^~~^~~^~~^^^^^^^~~^11/1976^
~01001~^~208~^717^0^^~4~^~NC~^~~^~~^^^^^^^~~^08/2010^
~01001~^~221~^0.0^0^^~7~^~~^~~^~~^^^^^^^~~^04/1985^
~01001~^~255~^15.87^522^0.061^~1~^~~^~~^~~^^^^^^^~~^11/1976^
~01001~^~262~^0^0^^~7~^~Z~^~~^~~^^^^^^^~~^02/2001^
~01001~^~263~^0^0^^~7~^~Z~^~~^~~^^^^^^^~~^02/2001^
~01001~^~268~^2999^0^^~4~^~NC~^~~^~~^^^^^^^~~^09/2015^
~01001~^~269~^0.06^0^^~4~^~NR~^~~^~~^^^^^^^~~^11/2002^
~01001~^~291~^0.0^0^^~4~^~~^~~^~~^^^^^^^~~^^
~01001~^~301~^24^17^0.789^~1~^~A~^~~^~~^7^19^30^4^22.021^26.496^~2, 3~^11/2002^`

	parser, error := newNutrientParser(strings.NewReader(input))
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
	if dataType != nutrientType {
		t.Error(`Wrong type`)
	}

	nutrients := result.([]nutrientDescription)

	if len(nutrients) != 1 {
		t.Errorf(`incorrect number of results %v`, len(nutrients))
	}
	if nutrients[0].calories != 717 {
		t.Errorf(`incorrect amount of calories %v`, nutrients[0].calories)
	}
	if nutrients[0].fiber != 0 {
		t.Errorf(`incorrect amount of fiber %v`, nutrients[0].fiber)
	}
	if nutrients[0].fat != 81.11 {
		t.Errorf(`incorrect amount of fiber %v`, nutrients[0].fat)
	}
	if nutrients[0].protein != 0.85 {
		t.Errorf(`incorrect amount of fiber %v`, nutrients[0].protein)
	}
	if nutrients[0].sugar != 0.06 {
		t.Errorf(`incorrect amount of fiber %v`, nutrients[0].sugar)
	}
}

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
	if dataType != foodType {
		t.Error(`Wrong type`)
	}
	groups := result.([]foodDescription)

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

	if dataType != fdGroupType {
		t.Error(`Wrong type`)
	}
	groups := result.([]fdGroup)

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
	parser, _ = newNutrientParser(strings.NewReader(input))
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

	parser, _ = newNutrientParser(strings.NewReader(input))
	_, _, error = parser.parse()

	if error != nil {
		t.Error(`empty data should not cause error`)
	}
}
