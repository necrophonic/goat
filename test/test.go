package test

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Status represents the current status of a test or scenario
type Status uint8

// Status constants
const (
	NotRun Status = iota
	Failed
	Passed
	Skipped
)

// Test represents a test to be excuted. It comprises metadata about the test as well as the
// execution script for each included scenario.
type Test struct {
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	Scenarios   []Scenario `yaml:"scenarios"`
	Status      Status     `yaml:"-"`
}

// Scenario represents the test script for a scenario supporting the main test
type Scenario struct {
	Name       string        `yaml:"name"`   // Name of the specific scenario
	ScriptText []string      `yaml:"script"` // Raw textual script from the specification file
	Script     []interface{} `yaml:"-"`      // Parsed lua script for execution
	Status     Status        `yaml:"-"`
}

// ImportFromFile creates a new spec object from a given spec file
func ImportFromFile(filename string) *Test {

	y, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read expected spec file [%s]: %v", filename, err)
	}

	test := &Test{}
	if err := yaml.Unmarshal(y, &test); err != nil {
		log.Fatalf("Failed to unmarshal spec file [%s]: %v", filename, err)
	}

	return test
}
