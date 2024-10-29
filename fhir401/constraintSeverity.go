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

// ConstraintSeverity is documented here http://hl7.org/fhir/ValueSet/constraint-severity
type ConstraintSeverity int

const (
	ConstraintSeverityError ConstraintSeverity = iota
	ConstraintSeverityWarning
)

func (code ConstraintSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ConstraintSeverity) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "error":
		*code = ConstraintSeverityError
	case "warning":
		*code = ConstraintSeverityWarning
	default:
		return fmt.Errorf("unknown ConstraintSeverity code `%s`", s)
	}
	return nil
}
func (code ConstraintSeverity) String() string {
	return code.Code()
}
func (code ConstraintSeverity) Code() string {
	switch code {
	case ConstraintSeverityError:
		return "error"
	case ConstraintSeverityWarning:
		return "warning"
	}
	return "<unknown>"
}
func (code ConstraintSeverity) Display() string {
	switch code {
	case ConstraintSeverityError:
		return "Error"
	case ConstraintSeverityWarning:
		return "Warning"
	}
	return "<unknown>"
}
func (code ConstraintSeverity) Definition() string {
	switch code {
	case ConstraintSeverityError:
		return "If the constraint is violated, the resource is not conformant."
	case ConstraintSeverityWarning:
		return "If the constraint is violated, the resource is conformant, but it is not necessarily following best practice."
	}
	return "<unknown>"
}
