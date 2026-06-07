package fhir401

import "testing"

func TestMedicationJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[Medication](t, "Medication", "medication")
}
