package fhir401

import "testing"

func TestPractitionerJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[Practitioner](t, "Practitioner", "practitioner")
}
