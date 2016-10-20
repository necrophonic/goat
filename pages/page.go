// Package page handles representations of pages
package pages

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"gopkg.in/yaml.v2"
)

type (
	// Page is the representation of a page
	Page struct {
		URL      string `json:"url",yaml:"url"`
		Elements []Element
	}

	// Element represents a page element such as a piece of text or an input field
	Element struct {
		Name       string   `json:"name",yaml:"name"`
		LocateByID string   `json:"locate-by-id",yaml:"locate-by-id"`
		Type       string   `json:"type",yaml:"type"`
		ExitsTo    []string `json:"exits-to",yaml:"exits-to"`
	}
)

var pageyaml = `
url: https://beta.companieshouse.gov.uk

elements:
  - name: page-title
    locate-by-xpath: //h1[@id='start-searching']/label[@class='heading-xlarge']
    type: text

  - name: search-box
    locate-by-id: site-search-text
    type: input-text

  - name: search-button
    locate-by-id: search-submit
    type: button
    exits-to:
      - search-results
      - somewhere

`

var pagejson = `
{
  "url"      : "https://beta.companieshouse.gov.uk",
  "elements" : [
    {
      "name" : "page-title",
      "locate-by-xpath" : "//h1[@id='start-searching']/label[@class='heading-xlarge']",
      "type" : "text"
    },
    {
      "name" : "search-box",
      "locate-by-id" : "site-search-text",
      "type" : "input-text"
    },
    {
      "name" : "search-button",
      "locate-by-id" : "search-submit",
      "type" : "button",
      "exits-to" : [
        "search-results",
        "somewhere"
      ]
    }
  ]
}
`

func Build() {
	y := Page{}

	err := yaml.Unmarshal([]byte(pageyaml), &y)
	if err != nil {
		panic(err)
	}
	fmt.Printf("YAML: %v\n", y)
	spew.Dump(y)

	j := Page{}

	err = json.Unmarshal([]byte(pagejson), &j)
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON: %v\n", j)
	spew.Dump(j)
}
