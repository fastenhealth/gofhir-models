package fhir401

import "testing"

func TestDocumentReferenceJSONRoundTrip(t *testing.T) {
	requireResourceJSONRoundTrip[DocumentReference](t, "DocumentReference", "document_reference")
}
