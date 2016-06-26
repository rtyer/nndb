package nndb

// Models representing data sourced from the National Nutrient Database

// Food represents a food item
type Food struct {
	ID            int
	FoodGroup     FoodGroup
	Name          string
	AlternateName string
	Manufacturer  string
	Nutrients     Nutrients
}

// FoodGroup represents a food group, such as `Cereal Grains and Pasta`
type FoodGroup struct {
	ID   int
	Name string
}

// Nutrients represents the nutritional information, such as `Calories`
type Nutrients struct {
	Calories int
}
