package fhir401

import "testing"

func TestQuestionnaireResponseJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[QuestionnaireResponse](t, "QuestionnaireResponse", "questionnaire_response")
}
