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

// CompartmentDefinition is documented here http://hl7.org/fhir/StructureDefinition/CompartmentDefinition
// A compartment definition that defines how resources are accessed on a server.
type CompartmentDefinition struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	Contained         []json.RawMessage               `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                          `bson:"url" json:"url"`
	Version           *string                         `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                          `bson:"name" json:"name"`
	Status            PublicationStatus               `bson:"status" json:"status"`
	Experimental      *bool                           `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                         `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                         `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail                 `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                         `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext                  `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Purpose           *string                         `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Code              CompartmentType                 `bson:"code" json:"code"`
	Search            bool                            `bson:"search" json:"search"`
	Resource          []CompartmentDefinitionResource `bson:"resource,omitempty" json:"resource,omitempty"`
}

// Information about how a resource is related to the compartment.
type CompartmentDefinitionResource struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              ResourceType `bson:"code" json:"code"`
	Param             []string     `bson:"param,omitempty" json:"param,omitempty"`
	Documentation     *string      `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

// This function returns resource reference information
func (r CompartmentDefinition) ResourceRef() (string, *string) {
	return "CompartmentDefinition", r.Id
}

// This function returns resource reference information
func (r CompartmentDefinition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCompartmentDefinition CompartmentDefinition

// MarshalJSON marshals the given CompartmentDefinition as JSON into a byte slice
func (r CompartmentDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCompartmentDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherCompartmentDefinition: OtherCompartmentDefinition(r),
		ResourceType:               "CompartmentDefinition",
	})
}

// UnmarshalCompartmentDefinition unmarshals a CompartmentDefinition.
func UnmarshalCompartmentDefinition(b []byte) (CompartmentDefinition, error) {
	var compartmentDefinition CompartmentDefinition
	if err := json.Unmarshal(b, &compartmentDefinition); err != nil {
		return compartmentDefinition, err
	}
	return compartmentDefinition, nil
}
