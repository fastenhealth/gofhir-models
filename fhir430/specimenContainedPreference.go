// Copyright 2022 - Fasten Health
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// SpecimenContainedPreference is documented here http://hl7.org/fhir/ValueSet/specimen-contained-preference
type SpecimenContainedPreference int

const (
	SpecimenContainedPreferencePreferred SpecimenContainedPreference = iota
	SpecimenContainedPreferenceAlternate
)

func (code SpecimenContainedPreference) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SpecimenContainedPreference) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "preferred":
		*code = SpecimenContainedPreferencePreferred
	case "alternate":
		*code = SpecimenContainedPreferenceAlternate
	default:
		return fmt.Errorf("unknown SpecimenContainedPreference code `%s`", s)
	}
	return nil
}
func (code SpecimenContainedPreference) String() string {
	return code.Code()
}
func (code SpecimenContainedPreference) Code() string {
	switch code {
	case SpecimenContainedPreferencePreferred:
		return "preferred"
	case SpecimenContainedPreferenceAlternate:
		return "alternate"
	}
	return "<unknown>"
}
func (code SpecimenContainedPreference) Display() string {
	switch code {
	case SpecimenContainedPreferencePreferred:
		return "Preferred"
	case SpecimenContainedPreferenceAlternate:
		return "Alternate"
	}
	return "<unknown>"
}
func (code SpecimenContainedPreference) Definition() string {
	switch code {
	case SpecimenContainedPreferencePreferred:
		return "This type of contained specimen is preferred to collect this kind of specimen."
	case SpecimenContainedPreferenceAlternate:
		return "This type of conditioned specimen is an alternate."
	}
	return "<unknown>"
}
