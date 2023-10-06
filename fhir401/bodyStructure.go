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

// BodyStructure is documented here http://hl7.org/fhir/StructureDefinition/BodyStructure
// Record details about an anatomical structure.  This resource may be used when a coded concept does not provide the necessary detail needed for the use case.
type BodyStructure struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Contained         []json.RawMessage `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool             `bson:"active,omitempty" json:"active,omitempty"`
	Morphology        *CodeableConcept  `bson:"morphology,omitempty" json:"morphology,omitempty"`
	Location          *CodeableConcept  `bson:"location,omitempty" json:"location,omitempty"`
	LocationQualifier []CodeableConcept `bson:"locationQualifier,omitempty" json:"locationQualifier,omitempty"`
	Description       *string           `bson:"description,omitempty" json:"description,omitempty"`
	Image             []Attachment      `bson:"image,omitempty" json:"image,omitempty"`
	Patient           Reference         `bson:"patient" json:"patient"`
}

// This function returns resource reference information
func (r BodyStructure) ResourceRef() (string, *string) {
	return "BodyStructure", r.Id
}

// This function returns resource reference information
func (r BodyStructure) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherBodyStructure BodyStructure

// MarshalJSON marshals the given BodyStructure as JSON into a byte slice
func (r BodyStructure) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBodyStructure
		ResourceType string `json:"resourceType"`
	}{
		OtherBodyStructure: OtherBodyStructure(r),
		ResourceType:       "BodyStructure",
	})
}

// UnmarshalBodyStructure unmarshals a BodyStructure.
func UnmarshalBodyStructure(b []byte) (BodyStructure, error) {
	var bodyStructure BodyStructure
	if err := json.Unmarshal(b, &bodyStructure); err != nil {
		return bodyStructure, err
	}
	return bodyStructure, nil
}
