package fhir401

import "testing"

func TestOrganizationJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[Organization](t, "Organization", "organization")
}
