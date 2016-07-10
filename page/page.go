// Package page deals with creation and management of web page structures
package page

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/necrophonic/log"
)

type (
	// Page is the representation of a page
	Page struct {
		Name     string `json:"name",yaml:"name"`
		URL      string `json:"url",yaml:"url"`
		Elements []Element
	}

	// Element represents a page element such as a piece of text or an input field
	Element struct {
		Name          string   `json:"name",yaml:"name"`
		LocateByID    string   `json:"locate-by-id",yaml:"locate-by-id"`
		LocateByXpath string   `json:"locate-by-xpath",yaml:"locate-by-xpath"`
		Type          string   `json:"type",yaml:"type"`
		ExitsTo       []string `json:"exits-to",yaml:"exits-to"`
	}
)

// BuildFromFile constructs a page from yaml or json data in a file
func BuildFromFile(filename string) (*Page, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	switch {
	case strings.HasSuffix(filename, ".json"):
		return BuildFromJSON(b)
	case strings.HasSuffix(filename, ".yaml"):
		return BuildFromYAML(b)
	default:
		return nil, errors.New("unable to build from file - unsupported file type")
	}
}

// BuildFromYAML constructs a page from yaml data
func BuildFromYAML(source []byte) (*Page, error) {
	log.Trace("building page from YAML")
	page := &Page{}
	if err := yaml.Unmarshal(source, &page); err != nil {
		return nil, err
	}
	// TODO ensure constructed fields are valid
	return page, nil
}

// BuildFromJSON constructs a page form json data
func BuildFromJSON(source []byte) (*Page, error) {
	log.Trace("building page from JSON")
	page := &Page{}
	if err := json.Unmarshal(source, &page); err != nil {
		return nil, err
	}
	// TODO ensure constructed fields are valid
	return page, nil
}
