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

// Media is documented here http://hl7.org/fhir/StructureDefinition/Media
// A photo, video, or audio recording acquired or used in healthcare. The actual content may be inline or provided by direct reference.
type Media struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Contained         []json.RawMessage `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn           []Reference       `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf            []Reference       `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status            EventStatus       `bson:"status" json:"status"`
	Type              *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Modality          *CodeableConcept  `bson:"modality,omitempty" json:"modality,omitempty"`
	View              *CodeableConcept  `bson:"view,omitempty" json:"view,omitempty"`
	Subject           *Reference        `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter         *Reference        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	CreatedDateTime   *string           `bson:"createdDateTime,omitempty" json:"createdDateTime,omitempty"`
	CreatedPeriod     *Period           `bson:"createdPeriod,omitempty" json:"createdPeriod,omitempty"`
	Issued            *string           `bson:"issued,omitempty" json:"issued,omitempty"`
	Operator          *Reference        `bson:"operator,omitempty" json:"operator,omitempty"`
	ReasonCode        []CodeableConcept `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	BodySite          *CodeableConcept  `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	DeviceName        *string           `bson:"deviceName,omitempty" json:"deviceName,omitempty"`
	Device            *Reference        `bson:"device,omitempty" json:"device,omitempty"`
	Height            *int              `bson:"height,omitempty" json:"height,omitempty"`
	Width             *int              `bson:"width,omitempty" json:"width,omitempty"`
	Frames            *int              `bson:"frames,omitempty" json:"frames,omitempty"`
	Duration          *json.Number      `bson:"duration,omitempty" json:"duration,omitempty"`
	Content           Attachment        `bson:"content" json:"content"`
	Note              []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
}

// This function returns resource reference information
func (r Media) ResourceRef() (string, *string) {
	return "Media", r.Id
}

// This function returns resource reference information
func (r Media) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMedia Media

// MarshalJSON marshals the given Media as JSON into a byte slice
func (r Media) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedia
		ResourceType string `json:"resourceType"`
	}{
		OtherMedia:   OtherMedia(r),
		ResourceType: "Media",
	})
}

// UnmarshalMedia unmarshals a Media.
func UnmarshalMedia(b []byte) (Media, error) {
	var media Media
	if err := json.Unmarshal(b, &media); err != nil {
		return media, err
	}
	return media, nil
}
