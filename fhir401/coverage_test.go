package fhir401

import "testing"

func TestCoverageJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[Coverage](t, "Coverage", "coverage")
}
