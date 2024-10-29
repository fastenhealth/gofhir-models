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

// SupplyRequestStatus is documented here http://hl7.org/fhir/ValueSet/supplyrequest-status
type SupplyRequestStatus int

const (
	SupplyRequestStatusDraft SupplyRequestStatus = iota
	SupplyRequestStatusActive
	SupplyRequestStatusSuspended
	SupplyRequestStatusCancelled
	SupplyRequestStatusCompleted
	SupplyRequestStatusEnteredInError
	SupplyRequestStatusUnknown
)

func (code SupplyRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SupplyRequestStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "draft":
		*code = SupplyRequestStatusDraft
	case "active":
		*code = SupplyRequestStatusActive
	case "suspended":
		*code = SupplyRequestStatusSuspended
	case "cancelled":
		*code = SupplyRequestStatusCancelled
	case "completed":
		*code = SupplyRequestStatusCompleted
	case "entered-in-error":
		*code = SupplyRequestStatusEnteredInError
	case "unknown":
		*code = SupplyRequestStatusUnknown
	default:
		return fmt.Errorf("unknown SupplyRequestStatus code `%s`", s)
	}
	return nil
}
func (code SupplyRequestStatus) String() string {
	return code.Code()
}
func (code SupplyRequestStatus) Code() string {
	switch code {
	case SupplyRequestStatusDraft:
		return "draft"
	case SupplyRequestStatusActive:
		return "active"
	case SupplyRequestStatusSuspended:
		return "suspended"
	case SupplyRequestStatusCancelled:
		return "cancelled"
	case SupplyRequestStatusCompleted:
		return "completed"
	case SupplyRequestStatusEnteredInError:
		return "entered-in-error"
	case SupplyRequestStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code SupplyRequestStatus) Display() string {
	switch code {
	case SupplyRequestStatusDraft:
		return "Draft"
	case SupplyRequestStatusActive:
		return "Active"
	case SupplyRequestStatusSuspended:
		return "Suspended"
	case SupplyRequestStatusCancelled:
		return "Cancelled"
	case SupplyRequestStatusCompleted:
		return "Completed"
	case SupplyRequestStatusEnteredInError:
		return "Entered in Error"
	case SupplyRequestStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code SupplyRequestStatus) Definition() string {
	switch code {
	case SupplyRequestStatusDraft:
		return "The request has been created but is not yet complete or ready for action."
	case SupplyRequestStatusActive:
		return "The request is ready to be acted upon."
	case SupplyRequestStatusSuspended:
		return "The authorization/request to act has been temporarily withdrawn but is expected to resume in the future."
	case SupplyRequestStatusCancelled:
		return "The authorization/request to act has been terminated prior to the full completion of the intended actions.  No further activity should occur."
	case SupplyRequestStatusCompleted:
		return "Activity against the request has been sufficiently completed to the satisfaction of the requester."
	case SupplyRequestStatusEnteredInError:
		return "This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be \"cancelled\" rather than \"entered-in-error\".)."
	case SupplyRequestStatusUnknown:
		return "The authoring/source system does not know which of the status values currently applies for this observation. Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply, but the authoring/source system does not know which."
	}
	return "<unknown>"
}
