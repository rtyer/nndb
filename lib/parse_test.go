package nndb

import (
	"strings"
	"testing"
)

func TestParseFoodGroup(t *testing.T) {
	parser, error := newReaderParser(strings.NewReader("hi"), fdGroupType)
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
