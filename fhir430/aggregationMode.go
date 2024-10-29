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

// AggregationMode is documented here http://hl7.org/fhir/ValueSet/resource-aggregation-mode
type AggregationMode int

const (
	AggregationModeContained AggregationMode = iota
	AggregationModeReferenced
	AggregationModeBundled
)

func (code AggregationMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AggregationMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "contained":
		*code = AggregationModeContained
	case "referenced":
		*code = AggregationModeReferenced
	case "bundled":
		*code = AggregationModeBundled
	default:
		return fmt.Errorf("unknown AggregationMode code `%s`", s)
	}
	return nil
}
func (code AggregationMode) String() string {
	return code.Code()
}
func (code AggregationMode) Code() string {
	switch code {
	case AggregationModeContained:
		return "contained"
	case AggregationModeReferenced:
		return "referenced"
	case AggregationModeBundled:
		return "bundled"
	}
	return "<unknown>"
}
func (code AggregationMode) Display() string {
	switch code {
	case AggregationModeContained:
		return "Contained"
	case AggregationModeReferenced:
		return "Referenced"
	case AggregationModeBundled:
		return "Bundled"
	}
	return "<unknown>"
}
func (code AggregationMode) Definition() string {
	switch code {
	case AggregationModeContained:
		return "The reference is a local reference to a contained resource."
	case AggregationModeReferenced:
		return "The reference to a resource that has to be resolved externally to the resource that includes the reference."
	case AggregationModeBundled:
		return "The resource the reference points to will be found in the same bundle as the resource that includes the reference."
	}
	return "<unknown>"
}
