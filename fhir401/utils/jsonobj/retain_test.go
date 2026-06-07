package jsonobj

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type S struct {
	raw Retain

	Name string `json:"name,omitempty"`
}

func (s *S) UnmarshalJSON(data []byte) error {
	return s.raw.FromJSON(data, s)
}

func (s *S) MarshalJSON() ([]byte, error) {
	return s.raw.ToJSON(s)
}

func TestRetain_RoundTrip(t *testing.T) {
	tests := []struct {
		name   string
		json   string
		verify func(testing.TB, S)

		update         func(*S)
		wantUpdateJSON string
	}{
		{
			name: "empty",
			json: `{}`,
			verify: func(t testing.TB, s S) {
				assert.Equal(t, S{}, s)
			},
		},
		{
			name: "only known",
			json: `{"name": "foo"}`,
			verify: func(t testing.TB, s S) {
				assert.Equal(t, "foo", s.Name)
			},

			update: func(s *S) {
				s.Name = "new name"
			},
			wantUpdateJSON: `{"name":"new name"}`,
		},
		{
			name: "known and unknown",
			json: `{"name":"foo", "num":1,"str":"foo","obj":{"k": "v"},"list":[1,2,3]}`,
			verify: func(t testing.TB, s S) {
				assert.Equal(t, "foo", s.Name)
			},

			update: func(s *S) {
				s.Name = "new name"
			},
			wantUpdateJSON: `{"name":"new name", "num":1,"str":"foo","obj":{"k": "v"},"list":[1,2,3]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s S
			err := json.Unmarshal([]byte(tt.json), &s)
			require.NoError(t, err)

			t.Run("verify Unmarshal", func(t *testing.T) {
				tt.verify(t, s)
				assert.JSONEq(t, tt.json, mustMarshal(t, &s))
			})

			if tt.update != nil {
				t.Run("verify Marshal after mutate", func(t *testing.T) {
					tt.update(&s)
					assert.JSONEq(t, tt.wantUpdateJSON, mustMarshal(t, &s))
				})
			}
		})
	}
}

func TestRetain_ToJSON_OmitEmpty(t *testing.T) {
	t.Run("primitives", func(t *testing.T) {
		var s struct {
			raw Retain

			Str  string `json:",omitempty"`
			Int  int    `json:",omitempty"`
			Bool bool   `json:",omitempty"`
		}
		checkToJSON(t, "zero", s)
		s.Str = "str"
		s.Int = 1
		s.Bool = true

		checkToJSON(t, "set", s)
	})

	t.Run("list", func(t *testing.T) {
		var s struct {
			List []string `json:",omitempty"`
		}
		checkToJSON(t, "zero", s)

		s.List = []string{}
		checkToJSON(t, "empty", s)

		s.List = []string{"", "a", "b"}
		checkToJSON(t, "set", s)
	})

	t.Run("map", func(t *testing.T) {
		var s struct {
			Map map[string]string `json:",omitempty"`
		}
		checkToJSON(t, "zero", s)

		s.Map = make(map[string]string)
		checkToJSON(t, "emtpy", s)

		s.Map = map[string]string{"": "emptykey", "emptyval": "", "foo": "bar"}
		checkToJSON(t, "set", s)
	})

	t.Run("array", func(t *testing.T) {
		var s struct {
			Arr0 [0]int `json:",omitempty"`
			Arr1 [1]int `json:",omitempty"`
		}
		checkToJSON(t, "zero", s)

		s.Arr1[0] = 1
		checkToJSON(t, "set", s)
	})
}

func TestRetain_ToJSON_Ignore(t *testing.T) {
	t.Run("ignore", func(t *testing.T) {
		var s struct {
			IgnoreField string `json:"-"`
		}
		checkToJSON(t, "zero", s)

		s.IgnoreField = "str"
		checkToJSON(t, "set", s)
	})

	t.Run("dash name", func(t *testing.T) {
		var s struct {
			DashName string `json:"-,"`
		}
		checkToJSON(t, "zero", s)

		s.DashName = "str"
		checkToJSON(t, "set", s)
	})

	t.Run("dash name and omitempty", func(t *testing.T) {
		var s struct {
			DashOmitEmpty string `json:"-,omitempty"`
		}
		checkToJSON(t, "zero", s)

		s.DashOmitEmpty = "str"
		checkToJSON(t, "set", s)
	})
}

func TestRetain_FromJSON_Types(t *testing.T) {
	type S struct{}
	var s S
	var str string

	tests := []struct {
		name    string
		obj     any
		wantErr string
	}{
		{
			name:    "string",
			obj:     str,
			wantErr: "FromJSON requires a struct pointer, got string",
		},
		{
			name:    "string pointer",
			obj:     &str,
			wantErr: "FromJSON requires a struct pointer, got *string",
		},
		{
			name:    "struct value",
			obj:     s,
			wantErr: "FromJSON requires a struct pointer, got jsonobj.S",
		},
		{
			name: "struct pointer",
			obj:  &s,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r Retain
			err := r.FromJSON([]byte(`{}`), tt.obj)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestRetain_ToJSON_Types(t *testing.T) {
	tests := []struct {
		name    string
		obj     any
		want    string
		wantErr string
	}{
		{
			name:    "string",
			obj:     "str",
			wantErr: "ToJSON requires a struct, got string",
		},
		{
			name:    "string pointer",
			obj:     ptr("str"),
			wantErr: "ToJSON requires a struct, got *string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r Retain
			got, err := r.ToJSON(tt.obj)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestRetainable(t *testing.T) {
	// Retainable ensures the type implements Marshaler/Unmarshaler.
	type base struct {
		json.Marshaler
		json.Unmarshaler
	}

	type Valid struct {
		base

		Name  string
		Name2 string `json:"name"` // different from Name
		Age   int    `json:",omitempty"`

		Ignored func() `json:"-"`
		Dash    string `json:"-,"`
	}

	type DuplicateNameWithTag struct {
		base
		Name      string
		OtherName string `json:"Name"`
	}

	type DuplicateNameTags struct {
		base
		Name1 string `json:"name"`
		Name2 string `json:"name"`
	}

	type UnsupportedStringTag struct {
		base
		Age int `json:",string"`
	}

	type InlineStruct struct {
		Name string
	}

	type UnsupportedInlineTag struct {
		base
		InlineStruct `json:",inline"`
	}

	tests := []struct {
		v interface {
			json.Marshaler
			json.Unmarshaler
		}
		wantErr string
	}{
		{
			v: &Valid{},
		},
		{
			v:       Valid{},
			wantErr: `jsonobj.Valid not Retainable: requires struct pointer`,
		},
		{
			v:       &DuplicateNameWithTag{},
			wantErr: `*jsonobj.DuplicateNameWithTag not Retainable: duplicate JSON field "Name"`,
		},
		{
			v:       &DuplicateNameTags{},
			wantErr: `*jsonobj.DuplicateNameTags not Retainable: duplicate JSON field "name"`,
		},
		{
			v:       &UnsupportedStringTag{},
			wantErr: `*jsonobj.UnsupportedStringTag not Retainable: field "Age" has unsupported tag "string"`,
		},
		{
			v:       &UnsupportedInlineTag{},
			wantErr: `*jsonobj.UnsupportedInlineTag not Retainable: field "InlineStruct" has unsupported tag "inline"`,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%T", tt.v), func(t *testing.T) {
			err := Retainable(tt.v)

			if tt.wantErr == "" {
				assert.NoError(t, err)
				return
			}
			assert.EqualError(t, err, tt.wantErr)

			t.Run("MustRetainable", func(t *testing.T) {
				assert.PanicsWithError(t, tt.wantErr, func() {
					MustRetainable(tt.v)
				})
			})
		})
	}
}

func TestRetain_FromJSON_Errors(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		wantErr string
	}{
		{
			name:    "invalid JSON",
			json:    `{"key"}`,
			wantErr: "invalid character",
		},
		{
			name:    "not object",
			json:    "true",
			wantErr: "cannot unmarshal bool into",
		},
		{
			name:    "wrong field type",
			json:    `{"name": 1}`,
			wantErr: "cannot unmarshal number into",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s S
			err := s.UnmarshalJSON([]byte(tt.json))
			assert.ErrorContains(t, err, tt.wantErr)
		})
	}
}

// checkToJSON marshals the object using ToJSON and compares
// it to marshalling the struct using `json.Marshal`.
func checkToJSON(t *testing.T, name string, obj any) {
	t.Run(name, func(t *testing.T) {
		// Ensure a custom MarshalJSON isn't implemented, as that may
		// call ToJSON, which is what we're trying to test.
		_, ok := obj.(json.Marshaler)
		require.False(t, ok)

		var r Retain
		got, err := r.ToJSON(obj)
		require.NoError(t, err)
		want := mustMarshal(t, obj)
		assert.JSONEq(t, want, string(got))
	})
}

func mustMarshal(t *testing.T, obj any) string {
	b, err := json.Marshal(obj)
	require.NoError(t, err)
	return string(b)
}

func ptr(v any) *any {
	return &v
}
