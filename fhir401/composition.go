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

// Composition is documented here http://hl7.org/fhir/StructureDefinition/Composition
// A set of healthcare-related information that is assembled together into a single logical package that provides a single coherent statement of meaning, establishes its own context and that has clinical attestation with regard to who is making the statement. A Composition defines the structure and narrative content necessary for a document. However, a Composition alone does not constitute a document. Rather, the Composition must be the first entry in a Bundle where Bundle.type=document, and any other resources referenced from Composition must be included as subsequent entries in the Bundle (for example Patient, Practitioner, Encounter, etc.).
type Composition struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Contained         []json.RawMessage      `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            CompositionStatus      `bson:"status" json:"status"`
	Type              CodeableConcept        `bson:"type" json:"type"`
	Category          []CodeableConcept      `bson:"category,omitempty" json:"category,omitempty"`
	Subject           *Reference             `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter         *Reference             `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Date              string                 `bson:"date" json:"date"`
	Author            []Reference            `bson:"author" json:"author"`
	Title             string                 `bson:"title" json:"title"`
	Confidentiality   *string                `bson:"confidentiality,omitempty" json:"confidentiality,omitempty"`
	Attester          []CompositionAttester  `bson:"attester,omitempty" json:"attester,omitempty"`
	Custodian         *Reference             `bson:"custodian,omitempty" json:"custodian,omitempty"`
	RelatesTo         []CompositionRelatesTo `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	Event             []CompositionEvent     `bson:"event,omitempty" json:"event,omitempty"`
	Section           []CompositionSection   `bson:"section,omitempty" json:"section,omitempty"`
}

// A participant who has attested to the accuracy of the composition/document.
// Only list each attester once.
type CompositionAttester struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              CompositionAttestationMode `bson:"mode" json:"mode"`
	Time              *string                    `bson:"time,omitempty" json:"time,omitempty"`
	Party             *Reference                 `bson:"party,omitempty" json:"party,omitempty"`
}

// Relationships that this composition has with other compositions or documents that already exist.
// A document is a version specific composition.
type CompositionRelatesTo struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              DocumentRelationshipType `bson:"code" json:"code"`
	TargetIdentifier  Identifier               `bson:"targetIdentifier" json:"targetIdentifier"`
	TargetReference   Reference                `bson:"targetReference" json:"targetReference"`
}

// The clinical service, such as a colonoscopy or an appendectomy, being documented.
// The event needs to be consistent with the type element, though can provide further information if desired.
type CompositionEvent struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Period            *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Detail            []Reference       `bson:"detail,omitempty" json:"detail,omitempty"`
}

// The root of the sections that make up the composition.
type CompositionSection struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Title             *string              `bson:"title,omitempty" json:"title,omitempty"`
	Code              *CodeableConcept     `bson:"code,omitempty" json:"code,omitempty"`
	Author            []Reference          `bson:"author,omitempty" json:"author,omitempty"`
	Focus             *Reference           `bson:"focus,omitempty" json:"focus,omitempty"`
	Text              *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	Mode              *ListMode            `bson:"mode,omitempty" json:"mode,omitempty"`
	OrderedBy         *CodeableConcept     `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	Entry             []Reference          `bson:"entry,omitempty" json:"entry,omitempty"`
	EmptyReason       *CodeableConcept     `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
	Section           []CompositionSection `bson:"section,omitempty" json:"section,omitempty"`
}

// This function returns resource reference information
func (r Composition) ResourceRef() (string, *string) {
	return "Composition", r.Id
}

// This function returns resource reference information
func (r Composition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherComposition Composition

// MarshalJSON marshals the given Composition as JSON into a byte slice
func (r Composition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherComposition
		ResourceType string `json:"resourceType"`
	}{
		OtherComposition: OtherComposition(r),
		ResourceType:     "Composition",
	})
}

// UnmarshalComposition unmarshals a Composition.
func UnmarshalComposition(b []byte) (Composition, error) {
	var composition Composition
	if err := json.Unmarshal(b, &composition); err != nil {
		return composition, err
	}
	return composition, nil
}
