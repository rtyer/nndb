// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	nndb "github.com/rtyer/nndb/lib"
	"github.com/spf13/cobra"
)

var sourceDirectory string
var outfile string

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parses the standard format from national nutrient database into a consolidated format.",
	Long:  `Parses the standard format from national nutrient database into a consolidated format.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("parse called with %v\n", sourceDirectory)

		parser, err := nndb.NewParser(
			newReaderForFilename(sourceDirectory, nndb.FoodDesFile),
			newReaderForFilename(sourceDirectory, nndb.FoodGroupFile),
			newReaderForFilename(sourceDirectory, nndb.NutrDefFile),
			newReaderForFilename(sourceDirectory, nndb.WeightFile))
		if err != nil {
			fmt.Println(err)
		} else {
			food, err := parser.Parse()
			if err != nil {
				fmt.Println(err)
			}
			b, err := json.Marshal(food)

			if err != nil {
				fmt.Println(err)
			}
			err = ioutil.WriteFile(outfile, b, 0644)
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(parseCmd)

	parseCmd.Flags().StringVarP(&sourceDirectory, "source", "s", "", "source directory to read from")
	parseCmd.Flags().StringVarP(&outfile, "outfile", "o", "", "file to output to")
}

func newReaderForFilename(folder string, filename string) io.Reader {
	path := folder + "/" + filename
	// fmt.Printf("called with %v\n", path)
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	return file
}
