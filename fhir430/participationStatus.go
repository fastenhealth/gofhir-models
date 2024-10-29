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

// ParticipationStatus is documented here http://hl7.org/fhir/ValueSet/participationstatus
type ParticipationStatus int

const (
	ParticipationStatusAccepted ParticipationStatus = iota
	ParticipationStatusDeclined
	ParticipationStatusTentative
	ParticipationStatusNeedsAction
)

func (code ParticipationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ParticipationStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "accepted":
		*code = ParticipationStatusAccepted
	case "declined":
		*code = ParticipationStatusDeclined
	case "tentative":
		*code = ParticipationStatusTentative
	case "needs-action":
		*code = ParticipationStatusNeedsAction
	default:
		return fmt.Errorf("unknown ParticipationStatus code `%s`", s)
	}
	return nil
}
func (code ParticipationStatus) String() string {
	return code.Code()
}
func (code ParticipationStatus) Code() string {
	switch code {
	case ParticipationStatusAccepted:
		return "accepted"
	case ParticipationStatusDeclined:
		return "declined"
	case ParticipationStatusTentative:
		return "tentative"
	case ParticipationStatusNeedsAction:
		return "needs-action"
	}
	return "<unknown>"
}
func (code ParticipationStatus) Display() string {
	switch code {
	case ParticipationStatusAccepted:
		return "Accepted"
	case ParticipationStatusDeclined:
		return "Declined"
	case ParticipationStatusTentative:
		return "Tentative"
	case ParticipationStatusNeedsAction:
		return "Needs Action"
	}
	return "<unknown>"
}
func (code ParticipationStatus) Definition() string {
	switch code {
	case ParticipationStatusAccepted:
		return "The participant has accepted the appointment."
	case ParticipationStatusDeclined:
		return "The participant has declined the appointment and will not participate in the appointment."
	case ParticipationStatusTentative:
		return "The participant has  tentatively accepted the appointment. This could be automatically created by a system and requires further processing before it can be accepted. There is no commitment that attendance will occur."
	case ParticipationStatusNeedsAction:
		return "The participant needs to indicate if they accept the appointment by changing this status to one of the other statuses."
	}
	return "<unknown>"
}
