package genderize_test

import (
	"context"
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
	ctx := context.TODO()

	table := []requestTableRow{
		{
			should: "https://api.genderize.io?name%5B%5D=Alice",
			given:  genderize.NewRequest(ctx).Name("Alice").Encode(),
		},
		{
			should: "https://api.genderize.io?name%5B%5D=Alice&country_id=US",
			given:  genderize.NewRequest(ctx).Name("Alice").CountryID("US").Encode(),
		},
		{
			should: "https://api.genderize.io?name%5B%5D=Alice&name%5B%5D=John&country_id=US",
			given:  genderize.NewRequest(ctx).Name("Alice").Name("John").CountryID("US").Encode(),
		},
		{
			should: "https://api.genderize.io?name%5B%5D=Alice&name%5B%5D=John&country_id=US&apikey=MyAwesomeAPIKey",
			given:  genderize.NewRequest(ctx).Name("Alice").Name("John").CountryID("US").Encode("MyAwesomeAPIKey"),
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
