package fhir401

import "testing"

func TestDeviceJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[Device](t, "Device", "device")
}
