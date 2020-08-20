package genderize_test

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/alexeyco/genderize"
)

type requestTableRow struct {
	should string
	given  string
}

func TestNewRequest_Encode(t *testing.T) {
	table := []requestTableRow{
		{
			should: "https://api.genderize.io?name=Alice",
			given:  genderize.NewRequest().Name("Alice").Encode(),
		},
		{
			should: "https://api.genderize.io?name=Alice&country_id=US",
			given:  genderize.NewRequest().Name("Alice").CountryID("US").Encode(),
		},
		{
			should: "https://api.genderize.io?name=Alice&name=John&country_id=US",
			given:  genderize.NewRequest().Name("Alice").Name("John").CountryID("US").Encode(),
		},
		{
			should: "https://api.genderize.io?name=Alice&name=John&country_id=US&apikey=MyAwesomeAPIKey",
			given:  genderize.NewRequest().Name("Alice").Name("John").CountryID("US").Encode("MyAwesomeAPIKey"),
		},
	}

	for _, r := range table {
		should, _ := url.Parse(r.should)
		given, _ := url.Parse(r.given)

		if !reflect.DeepEqual(should.Query(), given.Query()) {
			t.Errorf(`URL should be "%s", "%s" given`, should.String(), given.String())
		}
	}
}
