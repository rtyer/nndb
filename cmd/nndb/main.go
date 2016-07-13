package main

import (
	"fmt"

	"github.com/rtyer/nndb"
)

func main() {
	n := nndb.Nutrients{
		Calories: 100,
	}
	fmt.Printf("hi%v\n", n)
}
