package fhir401

import "testing"

func TestCareTeamJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[CareTeam](t, "CareTeam", "care_team")
}
