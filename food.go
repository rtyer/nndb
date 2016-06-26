package nndb

// Food represents a food item, sourced from the Nutrient Databank
type Food struct {
	// NBDID is the Nutrient Databank ID
	ID            int
	FoodGroup     FoodGroup
	Name          string
	AlternateName string
	Manufacturer  string
}

// FoodGroup represents a food group, sourced from the Nutrient Databank
type FoodGroup struct {
	ID   int
	Name string
}
