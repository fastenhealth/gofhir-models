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

// DiagnosticReportStatus is documented here http://hl7.org/fhir/ValueSet/diagnostic-report-status
type DiagnosticReportStatus int

const (
	DiagnosticReportStatusRegistered DiagnosticReportStatus = iota
	DiagnosticReportStatusPartial
	DiagnosticReportStatusPreliminary
	DiagnosticReportStatusFinal
	DiagnosticReportStatusAmended
	DiagnosticReportStatusCorrected
	DiagnosticReportStatusAppended
	DiagnosticReportStatusCancelled
	DiagnosticReportStatusEnteredInError
	DiagnosticReportStatusUnknown
)

func (code DiagnosticReportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DiagnosticReportStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "registered":
		*code = DiagnosticReportStatusRegistered
	case "partial":
		*code = DiagnosticReportStatusPartial
	case "preliminary":
		*code = DiagnosticReportStatusPreliminary
	case "final":
		*code = DiagnosticReportStatusFinal
	case "amended":
		*code = DiagnosticReportStatusAmended
	case "corrected":
		*code = DiagnosticReportStatusCorrected
	case "appended":
		*code = DiagnosticReportStatusAppended
	case "cancelled":
		*code = DiagnosticReportStatusCancelled
	case "entered-in-error":
		*code = DiagnosticReportStatusEnteredInError
	case "unknown":
		*code = DiagnosticReportStatusUnknown
	default:
		return fmt.Errorf("unknown DiagnosticReportStatus code `%s`", s)
	}
	return nil
}
func (code DiagnosticReportStatus) String() string {
	return code.Code()
}
func (code DiagnosticReportStatus) Code() string {
	switch code {
	case DiagnosticReportStatusRegistered:
		return "registered"
	case DiagnosticReportStatusPartial:
		return "partial"
	case DiagnosticReportStatusPreliminary:
		return "preliminary"
	case DiagnosticReportStatusFinal:
		return "final"
	case DiagnosticReportStatusAmended:
		return "amended"
	case DiagnosticReportStatusCorrected:
		return "corrected"
	case DiagnosticReportStatusAppended:
		return "appended"
	case DiagnosticReportStatusCancelled:
		return "cancelled"
	case DiagnosticReportStatusEnteredInError:
		return "entered-in-error"
	case DiagnosticReportStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code DiagnosticReportStatus) Display() string {
	switch code {
	case DiagnosticReportStatusRegistered:
		return "Registered"
	case DiagnosticReportStatusPartial:
		return "Partial"
	case DiagnosticReportStatusPreliminary:
		return "Preliminary"
	case DiagnosticReportStatusFinal:
		return "Final"
	case DiagnosticReportStatusAmended:
		return "Amended"
	case DiagnosticReportStatusCorrected:
		return "Corrected"
	case DiagnosticReportStatusAppended:
		return "Appended"
	case DiagnosticReportStatusCancelled:
		return "Cancelled"
	case DiagnosticReportStatusEnteredInError:
		return "Entered in Error"
	case DiagnosticReportStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code DiagnosticReportStatus) Definition() string {
	switch code {
	case DiagnosticReportStatusRegistered:
		return "The existence of the report is registered, but there is nothing yet available."
	case DiagnosticReportStatusPartial:
		return "This is a partial (e.g. initial, interim or preliminary) report: data in the report may be incomplete or unverified."
	case DiagnosticReportStatusPreliminary:
		return "Verified early results are available, but not all  results are final."
	case DiagnosticReportStatusFinal:
		return "The report is complete and verified by an authorized person."
	case DiagnosticReportStatusAmended:
		return "Subsequent to being final, the report has been modified.  This includes any change in the results, diagnosis, narrative text, or other content of a report that has been issued."
	case DiagnosticReportStatusCorrected:
		return "Subsequent to being final, the report has been modified  to correct an error in the report or referenced results."
	case DiagnosticReportStatusAppended:
		return "Subsequent to being final, the report has been modified by adding new content. The existing content is unchanged."
	case DiagnosticReportStatusCancelled:
		return "The report is unavailable because the measurement was not started or not completed (also sometimes called \"aborted\")."
	case DiagnosticReportStatusEnteredInError:
		return "The report has been withdrawn following a previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be \"cancelled\" rather than \"entered-in-error\".)."
	case DiagnosticReportStatusUnknown:
		return "The authoring/source system does not know which of the status values currently applies for this observation. Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply, but the authoring/source system does not know which."
	}
	return "<unknown>"
}
