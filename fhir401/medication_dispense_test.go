package fhir401

import "testing"

func TestMedicationDispenseJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[MedicationDispense](t, "MedicationDispense", "medication_dispense")
}
