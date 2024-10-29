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

// ConsentProvisionType is documented here http://hl7.org/fhir/ValueSet/consent-provision-type
type ConsentProvisionType int

const (
	ConsentProvisionTypeDeny ConsentProvisionType = iota
	ConsentProvisionTypePermit
)

func (code ConsentProvisionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ConsentProvisionType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "deny":
		*code = ConsentProvisionTypeDeny
	case "permit":
		*code = ConsentProvisionTypePermit
	default:
		return fmt.Errorf("unknown ConsentProvisionType code `%s`", s)
	}
	return nil
}
func (code ConsentProvisionType) String() string {
	return code.Code()
}
func (code ConsentProvisionType) Code() string {
	switch code {
	case ConsentProvisionTypeDeny:
		return "deny"
	case ConsentProvisionTypePermit:
		return "permit"
	}
	return "<unknown>"
}
func (code ConsentProvisionType) Display() string {
	switch code {
	case ConsentProvisionTypeDeny:
		return "Opt Out"
	case ConsentProvisionTypePermit:
		return "Opt In"
	}
	return "<unknown>"
}
func (code ConsentProvisionType) Definition() string {
	switch code {
	case ConsentProvisionTypeDeny:
		return "Consent is denied for actions meeting these rules."
	case ConsentProvisionTypePermit:
		return "Consent is provided for actions meeting these rules."
	}
	return "<unknown>"
}
