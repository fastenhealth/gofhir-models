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

// RequestStatus is documented here http://hl7.org/fhir/ValueSet/request-status
type RequestStatus int

const (
	RequestStatusDraft RequestStatus = iota
	RequestStatusActive
	RequestStatusOnHold
	RequestStatusRevoked
	RequestStatusCompleted
	RequestStatusEnteredInError
	RequestStatusUnknown
)

func (code RequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *RequestStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "draft":
		*code = RequestStatusDraft
	case "active":
		*code = RequestStatusActive
	case "on-hold":
		*code = RequestStatusOnHold
	case "revoked":
		*code = RequestStatusRevoked
	case "completed":
		*code = RequestStatusCompleted
	case "entered-in-error":
		*code = RequestStatusEnteredInError
	case "unknown":
		*code = RequestStatusUnknown
	default:
		return fmt.Errorf("unknown RequestStatus code `%s`", s)
	}
	return nil
}
func (code RequestStatus) String() string {
	return code.Code()
}
func (code RequestStatus) Code() string {
	switch code {
	case RequestStatusDraft:
		return "draft"
	case RequestStatusActive:
		return "active"
	case RequestStatusOnHold:
		return "on-hold"
	case RequestStatusRevoked:
		return "revoked"
	case RequestStatusCompleted:
		return "completed"
	case RequestStatusEnteredInError:
		return "entered-in-error"
	case RequestStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code RequestStatus) Display() string {
	switch code {
	case RequestStatusDraft:
		return "Draft"
	case RequestStatusActive:
		return "Active"
	case RequestStatusOnHold:
		return "On Hold"
	case RequestStatusRevoked:
		return "Revoked"
	case RequestStatusCompleted:
		return "Completed"
	case RequestStatusEnteredInError:
		return "Entered in Error"
	case RequestStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code RequestStatus) Definition() string {
	switch code {
	case RequestStatusDraft:
		return "The request has been created but is not yet complete or ready for action."
	case RequestStatusActive:
		return "The request is in force and ready to be acted upon."
	case RequestStatusOnHold:
		return "The request (and any implicit authorization to act) has been temporarily withdrawn but is expected to resume in the future."
	case RequestStatusRevoked:
		return "The request (and any implicit authorization to act) has been terminated prior to the known full completion of the intended actions.  No further activity should occur."
	case RequestStatusCompleted:
		return "The activity described by the request has been fully performed.  No further activity will occur."
	case RequestStatusEnteredInError:
		return "This request should never have existed and should be considered 'void'.  (It is possible that real-world decisions were based on it.  If real-world activity has occurred, the status should be \"revoked\" rather than \"entered-in-error\".)."
	case RequestStatusUnknown:
		return "The authoring/source system does not know which of the status values currently applies for this request.  Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply,  but the authoring/source system does not know which."
	}
	return "<unknown>"
}
