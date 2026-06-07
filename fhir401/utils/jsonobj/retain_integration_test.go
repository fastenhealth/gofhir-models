package jsonobj_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fastenhealth/gofhir-models/fhir401/utils/jsonobj"
)

// Ensure Page is Retainable.
var _ any = jsonobj.MustRetainable(&Page{})

type Page struct {
	raw jsonobj.Retain

	Title string `json:"title"`
	Slug  string `json:"slug"`
}

func (p *Page) UnmarshalJSON(data []byte) error {
	return p.raw.FromJSON(data, p)
}

func (p Page) MarshalJSON() ([]byte, error) {
	return p.raw.ToJSON(p)
}

func Example_retainRoundtrip() {
	var p Page
	input := `{"title":"Contact Us","slug":"contact","icon":"email"}`
	if err := json.Unmarshal([]byte(input), &p); err != nil {
		log.Fatalf("Unmarshal failed: %v", err)
	}

	p.Slug = "contact-us" // update the slug value
	marshalled, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Marshal failed: %v", err)
	}

	// Note that output includes the updated "slug"
	// and retains the unknown "icon" field.
	fmt.Println(string(marshalled))
	// Output:
	// {"icon":"email","slug":"contact-us","title":"Contact Us"}
}
