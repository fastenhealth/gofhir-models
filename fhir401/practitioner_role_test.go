package fhir401

import "testing"

func TestPractitionerRoleJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[PractitionerRole](t, "PractitionerRole", "practitioner_role")
}
