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

// Nutrients represents the nutritional information, such as `Calories` per 100g
type Nutrients struct {
	// Calories in kcal
	Calories float32
	// Fat in g
	Fat float32
	// Sugar in g
	Sugar float32
	// Protein in g
	Protein float32
	// Fiber in g
	Fiber float32
}

// Measurement represents the number of grams of weight that common measurements have (per Food)
type Measurement struct {
	Amount float32
	Unit   string
	Weight float32
}
