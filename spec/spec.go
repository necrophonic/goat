package spec

import (
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"

	"gopkg.in/yaml.v2"
)

// Spec represents test test script(s) read and created from a spec file
type Spec struct {
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	Scenarios   []Scenario `yaml:"scenarios"`
}

// Scenario represents the test script for a scenario supporting the main test
type Scenario struct {
	Name   string   `yaml:"name"`
	Script []string `yaml:"script"`
}

// Import creates a new spec object from a given spec file
func Import(filename string) *Spec {
	sp := &Spec{}

	y, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err) // TODO
	}

	if err := yaml.Unmarshal(y, &sp); err != nil {
		panic(err) // TODO
	}

	spew.Dump(&sp)

	return sp
}
