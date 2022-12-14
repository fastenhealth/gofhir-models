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

// Observation is documented here http://hl7.org/fhir/StructureDefinition/Observation
// Measurements and simple assertions made about a patient, device or other subject.
type Observation struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn           []Reference                 `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf            []Reference                 `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status            ObservationStatus           `bson:"status" json:"status"`
	Category          []CodeableConcept           `bson:"category,omitempty" json:"category,omitempty"`
	Code              CodeableConcept             `bson:"code" json:"code"`
	Subject           *Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Focus             []Reference                 `bson:"focus,omitempty" json:"focus,omitempty"`
	Encounter         *Reference                  `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Issued            *string                     `bson:"issued,omitempty" json:"issued,omitempty"`
	Performer         []Reference                 `bson:"performer,omitempty" json:"performer,omitempty"`
	DataAbsentReason  *CodeableConcept            `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Interpretation    []CodeableConcept           `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	Note              []Annotation                `bson:"note,omitempty" json:"note,omitempty"`
	BodySite          *CodeableConcept            `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Method            *CodeableConcept            `bson:"method,omitempty" json:"method,omitempty"`
	Specimen          *Reference                  `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Device            *Reference                  `bson:"device,omitempty" json:"device,omitempty"`
	ReferenceRange    []ObservationReferenceRange `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
	HasMember         []Reference                 `bson:"hasMember,omitempty" json:"hasMember,omitempty"`
	DerivedFrom       []Reference                 `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	Component         []ObservationComponent      `bson:"component,omitempty" json:"component,omitempty"`
}

// Guidance on how to interpret the value by comparison to a normal or recommended range.  Multiple reference ranges are interpreted as an "OR".   In other words, to represent two distinct target populations, two `referenceRange` elements would be used.
// Most observations only have one generic reference range. Systems MAY choose to restrict to only supplying the relevant reference range based on knowledge about the patient (e.g., specific to the patient's age, gender, weight and other factors), but this might not be possible or appropriate. Whenever more than one reference range is supplied, the differences between them SHOULD be provided in the reference range and/or age properties.
type ObservationReferenceRange struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Low               *Quantity         `bson:"low,omitempty" json:"low,omitempty"`
	High              *Quantity         `bson:"high,omitempty" json:"high,omitempty"`
	Type              *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	AppliesTo         []CodeableConcept `bson:"appliesTo,omitempty" json:"appliesTo,omitempty"`
	Age               *Range            `bson:"age,omitempty" json:"age,omitempty"`
	Text              *string           `bson:"text,omitempty" json:"text,omitempty"`
}

// Some observations have multiple component observations.  These component observations are expressed as separate code value pairs that share the same attributes.  Examples include systolic and diastolic component observations for blood pressure measurement and multiple component observations for genetics observations.
// For a discussion on the ways Observations can be assembled in groups together see [Notes](observation.html#notes) below.
type ObservationComponent struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept             `bson:"code" json:"code"`
	DataAbsentReason  *CodeableConcept            `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Interpretation    []CodeableConcept           `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	ReferenceRange    []ObservationReferenceRange `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
}

// This function returns resource reference information
func (r Observation) ResourceRef() (string, *string) {
	return "Observation", r.Id
}

type OtherObservation Observation

// MarshalJSON marshals the given Observation as JSON into a byte slice
func (r Observation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherObservation
		ResourceType string `json:"resourceType"`
	}{
		OtherObservation: OtherObservation(r),
		ResourceType:     "Observation",
	})
}

// UnmarshalObservation unmarshals a Observation.
func UnmarshalObservation(b []byte) (Observation, error) {
	var observation Observation
	if err := json.Unmarshal(b, &observation); err != nil {
		return observation, err
	}
	return observation, nil
}
