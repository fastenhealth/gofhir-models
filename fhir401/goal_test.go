package fhir401

import "testing"

func TestGoalJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[Goal](t, "Goal", "goal")
}
