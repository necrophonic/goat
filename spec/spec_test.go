package spec

import (
	"testing"

	"github.com/necrophonic/goat/spec"
)

func TestImportSpecFile(t *testing.T) {

	s := spec.Import("../testing/simple-1.test.yml")

	if s.Name != "Simple test 1" {
		t.Error("Test name not as expected")
	}
	if s.Description != "A simple test layout" {
		t.Error("Description not as expected")
	}
	if len(s.Scenarios) != 2 {
		t.Error("Expected two scenarios")
	}
}
