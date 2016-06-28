package nndb

import (
	"strings"
	"testing"
)

func TestParseFoodGroup(t *testing.T) {
	input := "~0100~^~Dairy and Egg Products~\n~0200~^~Spices and Herbs~\n"
	parser, error := newReaderParser(strings.NewReader(input), fdGroupType)

	if error != nil {
		t.Errorf(`newReaderParser returned an error %v`, error)
	}
	if parser == nil {
		t.Error(`newReaderParser returned nil parser`)
	}

	foodGroups, error := parser.parse()

	if error != nil {
		t.Errorf(`parse() returned an error %v`, error)
	}
	if foodGroups == nil {
		t.Error(`parse returned nil`)
	}
}
