package fhir401

import "testing"

func TestAllergyIntoleranceJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[AllergyIntolerance](t, "AllergyIntolerance", "allergy_intolerance")
}
