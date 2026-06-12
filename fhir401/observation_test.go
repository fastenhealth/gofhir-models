package fhir401

import "testing"

func TestObservationJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[Observation](t, "Observation", "observation")
}
