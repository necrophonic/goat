package page

import (
	"io/ioutil"
	"os"
	"testing"
)

var examplejson = `
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

func TestBuildingFileFile(t *testing.T) {
	// Start by writing an example file we can read from
	fn := "testpage.json"
	if err := ioutil.WriteFile(fn, []byte(examplejson), 0777); err != nil {
		t.Error("failed to write example file for reading")
		t.FailNow()
	}
	defer os.Remove(fn)

	_, err := BuildFromFile(fn)
	if err != nil {
		t.Errorf("failed to build from file: %v", err)
		t.Fail()
	}
}
