package fhir401

import "testing"

func TestProcedureJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[Procedure](t, "Procedure", "procedure")
}
