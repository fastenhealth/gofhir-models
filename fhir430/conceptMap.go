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

// ConceptMap is documented here http://hl7.org/fhir/StructureDefinition/ConceptMap
// A statement of relationships from one set of concepts to one or more other concepts - either concepts in code systems, or data element/data element concepts, or classes in class models.
type ConceptMap struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string           `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        *Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string           `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string           `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string           `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus `bson:"status" json:"status"`
	Experimental      *bool             `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string           `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string           `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail   `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string           `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext    `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string           `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string           `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Group             []ConceptMapGroup `bson:"group,omitempty" json:"group,omitempty"`
}

// A group of mappings that all have the same source and target system.
type ConceptMapGroup struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Source            *string                  `bson:"source,omitempty" json:"source,omitempty"`
	SourceVersion     *string                  `bson:"sourceVersion,omitempty" json:"sourceVersion,omitempty"`
	Target            *string                  `bson:"target,omitempty" json:"target,omitempty"`
	TargetVersion     *string                  `bson:"targetVersion,omitempty" json:"targetVersion,omitempty"`
	Element           []ConceptMapGroupElement `bson:"element" json:"element"`
	Unmapped          *ConceptMapGroupUnmapped `bson:"unmapped,omitempty" json:"unmapped,omitempty"`
}

// Mappings for an individual concept in the source to one or more concepts in the target.
// Generally, the ideal is that there would only be one mapping for each concept in the source value set, but a given concept may be mapped multiple times with different comments or dependencies.
type ConceptMapGroupElement struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              *string                        `bson:"code,omitempty" json:"code,omitempty"`
	Display           *string                        `bson:"display,omitempty" json:"display,omitempty"`
	Target            []ConceptMapGroupElementTarget `bson:"target,omitempty" json:"target,omitempty"`
}

// A concept from the target value set that this concept maps to.
// Ideally there would only be one map, with equal or equivalent mapping. But multiple maps are allowed for several narrower options, or to assert that other concepts are unmatched.
type ConceptMapGroupElementTarget struct {
	Id                *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              *string                                 `bson:"code,omitempty" json:"code,omitempty"`
	Display           *string                                 `bson:"display,omitempty" json:"display,omitempty"`
	Equivalence       ConceptMapEquivalence                   `bson:"equivalence" json:"equivalence"`
	Comment           *string                                 `bson:"comment,omitempty" json:"comment,omitempty"`
	DependsOn         []ConceptMapGroupElementTargetDependsOn `bson:"dependsOn,omitempty" json:"dependsOn,omitempty"`
	Product           []ConceptMapGroupElementTargetDependsOn `bson:"product,omitempty" json:"product,omitempty"`
}

// A set of additional dependencies for this mapping to hold. This mapping is only applicable if the specified element can be resolved, and it has the specified value.
type ConceptMapGroupElementTargetDependsOn struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Property          string      `bson:"property" json:"property"`
	System            *string     `bson:"system,omitempty" json:"system,omitempty"`
	Value             string      `bson:"value" json:"value"`
	Display           *string     `bson:"display,omitempty" json:"display,omitempty"`
}

// What to do when there is no mapping for the source concept. "Unmapped" does not include codes that are unmatched, and the unmapped element is ignored in a code is specified to have equivalence = unmatched.
// This only applies if the source code has a system value that matches the system defined for the group.
type ConceptMapGroupUnmapped struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              ConceptMapGroupUnmappedMode `bson:"mode" json:"mode"`
	Code              *string                     `bson:"code,omitempty" json:"code,omitempty"`
	Display           *string                     `bson:"display,omitempty" json:"display,omitempty"`
	Url               *string                     `bson:"url,omitempty" json:"url,omitempty"`
}
type OtherConceptMap ConceptMap

// MarshalJSON marshals the given ConceptMap as JSON into a byte slice
func (r ConceptMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherConceptMap
		ResourceType string `json:"resourceType"`
	}{
		OtherConceptMap: OtherConceptMap(r),
		ResourceType:    "ConceptMap",
	})
}

// UnmarshalConceptMap unmarshals a ConceptMap.
func UnmarshalConceptMap(b []byte) (ConceptMap, error) {
	var conceptMap ConceptMap
	if err := json.Unmarshal(b, &conceptMap); err != nil {
		return conceptMap, err
	}
	return conceptMap, nil
}
