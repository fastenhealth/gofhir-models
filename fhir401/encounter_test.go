package fhir401

import "testing"

func TestEncounterJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[Encounter](t, "Encounter", "encounter")
}
