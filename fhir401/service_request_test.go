package fhir401

import "testing"

func TestServiceRequestJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[ServiceRequest](t, "ServiceRequest", "service_request")
}
