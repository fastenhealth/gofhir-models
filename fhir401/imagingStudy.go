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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// ImagingStudy is documented here http://hl7.org/fhir/StructureDefinition/ImagingStudy
// Representation of the content produced in a DICOM imaging study. A study comprises a set of series, each of which includes a set of Service-Object Pair Instances (SOP Instances - images or other data) acquired or produced in a common context.  A series is of only one modality (e.g. X-ray, CT, MR, ultrasound), but a study may have multiple series of different modalities.
type ImagingStudy struct {
	Id                 *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	Extension          []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status             ImagingStudyStatus   `bson:"status" json:"status"`
	Modality           []Coding             `bson:"modality,omitempty" json:"modality,omitempty"`
	Subject            Reference            `bson:"subject" json:"subject"`
	Encounter          *Reference           `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Started            *string              `bson:"started,omitempty" json:"started,omitempty"`
	BasedOn            []Reference          `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Referrer           *Reference           `bson:"referrer,omitempty" json:"referrer,omitempty"`
	Interpreter        []Reference          `bson:"interpreter,omitempty" json:"interpreter,omitempty"`
	Endpoint           []Reference          `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	NumberOfSeries     *int                 `bson:"numberOfSeries,omitempty" json:"numberOfSeries,omitempty"`
	NumberOfInstances  *int                 `bson:"numberOfInstances,omitempty" json:"numberOfInstances,omitempty"`
	ProcedureReference *Reference           `bson:"procedureReference,omitempty" json:"procedureReference,omitempty"`
	ProcedureCode      []CodeableConcept    `bson:"procedureCode,omitempty" json:"procedureCode,omitempty"`
	Location           *Reference           `bson:"location,omitempty" json:"location,omitempty"`
	ReasonCode         []CodeableConcept    `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference    []Reference          `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note               []Annotation         `bson:"note,omitempty" json:"note,omitempty"`
	Description        *string              `bson:"description,omitempty" json:"description,omitempty"`
	Series             []ImagingStudySeries `bson:"series,omitempty" json:"series,omitempty"`
}

// Each study has one or more series of images or other content.
type ImagingStudySeries struct {
	Id                *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Uid               string                        `bson:"uid" json:"uid"`
	Number            *int                          `bson:"number,omitempty" json:"number,omitempty"`
	Modality          Coding                        `bson:"modality" json:"modality"`
	Description       *string                       `bson:"description,omitempty" json:"description,omitempty"`
	NumberOfInstances *int                          `bson:"numberOfInstances,omitempty" json:"numberOfInstances,omitempty"`
	Endpoint          []Reference                   `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	BodySite          *Coding                       `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Laterality        *Coding                       `bson:"laterality,omitempty" json:"laterality,omitempty"`
	Specimen          []Reference                   `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Started           *string                       `bson:"started,omitempty" json:"started,omitempty"`
	Performer         []ImagingStudySeriesPerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	Instance          []ImagingStudySeriesInstance  `bson:"instance,omitempty" json:"instance,omitempty"`
}

// Indicates who or what performed the series and how they were involved.
// If the person who performed the series is not known, their Organization may be recorded. A patient, or related person, may be the performer, e.g. for patient-captured images.
type ImagingStudySeriesPerformer struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	Actor             Reference        `bson:"actor" json:"actor"`
}

// A single SOP instance within the series, e.g. an image, or presentation state.
type ImagingStudySeriesInstance struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Uid               string      `bson:"uid" json:"uid"`
	SopClass          Coding      `bson:"sopClass" json:"sopClass"`
	Number            *int        `bson:"number,omitempty" json:"number,omitempty"`
	Title             *string     `bson:"title,omitempty" json:"title,omitempty"`
}
type OtherImagingStudy ImagingStudy

// MarshalJSON marshals the given ImagingStudy as JSON into a byte slice
func (r ImagingStudy) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImagingStudy
		ResourceType string `json:"resourceType"`
	}{
		OtherImagingStudy: OtherImagingStudy(r),
		ResourceType:      "ImagingStudy",
	})
}

// UnmarshalImagingStudy unmarshals a ImagingStudy.
func UnmarshalImagingStudy(b []byte) (ImagingStudy, error) {
	var imagingStudy ImagingStudy
	if err := json.Unmarshal(b, &imagingStudy); err != nil {
		return imagingStudy, err
	}
	return imagingStudy, nil
}
