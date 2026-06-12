package fhir401

import "testing"

func TestProvenanceJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[Provenance](t, "Provenance", "provenance")
}
