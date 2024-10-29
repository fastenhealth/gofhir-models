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

// Use is documented here http://hl7.org/fhir/ValueSet/claim-use
type Use int

const (
	UseClaim Use = iota
	UsePreauthorization
	UsePredetermination
)

func (code Use) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *Use) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "claim":
		*code = UseClaim
	case "preauthorization":
		*code = UsePreauthorization
	case "predetermination":
		*code = UsePredetermination
	default:
		return fmt.Errorf("unknown Use code `%s`", s)
	}
	return nil
}
func (code Use) String() string {
	return code.Code()
}
func (code Use) Code() string {
	switch code {
	case UseClaim:
		return "claim"
	case UsePreauthorization:
		return "preauthorization"
	case UsePredetermination:
		return "predetermination"
	}
	return "<unknown>"
}
func (code Use) Display() string {
	switch code {
	case UseClaim:
		return "Claim"
	case UsePreauthorization:
		return "Preauthorization"
	case UsePredetermination:
		return "Predetermination"
	}
	return "<unknown>"
}
func (code Use) Definition() string {
	switch code {
	case UseClaim:
		return "The treatment is complete and this represents a Claim for the services."
	case UsePreauthorization:
		return "The treatment is proposed and this represents a Pre-authorization for the services."
	case UsePredetermination:
		return "The treatment is proposed and this represents a Pre-determination for the services."
	}
	return "<unknown>"
}
