package nndb

import (
	"strings"
	"testing"
)

var nutrientInput = `~01001~^~203~^0.85^16^0.074^~1~^~~^~~^~~^^^^^^^~~^11/1976^
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
~01001~^~301~^24^17^0.789^~1~^~A~^~~^~~^7^19^30^4^22.021^26.496^~2, 3~^11/2002^
~01002~^~208~^555^0^^~4~^~NC~^~~^~~^^^^^^^~~^08/2010^
~01116~^~203~^3.47^0^^~1~^~~^~~^~~^^^^^^^~~^11/1976^
~01116~^~204~^3.25^0^^~1~^~~^~~^~~^^^^^^^~~^11/1976^
~01116~^~205~^4.66^0^^~4~^~NC~^~~^~~^^^^^^^~~^11/1976^
~01116~^~207~^0.72^0^^~1~^~~^~~^~~^^^^^^^~~^11/1976^
~01116~^~208~^61^0^^~4~^~NC~^~~^~~^^^^^^^~~^02/2009^
~01116~^~221~^0.0^0^^~7~^~~^~~^~~^^^^^^^~~^04/1985^
~01116~^~255~^87.90^0^^~1~^~~^~~^~~^^^^^^^~~^11/1976^
~01116~^~262~^0^0^^~7~^~Z~^~~^~~^^^^^^^~~^01/2003^
~01116~^~263~^0^0^^~7~^~Z~^~~^~~^^^^^^^~~^01/2003^
~01116~^~268~^257^0^^~4~^~NC~^~~^~~^^^^^^^~~^02/2009^
~01116~^~269~^4.66^0^^~4~^~NR~^~~^~~^^^^^^^~~^01/2003^
~01116~^~291~^0.0^0^^~4~^~~^~~^~~^^^^^^^~~^11/1976^`
var foodInput = "~01116~^~0100~^~Yogurt, plain, whole milk, 8 grams protein per 8 ounce~^~YOGURT,PLN,WHL MILK,8 GRAMS PROT PER 8 OZ~^~~^~~^~Y~^~~^0^~~^6.38^4.27^8.79^3.87"
var foodGroupInput = "~0100~^~Dairy and Egg Products~\n~0200~^~Spices and Herbs~\n"

func TestParseNutrientDescription(t *testing.T) {

	parser, error := NewParser(strings.NewReader(""), strings.NewReader(""), strings.NewReader(nutrientInput))
	if error != nil {
		t.Errorf(`newReaderParser returned an error %v`, error)
	}
	if parser == nil {
		t.Error(`newReaderParser returned nil parser`)
	}
	nutrients, error := parser.parseNutrients()
	if error != nil {
		t.Errorf(`Parse() returned an error %v`, error)
	}
	if nutrients == nil {
		t.Error(`parse returned nil`)
	}

	if len(nutrients) != 3 {
		t.Errorf(`incorrect number of results %v`, len(nutrients))
	}
	if nutrients[1001].Calories != 717 {
		t.Errorf(`incorrect amount of calories %v`, nutrients[1001].Calories)
	}
	if nutrients[1001].Fiber != 0 {
		t.Errorf(`incorrect amount of fiber %v`, nutrients[1001].Fiber)
	}
	if nutrients[1001].Fat != 81.11 {
		t.Errorf(`incorrect amount of fat %v`, nutrients[1001].Fat)
	}
	if nutrients[1001].Protein != 0.85 {
		t.Errorf(`incorrect amount of protein %v`, nutrients[1001].Protein)
	}
	if nutrients[1001].Sugar != 0.06 {
		t.Errorf(`incorrect amount of sugar %v`, nutrients[1001].Sugar)
	}
	if nutrients[1002].Calories != 555 {
		t.Errorf(`incorrect amount of calories %v`, nutrients[1002].Calories)
	}
	if nutrients[1002].Fat != 0 {
		t.Errorf(`incorrect amount of Fat %v`, nutrients[1002].Fat)
	}
}

func TestParse(t *testing.T) {
	parser, error := NewParser(strings.NewReader(foodInput), strings.NewReader(foodGroupInput), strings.NewReader(nutrientInput))
	if error != nil {
		t.Errorf(`newReaderParser returned an error %v`, error)
	}
	if parser == nil {
		t.Error(`newReaderParser returned nil parser`)
	}

	food, error := parser.Parse()

	if error != nil {
		t.Errorf(`Parse() returned an error %v`, error)
	}
	if food == nil {
		t.Error(`parse returned nil`)
	}
	if food[0].ID != 1116 {
		t.Error(`incorect value for foodGroups[0].ndbNo`)
	}
	if food[0].FoodGroup.ID != 100 {
		t.Error(`incorect food group`)
	}
	if food[0].Nutrients.Calories == 0 {
		t.Error(`Nutrients should be present`)
	}
	if food[0].Name != "Yogurt, plain, whole milk, 8 grams protein per 8 ounce" {
		t.Error(`Incorrect value for name`)
	}
}

func TestParseFoodGroup(t *testing.T) {
	parser, error := NewParser(strings.NewReader(""), strings.NewReader(foodGroupInput), strings.NewReader(""))

	if error != nil {
		t.Errorf(`newReaderParser returned an error %v`, error)
	}
	if parser == nil {
		t.Error(`newReaderParser returned nil parser`)
	}

	groups, error := parser.parseFoodGroups()

	if error != nil {
		t.Errorf(`Parse() returned an error %v`, error)
	}
	if groups == nil {
		t.Error(`parse returned nil`)
	}

	if groups[100].ID != 100 {
		t.Error(`incorect value for foodGroups[100].code`)
	}
	if groups[100].Name != "Dairy and Egg Products" {
		t.Error(`incorect value for foodGroups[100].code`)
	}

	if groups[200].ID != 200 {
		t.Error(`incorect value for foodGroups[200].code`)
	}
	if groups[200].Name != "Spices and Herbs" {
		t.Error(`incorect value for foodGroups[200].code`)
	}
}

func TestParseWrongFileFormat(t *testing.T) {
	input := "blahblah"
	parser, _ := NewParser(strings.NewReader(input), strings.NewReader(input), strings.NewReader(input))

	_, error := parser.parseFoodGroups()
	if error == nil {
		t.Error(`invalid data should cause error`)
	}

	_, error = parser.parseNutrients()
	if error == nil {
		t.Error(`invalid data should cause error`)
	}

	_, error = parser.Parse()
	if error == nil {
		t.Error(`invalid data should cause error`)
	}
}

func TestParseEmptyDataOK(t *testing.T) {
	input := ""
	parser, _ := NewParser(strings.NewReader(input), strings.NewReader(input), strings.NewReader(input))

	_, error := parser.parseFoodGroups()
	if error != nil {
		t.Error(`empty data should not cause error`)
	}

	_, error = parser.parseNutrients()
	if error != nil {
		t.Error(`empty data should not cause error`)
	}

	_, error = parser.Parse()
	if error != nil {
		t.Error(`empty data should not cause error`)
	}

}
