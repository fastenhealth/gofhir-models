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

// FHIRDeviceStatus is documented here http://hl7.org/fhir/ValueSet/device-status
type FHIRDeviceStatus int

const (
	FHIRDeviceStatusActive FHIRDeviceStatus = iota
	FHIRDeviceStatusInactive
	FHIRDeviceStatusEnteredInError
	FHIRDeviceStatusUnknown
)

func (code FHIRDeviceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *FHIRDeviceStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "active":
		*code = FHIRDeviceStatusActive
	case "inactive":
		*code = FHIRDeviceStatusInactive
	case "entered-in-error":
		*code = FHIRDeviceStatusEnteredInError
	case "unknown":
		*code = FHIRDeviceStatusUnknown
	default:
		return fmt.Errorf("unknown FHIRDeviceStatus code `%s`", s)
	}
	return nil
}
func (code FHIRDeviceStatus) String() string {
	return code.Code()
}
func (code FHIRDeviceStatus) Code() string {
	switch code {
	case FHIRDeviceStatusActive:
		return "active"
	case FHIRDeviceStatusInactive:
		return "inactive"
	case FHIRDeviceStatusEnteredInError:
		return "entered-in-error"
	case FHIRDeviceStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code FHIRDeviceStatus) Display() string {
	switch code {
	case FHIRDeviceStatusActive:
		return "Active"
	case FHIRDeviceStatusInactive:
		return "Inactive"
	case FHIRDeviceStatusEnteredInError:
		return "Entered in Error"
	case FHIRDeviceStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code FHIRDeviceStatus) Definition() string {
	switch code {
	case FHIRDeviceStatusActive:
		return "The device is available for use.  Note: For *implanted devices*  this means that the device is implanted in the patient."
	case FHIRDeviceStatusInactive:
		return "The device is no longer available for use (e.g. lost, expired, damaged).  Note: For *implanted devices*  this means that the device has been removed from the patient."
	case FHIRDeviceStatusEnteredInError:
		return "The device was entered in error and voided."
	case FHIRDeviceStatusUnknown:
		return "The status of the device has not been determined."
	}
	return "<unknown>"
}
