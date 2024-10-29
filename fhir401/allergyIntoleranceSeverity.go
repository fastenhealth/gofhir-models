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

package fhir401

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// AllergyIntoleranceSeverity is documented here http://hl7.org/fhir/ValueSet/reaction-event-severity
type AllergyIntoleranceSeverity int

const (
	AllergyIntoleranceSeverityMild AllergyIntoleranceSeverity = iota
	AllergyIntoleranceSeverityModerate
	AllergyIntoleranceSeveritySevere
)

func (code AllergyIntoleranceSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AllergyIntoleranceSeverity) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "mild":
		*code = AllergyIntoleranceSeverityMild
	case "moderate":
		*code = AllergyIntoleranceSeverityModerate
	case "severe":
		*code = AllergyIntoleranceSeveritySevere
	default:
		return fmt.Errorf("unknown AllergyIntoleranceSeverity code `%s`", s)
	}
	return nil
}
func (code AllergyIntoleranceSeverity) String() string {
	return code.Code()
}
func (code AllergyIntoleranceSeverity) Code() string {
	switch code {
	case AllergyIntoleranceSeverityMild:
		return "mild"
	case AllergyIntoleranceSeverityModerate:
		return "moderate"
	case AllergyIntoleranceSeveritySevere:
		return "severe"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceSeverity) Display() string {
	switch code {
	case AllergyIntoleranceSeverityMild:
		return "Mild"
	case AllergyIntoleranceSeverityModerate:
		return "Moderate"
	case AllergyIntoleranceSeveritySevere:
		return "Severe"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceSeverity) Definition() string {
	switch code {
	case AllergyIntoleranceSeverityMild:
		return "Causes mild physiological effects."
	case AllergyIntoleranceSeverityModerate:
		return "Causes moderate physiological effects."
	case AllergyIntoleranceSeveritySevere:
		return "Causes severe physiological effects."
	}
	return "<unknown>"
}
