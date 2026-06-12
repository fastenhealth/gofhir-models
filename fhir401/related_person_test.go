package fhir401

import "testing"

func TestRelatedPersonJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[RelatedPerson](t, "RelatedPerson", "related_person")
}
