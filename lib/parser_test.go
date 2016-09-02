package nndb

import (
	"fmt"
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
var weightInput = `~01001~^1^1^~pat (1" sq, 1/3" high)~^5.0^^
~01001~^2^1^~tbsp~^14.2^^
~01001~^3^1^~cup~^227^^
~01001~^4^1^~stick~^113^^
~01002~^1^1^~pat (1" sq, 1/3" high)~^3.8^^
~01002~^2^1^~tbsp~^9.4^^
~01002~^3^1^~cup~^151^^
~01002~^4^1^~stick~^76^^
~01116~^1^4^~oz~^113^^
~01116~^2^1^~cup (not packed)~^226^^`

func TestParseNutrientDescription(t *testing.T) {

	parser, error := NewParser(strings.NewReader(""), strings.NewReader(""), strings.NewReader(nutrientInput), strings.NewReader(""))
	isValidParser(parser, error, t)

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

	nutrientsShouldMatch(nutrients[1001], Nutrients{Calories: 717, Fiber: 0, Fat: 81.11, Protein: 0.85, Sugar: 0.06}, t)
	nutrientsShouldMatch(nutrients[1002], Nutrients{Calories: 555}, t)
}

func nutrientsShouldMatch(actual Nutrients, expected Nutrients, t *testing.T) {
	if actual.Calories != expected.Calories {
		t.Errorf(`incorrect amount of calories %v when expected %v`, actual.Calories, expected.Calories)
	}
	if actual.Fiber != expected.Fiber {
		t.Errorf(`incorrect amount of fiber %v when expected %v`, actual.Fiber, expected.Fiber)
	}
	if actual.Fat != expected.Fat {
		t.Errorf(`incorrect amount of fat %v when expected %v`, actual.Fat, expected.Fat)
	}
	if actual.Protein != expected.Protein {
		t.Errorf(`incorrect amount of protein %v when expected %v`, actual.Protein, expected.Protein)
	}
	if actual.Sugar != expected.Sugar {
		t.Errorf(`incorrect amount of sugar %v when expected %v`, actual.Sugar, expected.Sugar)
	}
}

func isValidParser(parser Parser, e error, t *testing.T) {
	if e != nil {
		t.Errorf(`newReaderParser returned an error %v`, e)
	}
	if parser == nil {
		t.Error(`newReaderParser returned nil parser`)
	}
}

func TestParse(t *testing.T) {
	parser, error := NewParser(strings.NewReader(foodInput), strings.NewReader(foodGroupInput), strings.NewReader(nutrientInput), strings.NewReader(weightInput))
	isValidParser(parser, error, t)

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
	if len(food[0].Measurements) != 2 {
		t.Error(`Incorrect number of measurements`)
	}

	fmt.Println(food[0])
}

func TestParseFoodGroup(t *testing.T) {
	parser, error := NewParser(strings.NewReader(""), strings.NewReader(foodGroupInput), strings.NewReader(""), strings.NewReader(""))
	isValidParser(parser, error, t)

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

func TestParseWeight(t *testing.T) {
	parser, error := NewParser(strings.NewReader(""), strings.NewReader(foodGroupInput), strings.NewReader(""), strings.NewReader(weightInput))
	isValidParser(parser, error, t)

	weights, error := parser.parseWeights()

	if error != nil {
		t.Errorf(`Parse() returned an error %v`, error)
	}
	if weights == nil {
		t.Errorf(`weights returned nil`)
	}

	if len(weights[1001]) != 4 {
		t.Error(`incorrect number of measurements for weights[1001]`)
	}

	if weights[1001][1].Amount != 1 || weights[1001][1].Unit != "tbsp" || weights[1001][1].Weight != 14.2 {
		t.Errorf("Measurement does not match expected.  1 tbsp : 14.2 expected, saw %v", weights[1001][1])
	}
}

func TestParseWrongFileFormat(t *testing.T) {
	input := "blahblah"
	parser, _ := NewParser(strings.NewReader(input), strings.NewReader(input), strings.NewReader(input), strings.NewReader(input))

	_, error := parser.parseFoodGroups()
	if error == nil {
		t.Error(`invalid data should cause error`)
	}

	_, error = parser.parseNutrients()
	if error == nil {
		t.Error(`invalid data should cause error`)
	}

	_, error = parser.parseWeights()
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
	parser, _ := NewParser(strings.NewReader(input), strings.NewReader(input), strings.NewReader(input), strings.NewReader(input))

	_, error := parser.parseFoodGroups()
	if error != nil {
		t.Error(`empty data should not cause error`)
	}

	_, error = parser.parseNutrients()
	if error != nil {
		t.Error(`empty data should not cause error`)
	}

	_, error = parser.parseWeights()
	if error != nil {
		t.Error(`empty data should not cause error`)
	}

	_, error = parser.Parse()
	if error != nil {
		t.Error(`empty data should not cause error`)
	}

}
