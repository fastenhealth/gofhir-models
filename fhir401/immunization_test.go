package fhir401

import "testing"

func TestImmunizationJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[Immunization](t, "Immunization", "immunization")
}
