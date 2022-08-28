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

// TerminologyCapabilities is documented here http://hl7.org/fhir/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilities struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                             `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                                `bson:"url,omitempty" json:"url,omitempty"`
	Version           *string                                `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string                                `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                                `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus                      `bson:"status" json:"status"`
	Experimental      *bool                                  `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              string                                 `bson:"date" json:"date"`
	Publisher         *string                                `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail                        `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                                `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext                         `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                      `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                                `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string                                `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Kind              CapabilityStatementKind                `bson:"kind" json:"kind"`
	Software          *TerminologyCapabilitiesSoftware       `bson:"software,omitempty" json:"software,omitempty"`
	Implementation    *TerminologyCapabilitiesImplementation `bson:"implementation,omitempty" json:"implementation,omitempty"`
	LockedDate        *bool                                  `bson:"lockedDate,omitempty" json:"lockedDate,omitempty"`
	CodeSystem        []TerminologyCapabilitiesCodeSystem    `bson:"codeSystem,omitempty" json:"codeSystem,omitempty"`
	Expansion         *TerminologyCapabilitiesExpansion      `bson:"expansion,omitempty" json:"expansion,omitempty"`
	CodeSearch        *CodeSearchSupport                     `bson:"codeSearch,omitempty" json:"codeSearch,omitempty"`
	ValidateCode      *TerminologyCapabilitiesValidateCode   `bson:"validateCode,omitempty" json:"validateCode,omitempty"`
	Translation       *TerminologyCapabilitiesTranslation    `bson:"translation,omitempty" json:"translation,omitempty"`
	Closure           *TerminologyCapabilitiesClosure        `bson:"closure,omitempty" json:"closure,omitempty"`
}
type TerminologyCapabilitiesSoftware struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name" json:"name"`
	Version           *string     `bson:"version,omitempty" json:"version,omitempty"`
}
type TerminologyCapabilitiesImplementation struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       string      `bson:"description" json:"description"`
	Url               *string     `bson:"url,omitempty" json:"url,omitempty"`
}
type TerminologyCapabilitiesCodeSystem struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Uri               *string                                    `bson:"uri,omitempty" json:"uri,omitempty"`
	Version           []TerminologyCapabilitiesCodeSystemVersion `bson:"version,omitempty" json:"version,omitempty"`
	Subsumption       *bool                                      `bson:"subsumption,omitempty" json:"subsumption,omitempty"`
}
type TerminologyCapabilitiesCodeSystemVersion struct {
	Id                *string                                          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              *string                                          `bson:"code,omitempty" json:"code,omitempty"`
	IsDefault         *bool                                            `bson:"isDefault,omitempty" json:"isDefault,omitempty"`
	Compositional     *bool                                            `bson:"compositional,omitempty" json:"compositional,omitempty"`
	Language          []string                                         `bson:"language,omitempty" json:"language,omitempty"`
	Filter            []TerminologyCapabilitiesCodeSystemVersionFilter `bson:"filter,omitempty" json:"filter,omitempty"`
	Property          []string                                         `bson:"property,omitempty" json:"property,omitempty"`
}
type TerminologyCapabilitiesCodeSystemVersionFilter struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string      `bson:"code" json:"code"`
	Op                []string    `bson:"op" json:"op"`
}
type TerminologyCapabilitiesExpansion struct {
	Id                *string                                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Hierarchical      *bool                                       `bson:"hierarchical,omitempty" json:"hierarchical,omitempty"`
	Paging            *bool                                       `bson:"paging,omitempty" json:"paging,omitempty"`
	Incomplete        *bool                                       `bson:"incomplete,omitempty" json:"incomplete,omitempty"`
	Parameter         []TerminologyCapabilitiesExpansionParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	TextFilter        *string                                     `bson:"textFilter,omitempty" json:"textFilter,omitempty"`
}
type TerminologyCapabilitiesExpansionParameter struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name" json:"name"`
	Documentation     *string     `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type TerminologyCapabilitiesValidateCode struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Translations      bool        `bson:"translations" json:"translations"`
}
type TerminologyCapabilitiesTranslation struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	NeedsMap          bool        `bson:"needsMap" json:"needsMap"`
}
type TerminologyCapabilitiesClosure struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Translation       *bool       `bson:"translation,omitempty" json:"translation,omitempty"`
}
type OtherTerminologyCapabilities TerminologyCapabilities

// MarshalJSON marshals the given TerminologyCapabilities as JSON into a byte slice
func (r TerminologyCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTerminologyCapabilities
		ResourceType string `json:"resourceType"`
	}{
		OtherTerminologyCapabilities: OtherTerminologyCapabilities(r),
		ResourceType:                 "TerminologyCapabilities",
	})
}

// UnmarshalTerminologyCapabilities unmarshals a TerminologyCapabilities.
func UnmarshalTerminologyCapabilities(b []byte) (TerminologyCapabilities, error) {
	var terminologyCapabilities TerminologyCapabilities
	if err := json.Unmarshal(b, &terminologyCapabilities); err != nil {
		return terminologyCapabilities, err
	}
	return terminologyCapabilities, nil
}