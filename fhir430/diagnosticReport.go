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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// DiagnosticReport is documented here http://hl7.org/fhir/StructureDefinition/DiagnosticReport
// The findings and interpretation of diagnostic  tests performed on patients, groups of patients, devices, and locations, and/or specimens derived from these. The report includes clinical context such as requesting and provider information, and some mix of atomic results, images, textual and coded interpretations, and formatted representation of diagnostic reports.
type DiagnosticReport struct {
	Id                 *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	Contained          []json.RawMessage       `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension          []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn            []Reference             `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Status             DiagnosticReportStatus  `bson:"status" json:"status"`
	Category           []CodeableConcept       `bson:"category,omitempty" json:"category,omitempty"`
	Code               CodeableConcept         `bson:"code" json:"code"`
	Subject            *Reference              `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter          *Reference              `bson:"encounter,omitempty" json:"encounter,omitempty"`
	EffectiveDateTime  *string                 `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod    *Period                 `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Issued             *string                 `bson:"issued,omitempty" json:"issued,omitempty"`
	Performer          []Reference             `bson:"performer,omitempty" json:"performer,omitempty"`
	ResultsInterpreter []Reference             `bson:"resultsInterpreter,omitempty" json:"resultsInterpreter,omitempty"`
	Specimen           []Reference             `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Result             []Reference             `bson:"result,omitempty" json:"result,omitempty"`
	ImagingStudy       []Reference             `bson:"imagingStudy,omitempty" json:"imagingStudy,omitempty"`
	Media              []DiagnosticReportMedia `bson:"media,omitempty" json:"media,omitempty"`
	Conclusion         *string                 `bson:"conclusion,omitempty" json:"conclusion,omitempty"`
	ConclusionCode     []CodeableConcept       `bson:"conclusionCode,omitempty" json:"conclusionCode,omitempty"`
	PresentedForm      []Attachment            `bson:"presentedForm,omitempty" json:"presentedForm,omitempty"`
}

// A list of key images associated with this report. The images are generally created during the diagnostic process, and may be directly of the patient, or of treated specimens (i.e. slides of interest).
type DiagnosticReportMedia struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Comment           *string     `bson:"comment,omitempty" json:"comment,omitempty"`
	Link              Reference   `bson:"link" json:"link"`
}

// This function returns resource reference information
func (r DiagnosticReport) ResourceRef() (string, *string) {
	return "DiagnosticReport", r.Id
}

// This function returns resource reference information
func (r DiagnosticReport) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherDiagnosticReport DiagnosticReport

// MarshalJSON marshals the given DiagnosticReport as JSON into a byte slice
func (r DiagnosticReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDiagnosticReport
		ResourceType string `json:"resourceType"`
	}{
		OtherDiagnosticReport: OtherDiagnosticReport(r),
		ResourceType:          "DiagnosticReport",
	})
}

// UnmarshalDiagnosticReport unmarshals a DiagnosticReport.
func UnmarshalDiagnosticReport(b []byte) (DiagnosticReport, error) {
	var diagnosticReport DiagnosticReport
	if err := json.Unmarshal(b, &diagnosticReport); err != nil {
		return diagnosticReport, err
	}
	return diagnosticReport, nil
}
