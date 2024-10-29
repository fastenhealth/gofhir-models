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

// EncounterStatus is documented here http://hl7.org/fhir/ValueSet/encounter-status
type EncounterStatus int

const (
	EncounterStatusPlanned EncounterStatus = iota
	EncounterStatusArrived
	EncounterStatusTriaged
	EncounterStatusInProgress
	EncounterStatusOnleave
	EncounterStatusFinished
	EncounterStatusCancelled
	EncounterStatusEnteredInError
	EncounterStatusUnknown
)

func (code EncounterStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *EncounterStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "planned":
		*code = EncounterStatusPlanned
	case "arrived":
		*code = EncounterStatusArrived
	case "triaged":
		*code = EncounterStatusTriaged
	case "in-progress":
		*code = EncounterStatusInProgress
	case "onleave":
		*code = EncounterStatusOnleave
	case "finished":
		*code = EncounterStatusFinished
	case "cancelled":
		*code = EncounterStatusCancelled
	case "entered-in-error":
		*code = EncounterStatusEnteredInError
	case "unknown":
		*code = EncounterStatusUnknown
	default:
		return fmt.Errorf("unknown EncounterStatus code `%s`", s)
	}
	return nil
}
func (code EncounterStatus) String() string {
	return code.Code()
}
func (code EncounterStatus) Code() string {
	switch code {
	case EncounterStatusPlanned:
		return "planned"
	case EncounterStatusArrived:
		return "arrived"
	case EncounterStatusTriaged:
		return "triaged"
	case EncounterStatusInProgress:
		return "in-progress"
	case EncounterStatusOnleave:
		return "onleave"
	case EncounterStatusFinished:
		return "finished"
	case EncounterStatusCancelled:
		return "cancelled"
	case EncounterStatusEnteredInError:
		return "entered-in-error"
	case EncounterStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code EncounterStatus) Display() string {
	switch code {
	case EncounterStatusPlanned:
		return "Planned"
	case EncounterStatusArrived:
		return "Arrived"
	case EncounterStatusTriaged:
		return "Triaged"
	case EncounterStatusInProgress:
		return "In Progress"
	case EncounterStatusOnleave:
		return "On Leave"
	case EncounterStatusFinished:
		return "Finished"
	case EncounterStatusCancelled:
		return "Cancelled"
	case EncounterStatusEnteredInError:
		return "Entered in Error"
	case EncounterStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code EncounterStatus) Definition() string {
	switch code {
	case EncounterStatusPlanned:
		return "The Encounter has not yet started."
	case EncounterStatusArrived:
		return "The Patient is present for the encounter, however is not currently meeting with a practitioner."
	case EncounterStatusTriaged:
		return "The patient has been assessed for the priority of their treatment based on the severity of their condition."
	case EncounterStatusInProgress:
		return "The Encounter has begun and the patient is present / the practitioner and the patient are meeting."
	case EncounterStatusOnleave:
		return "The Encounter has begun, but the patient is temporarily on leave."
	case EncounterStatusFinished:
		return "The Encounter has ended."
	case EncounterStatusCancelled:
		return "The Encounter has ended before it has begun."
	case EncounterStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	case EncounterStatusUnknown:
		return "The encounter status is unknown. Note that \"unknown\" is a value of last resort and every attempt should be made to provide a meaningful value other than \"unknown\"."
	}
	return "<unknown>"
}
