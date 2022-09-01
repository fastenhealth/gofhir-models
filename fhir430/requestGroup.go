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

// RequestGroup is documented here http://hl7.org/fhir/StructureDefinition/RequestGroup
// A group of related requests that can be used to capture intended activities that have inter-dependencies such as "give this medication after that one".
type RequestGroup struct {
	Id                    *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical []string             `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string             `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	BasedOn               []Reference          `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces              []Reference          `bson:"replaces,omitempty" json:"replaces,omitempty"`
	GroupIdentifier       *Identifier          `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status                RequestStatus        `bson:"status" json:"status"`
	Intent                RequestIntent        `bson:"intent" json:"intent"`
	Priority              *RequestPriority     `bson:"priority,omitempty" json:"priority,omitempty"`
	Code                  *CodeableConcept     `bson:"code,omitempty" json:"code,omitempty"`
	Subject               *Reference           `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter             *Reference           `bson:"encounter,omitempty" json:"encounter,omitempty"`
	AuthoredOn            *string              `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Author                *Reference           `bson:"author,omitempty" json:"author,omitempty"`
	ReasonCode            []CodeableConcept    `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference       []Reference          `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note                  []Annotation         `bson:"note,omitempty" json:"note,omitempty"`
	Action                []RequestGroupAction `bson:"action,omitempty" json:"action,omitempty"`
}

// The actions, if any, produced by the evaluation of the artifact.
type RequestGroupAction struct {
	Id                  *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Prefix              *string                           `bson:"prefix,omitempty" json:"prefix,omitempty"`
	Title               *string                           `bson:"title,omitempty" json:"title,omitempty"`
	Description         *string                           `bson:"description,omitempty" json:"description,omitempty"`
	TextEquivalent      *string                           `bson:"textEquivalent,omitempty" json:"textEquivalent,omitempty"`
	Priority            *RequestPriority                  `bson:"priority,omitempty" json:"priority,omitempty"`
	Code                []CodeableConcept                 `bson:"code,omitempty" json:"code,omitempty"`
	Documentation       []RelatedArtifact                 `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Condition           []RequestGroupActionCondition     `bson:"condition,omitempty" json:"condition,omitempty"`
	RelatedAction       []RequestGroupActionRelatedAction `bson:"relatedAction,omitempty" json:"relatedAction,omitempty"`
	Participant         []Reference                       `bson:"participant,omitempty" json:"participant,omitempty"`
	Type                *CodeableConcept                  `bson:"type,omitempty" json:"type,omitempty"`
	GroupingBehavior    *ActionGroupingBehavior           `bson:"groupingBehavior,omitempty" json:"groupingBehavior,omitempty"`
	SelectionBehavior   *ActionSelectionBehavior          `bson:"selectionBehavior,omitempty" json:"selectionBehavior,omitempty"`
	RequiredBehavior    *ActionRequiredBehavior           `bson:"requiredBehavior,omitempty" json:"requiredBehavior,omitempty"`
	PrecheckBehavior    *ActionPrecheckBehavior           `bson:"precheckBehavior,omitempty" json:"precheckBehavior,omitempty"`
	CardinalityBehavior *ActionCardinalityBehavior        `bson:"cardinalityBehavior,omitempty" json:"cardinalityBehavior,omitempty"`
	Resource            *Reference                        `bson:"resource,omitempty" json:"resource,omitempty"`
	Action              []RequestGroupAction              `bson:"action,omitempty" json:"action,omitempty"`
}

// An expression that describes applicability criteria, or start/stop conditions for the action.
// When multiple conditions of the same kind are present, the effects are combined using AND semantics, so the overall condition is true only if all of the conditions are true.
type RequestGroupActionCondition struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Kind              ActionConditionKind `bson:"kind" json:"kind"`
	Expression        *Expression         `bson:"expression,omitempty" json:"expression,omitempty"`
}

// A relationship to another action such as "before" or "30-60 minutes after start of".
type RequestGroupActionRelatedAction struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ActionId          string                 `bson:"actionId" json:"actionId"`
	Relationship      ActionRelationshipType `bson:"relationship" json:"relationship"`
}
type OtherRequestGroup RequestGroup

// MarshalJSON marshals the given RequestGroup as JSON into a byte slice
func (r RequestGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRequestGroup
		ResourceType string `json:"resourceType"`
	}{
		OtherRequestGroup: OtherRequestGroup(r),
		ResourceType:      "RequestGroup",
	})
}

// UnmarshalRequestGroup unmarshals a RequestGroup.
func UnmarshalRequestGroup(b []byte) (RequestGroup, error) {
	var requestGroup RequestGroup
	if err := json.Unmarshal(b, &requestGroup); err != nil {
		return requestGroup, err
	}
	return requestGroup, nil
}
