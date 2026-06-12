package fhir401

import "testing"

func TestLocationJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[Location](t, "Location", "location")
}
