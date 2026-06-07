package fhir401

import "testing"

func TestMedicationRequestJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[MedicationRequest](t, "MedicationRequest", "medication_request")
}
