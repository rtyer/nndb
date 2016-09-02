package nndb

import "fmt"

// Models representing data sourced from the National Nutrient Database

// Food represents a food item
// TODO: Add Measurements
type Food struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	AlternateName string        `json:"alt_name"`
	Manufacturer  string        `json:"manufacturer"`
	FoodGroup     FoodGroup     `json:"food_group"`
	Nutrients     Nutrients     `json:"nutrients"`
	Measurements  []Measurement `json:"measurements"`
}

func (f Food) String() string {
	return fmt.Sprintf("Food {ID: %v Name: %v, AltName: %v, Manufacturer: %v, %v, %v, %v}", f.ID, f.Name, strOrDefault(f.AlternateName, "n/a"), strOrDefault(f.Manufacturer, "n/a"), f.FoodGroup, f.Nutrients, f.Measurements)
}

// FoodGroup represents a food group, such as `Cereal Grains and Pasta`
type FoodGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (fg FoodGroup) String() string {
	return fmt.Sprintf("FoodGroup {ID: %v Name: %v}", fg.ID, fg.Name)
}

// Nutrients represents the nutritional information, such as `Calories` per 100g.  All values in g other than kCal.
type Nutrients struct {
	Calories float64 `json:"calories"`
	Fat      float64 `json:"fat"`
	Sugar    float64 `json:"sugar"`
	Protein  float64 `json:"protein"`
	Fiber    float64 `json:"fiber"`
}

func (n Nutrients) String() string {
	return fmt.Sprintf("Nutrients {Calories: %.2f kCal, Fat: %.2fg, Sugar: %.2fg, Protein: %.2fg, Fiber: %.2fg}", n.Calories, n.Fat, n.Sugar, n.Protein, n.Fiber)
}

// Measurement represents the number of grams of weight that common measurements have (per Food)
type Measurement struct {
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
	Weight float64 `json:"weight"`
}

func (m Measurement) String() string {
	return fmt.Sprintf("Measurements {Amount: %.2f, Unit: %v, Weight: %.2f}", m.Amount, m.Unit, m.Weight)
}

func strOrDefault(s string, def string) string {
	if s == "" {
		return def
	}
	return s
}
