package fhir401

import "testing"

func TestConditionJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[Condition](t, "Condition", "condition")
}
