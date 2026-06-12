package fhir401

import "testing"

func TestDiagnosticReportJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[DiagnosticReport](t, "DiagnosticReport", "diagnostic_report")
}
