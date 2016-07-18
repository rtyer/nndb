package nndb

import (
	"fmt"
	"strconv"
)

// Models representing data sourced from the National Nutrient Database

// Food represents a food item
// TODO: Add Measurements
type Food struct {
	ID            int
	Name          string
	AlternateName string
	Manufacturer  string
	FoodGroup     FoodGroup
	Nutrients     Nutrients
}

func (f Food) String() string {
	return fmt.Sprintf("Food {ID: %v Name: %v, AltName: %v, Manufacturer: %v, %v, %v}", f.ID, f.Name, strOrDefault(f.AlternateName, "n/a"), strOrDefault(f.Manufacturer, "n/a"), f.FoodGroup, f.Nutrients)
}

//Fields exports all fields (flattened) as a slice of strings
func (f Food) Fields() []string {
	return []string{
		strconv.Itoa(f.ID),
		f.Name,
		strOrDefault(f.AlternateName, "n/a"),
		strOrDefault(f.Manufacturer, "n/a"),
		strconv.Itoa(f.FoodGroup.ID),
		f.FoodGroup.Name,
		strconv.FormatFloat(f.Nutrients.Calories, 'f', 6, 64),
		strconv.FormatFloat(f.Nutrients.Fat, 'f', 6, 64),
		strconv.FormatFloat(f.Nutrients.Sugar, 'f', 6, 64),
		strconv.FormatFloat(f.Nutrients.Protein, 'f', 6, 64),
		strconv.FormatFloat(f.Nutrients.Fiber, 'f', 6, 64),
	}
}

// FoodGroup represents a food group, such as `Cereal Grains and Pasta`
type FoodGroup struct {
	ID   int
	Name string
}

func (fg FoodGroup) String() string {
	return fmt.Sprintf("FoodGroup {ID: %v Name: %v}", fg.ID, fg.Name)
}

// Nutrients represents the nutritional information, such as `Calories` per 100g
type Nutrients struct {
	// Calories in kcal
	Calories float64
	// Fat in g
	Fat float64
	// Sugar in g
	Sugar float64
	// Protein in g
	Protein float64
	// Fiber in g
	Fiber float64
}

func (n Nutrients) String() string {
	return fmt.Sprintf("Nutrients {Calories: %v kcal, Fat: %vg, Sugar: %vg, Protein: %vg, Fiber: %vg}", n.Calories, n.Fat, n.Sugar, n.Protein, n.Fiber)
}

// Measurement represents the number of grams of weight that common measurements have (per Food)
type Measurement struct {
	Amount float64
	Unit   string
	Weight float64
}

func strOrDefault(s string, def string) string {
	if s == "" {
		return def
	}
	return s
}
