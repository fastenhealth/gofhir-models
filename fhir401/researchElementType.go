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

// ResearchElementType is documented here http://hl7.org/fhir/ValueSet/research-element-type
type ResearchElementType int

const (
	ResearchElementTypePopulation ResearchElementType = iota
	ResearchElementTypeExposure
	ResearchElementTypeOutcome
)

func (code ResearchElementType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ResearchElementType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "population":
		*code = ResearchElementTypePopulation
	case "exposure":
		*code = ResearchElementTypeExposure
	case "outcome":
		*code = ResearchElementTypeOutcome
	default:
		return fmt.Errorf("unknown ResearchElementType code `%s`", s)
	}
	return nil
}
func (code ResearchElementType) String() string {
	return code.Code()
}
func (code ResearchElementType) Code() string {
	switch code {
	case ResearchElementTypePopulation:
		return "population"
	case ResearchElementTypeExposure:
		return "exposure"
	case ResearchElementTypeOutcome:
		return "outcome"
	}
	return "<unknown>"
}
func (code ResearchElementType) Display() string {
	switch code {
	case ResearchElementTypePopulation:
		return "Population"
	case ResearchElementTypeExposure:
		return "Exposure"
	case ResearchElementTypeOutcome:
		return "Outcome"
	}
	return "<unknown>"
}
func (code ResearchElementType) Definition() string {
	switch code {
	case ResearchElementTypePopulation:
		return "The element defines the population that forms the basis for research."
	case ResearchElementTypeExposure:
		return "The element defines an exposure within the population that is being researched."
	case ResearchElementTypeOutcome:
		return "The element defines an outcome within the population that is being researched."
	}
	return "<unknown>"
}
