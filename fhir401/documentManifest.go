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

// DocumentManifest is documented here http://hl7.org/fhir/StructureDefinition/DocumentManifest
// A collection of documents compiled for a purpose together with metadata that applies to the collection.
type DocumentManifest struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	Contained         []json.RawMessage         `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	MasterIdentifier  *Identifier               `bson:"masterIdentifier,omitempty" json:"masterIdentifier,omitempty"`
	Identifier        []Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            DocumentReferenceStatus   `bson:"status" json:"status"`
	Type              *CodeableConcept          `bson:"type,omitempty" json:"type,omitempty"`
	Subject           *Reference                `bson:"subject,omitempty" json:"subject,omitempty"`
	Created           *string                   `bson:"created,omitempty" json:"created,omitempty"`
	Author            []Reference               `bson:"author,omitempty" json:"author,omitempty"`
	Recipient         []Reference               `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Source            *string                   `bson:"source,omitempty" json:"source,omitempty"`
	Description       *string                   `bson:"description,omitempty" json:"description,omitempty"`
	Content           []Reference               `bson:"content" json:"content"`
	Related           []DocumentManifestRelated `bson:"related,omitempty" json:"related,omitempty"`
}

// Related identifiers or resources associated with the DocumentManifest.
// May be identifiers or resources that caused the DocumentManifest to be created.
type DocumentManifestRelated struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ref               *Reference  `bson:"ref,omitempty" json:"ref,omitempty"`
}

// This function returns resource reference information
func (r DocumentManifest) ResourceRef() (string, *string) {
	return "DocumentManifest", r.Id
}

// This function returns resource reference information
func (r DocumentManifest) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherDocumentManifest DocumentManifest

// MarshalJSON marshals the given DocumentManifest as JSON into a byte slice
func (r DocumentManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDocumentManifest
		ResourceType string `json:"resourceType"`
	}{
		OtherDocumentManifest: OtherDocumentManifest(r),
		ResourceType:          "DocumentManifest",
	})
}

// UnmarshalDocumentManifest unmarshals a DocumentManifest.
func UnmarshalDocumentManifest(b []byte) (DocumentManifest, error) {
	var documentManifest DocumentManifest
	if err := json.Unmarshal(b, &documentManifest); err != nil {
		return documentManifest, err
	}
	return documentManifest, nil
}
