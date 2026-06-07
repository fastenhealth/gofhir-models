package fhir401

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPatientJSONRoundTrip(t *testing.T) {
	patientFiles, err := filepath.Glob("testdata/patient/*.json")
	require.NoError(t, err)
	require.NotEmpty(t, patientFiles)

	for _, patientFile := range patientFiles {
		t.Run(strings.TrimSuffix(filepath.Base(patientFile), ".json"), func(t *testing.T) {
			patientData, err := os.ReadFile(patientFile)
			require.NoError(t, err)

			var patientResource Patient
			require.NoError(t, json.Unmarshal(patientData, &patientResource))
			resourceType, resourceID := patientResource.ResourceRef()
			require.Equal(t, "Patient", resourceType)

			require.NotNil(t, resourceID)

			marshaledPatient, err := json.Marshal(patientResource)
			require.NoError(t, err)

			t.Logf("Original JSON:\n%s\nMarshaled JSON:\n%s\n", string(patientData), string(marshaledPatient))

			requireJSONFieldsPresent(t, patientData, marshaledPatient)
			require.Equal(t, canonicalJSON(t, patientData), canonicalJSON(t, marshaledPatient))

			var unmarshaledPatient Patient
			require.NoError(t, json.Unmarshal(marshaledPatient, &unmarshaledPatient))

			remarshaledPatient, err := json.Marshal(unmarshaledPatient)
			require.NoError(t, err)
			require.JSONEq(t, string(marshaledPatient), string(remarshaledPatient))
		})
	}
}

func requireJSONFieldsPresent(t *testing.T, originalJSON []byte, actualJSON []byte) {
	t.Helper()

	var original any
	require.NoError(t, json.Unmarshal(originalJSON, &original))

	var actual any
	require.NoError(t, json.Unmarshal(actualJSON, &actual))

	require.Empty(t, missingJSONFields(original, actual, "$"))
}

func missingJSONFields(original any, actual any, path string) []string {
	switch originalValue := original.(type) {
	case map[string]any:
		actualValue, ok := actual.(map[string]any)
		if !ok {
			return []string{path}
		}

		var missing []string
		for key, originalField := range originalValue {
			fieldPath := path + "." + key
			actualField, ok := actualValue[key]
			if !ok {
				missing = append(missing, fieldPath)
				continue
			}
			missing = append(missing, missingJSONFields(originalField, actualField, fieldPath)...)
		}
		return missing
	case []any:
		actualValue, ok := actual.([]any)
		if !ok {
			return []string{path}
		}

		var missing []string
		for i, originalField := range originalValue {
			fieldPath := fmt.Sprintf("%s[%d]", path, i)
			if i >= len(actualValue) {
				missing = append(missing, fieldPath)
				continue
			}
			missing = append(missing, missingJSONFields(originalField, actualValue[i], fieldPath)...)
		}
		return missing
	default:
		return nil
	}
}

func canonicalJSON(t *testing.T, data []byte) string {
	t.Helper()

	var value any
	require.NoError(t, json.Unmarshal(data, &value))

	canonical, err := json.Marshal(value)
	require.NoError(t, err)
	return string(canonical)
}
