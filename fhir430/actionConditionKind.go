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

// ActionConditionKind is documented here http://hl7.org/fhir/ValueSet/action-condition-kind
type ActionConditionKind int

const (
	ActionConditionKindApplicability ActionConditionKind = iota
	ActionConditionKindStart
	ActionConditionKindStop
)

func (code ActionConditionKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ActionConditionKind) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "applicability":
		*code = ActionConditionKindApplicability
	case "start":
		*code = ActionConditionKindStart
	case "stop":
		*code = ActionConditionKindStop
	default:
		return fmt.Errorf("unknown ActionConditionKind code `%s`", s)
	}
	return nil
}
func (code ActionConditionKind) String() string {
	return code.Code()
}
func (code ActionConditionKind) Code() string {
	switch code {
	case ActionConditionKindApplicability:
		return "applicability"
	case ActionConditionKindStart:
		return "start"
	case ActionConditionKindStop:
		return "stop"
	}
	return "<unknown>"
}
func (code ActionConditionKind) Display() string {
	switch code {
	case ActionConditionKindApplicability:
		return "Applicability"
	case ActionConditionKindStart:
		return "Start"
	case ActionConditionKindStop:
		return "Stop"
	}
	return "<unknown>"
}
func (code ActionConditionKind) Definition() string {
	switch code {
	case ActionConditionKindApplicability:
		return "The condition describes whether or not a given action is applicable."
	case ActionConditionKindStart:
		return "The condition is a starting condition for the action."
	case ActionConditionKindStop:
		return "The condition is a stop, or exit condition for the action."
	}
	return "<unknown>"
}
