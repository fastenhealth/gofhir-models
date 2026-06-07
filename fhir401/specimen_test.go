package fhir401

import "testing"

func TestSpecimenJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[Specimen](t, "Specimen", "specimen")
}
