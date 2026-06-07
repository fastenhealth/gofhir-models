package fhir401

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type testResource interface {
	ResourceRef() (string, *string)
}

func requireResourceJSONRoundTrip[T testResource](t *testing.T, resourceType string, folder string) {
	t.Helper()

	resourceFiles, err := filepath.Glob(filepath.Join("testdata", folder, "*.json"))
	require.NoError(t, err)
	require.NotEmpty(t, resourceFiles)
	require.LessOrEqual(t, len(resourceFiles), 25)

	for _, resourceFile := range resourceFiles {
		t.Run(strings.TrimSuffix(filepath.Base(resourceFile), ".json"), func(t *testing.T) {
			resourceData, err := os.ReadFile(resourceFile)
			require.NoError(t, err)

			var resource T
			require.NoError(t, json.Unmarshal(resourceData, &resource))

			actualResourceType, resourceID := resource.ResourceRef()
			require.Equal(t, resourceType, actualResourceType)
			require.NotNil(t, resourceID)

			marshaledResource, err := json.Marshal(resource)
			require.NoError(t, err)

			requireJSONFieldsPresent(t, resourceData, marshaledResource)

			var unmarshaledResource T
			require.NoError(t, json.Unmarshal(marshaledResource, &unmarshaledResource))

			remarshaledResource, err := json.Marshal(unmarshaledResource)
			require.NoError(t, err)
			require.JSONEq(t, string(marshaledResource), string(remarshaledResource))
		})
	}
}
