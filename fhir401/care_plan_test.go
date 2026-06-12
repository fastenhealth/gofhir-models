package fhir401

import "testing"

func TestCarePlanJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[CarePlan](t, "CarePlan", "care_plan")
}
