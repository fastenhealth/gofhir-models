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

// ImmunizationStatusCodes is documented here http://hl7.org/fhir/ValueSet/immunization-status
type ImmunizationStatusCodes int

const (
	ImmunizationStatusCodesPreparation ImmunizationStatusCodes = iota
	ImmunizationStatusCodesInProgress
	ImmunizationStatusCodesNotDone
	ImmunizationStatusCodesOnHold
	ImmunizationStatusCodesStopped
	ImmunizationStatusCodesCompleted
	ImmunizationStatusCodesEnteredInError
	ImmunizationStatusCodesUnknown
)

func (code ImmunizationStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ImmunizationStatusCodes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "preparation":
		*code = ImmunizationStatusCodesPreparation
	case "in-progress":
		*code = ImmunizationStatusCodesInProgress
	case "not-done":
		*code = ImmunizationStatusCodesNotDone
	case "on-hold":
		*code = ImmunizationStatusCodesOnHold
	case "stopped":
		*code = ImmunizationStatusCodesStopped
	case "completed":
		*code = ImmunizationStatusCodesCompleted
	case "entered-in-error":
		*code = ImmunizationStatusCodesEnteredInError
	case "unknown":
		*code = ImmunizationStatusCodesUnknown
	default:
		return fmt.Errorf("unknown ImmunizationStatusCodes code `%s`", s)
	}
	return nil
}
func (code ImmunizationStatusCodes) String() string {
	return code.Code()
}
func (code ImmunizationStatusCodes) Code() string {
	switch code {
	case ImmunizationStatusCodesPreparation:
		return "preparation"
	case ImmunizationStatusCodesInProgress:
		return "in-progress"
	case ImmunizationStatusCodesNotDone:
		return "not-done"
	case ImmunizationStatusCodesOnHold:
		return "on-hold"
	case ImmunizationStatusCodesStopped:
		return "stopped"
	case ImmunizationStatusCodesCompleted:
		return "completed"
	case ImmunizationStatusCodesEnteredInError:
		return "entered-in-error"
	case ImmunizationStatusCodesUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code ImmunizationStatusCodes) Display() string {
	switch code {
	case ImmunizationStatusCodesPreparation:
		return "Preparation"
	case ImmunizationStatusCodesInProgress:
		return "In Progress"
	case ImmunizationStatusCodesNotDone:
		return "Not Done"
	case ImmunizationStatusCodesOnHold:
		return "On Hold"
	case ImmunizationStatusCodesStopped:
		return "Stopped"
	case ImmunizationStatusCodesCompleted:
		return "Completed"
	case ImmunizationStatusCodesEnteredInError:
		return "Entered in Error"
	case ImmunizationStatusCodesUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code ImmunizationStatusCodes) Definition() string {
	switch code {
	case ImmunizationStatusCodesPreparation:
		return "The core event has not started yet, but some staging activities have begun (e.g. surgical suite preparation).  Preparation stages may be tracked for billing purposes."
	case ImmunizationStatusCodesInProgress:
		return "The event is currently occurring."
	case ImmunizationStatusCodesNotDone:
		return "The event was terminated prior to any activity beyond preparation.  I.e. The 'main' activity has not yet begun.  The boundary between preparatory and the 'main' activity is context-specific."
	case ImmunizationStatusCodesOnHold:
		return "The event has been temporarily stopped but is expected to resume in the future."
	case ImmunizationStatusCodesStopped:
		return "The event was terminated prior to the full completion of the intended activity but after at least some of the 'main' activity (beyond preparation) has occurred."
	case ImmunizationStatusCodesCompleted:
		return "The event has now concluded."
	case ImmunizationStatusCodesEnteredInError:
		return "This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be \"stopped\" rather than \"entered-in-error\".)."
	case ImmunizationStatusCodesUnknown:
		return "The authoring/source system does not know which of the status values currently applies for this event.  Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply,  but the authoring/source system does not know which."
	}
	return "<unknown>"
}
