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

// CarePlan is documented here http://hl7.org/fhir/StructureDefinition/CarePlan
// Describes the intention of how one or more practitioners intend to deliver care for a particular patient, group or community for a period of time, possibly limited to care for a specific condition or set of conditions.
type CarePlan struct {
	Id                    *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	Contained             []json.RawMessage  `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension             []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical []string           `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string           `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	BasedOn               []Reference        `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces              []Reference        `bson:"replaces,omitempty" json:"replaces,omitempty"`
	PartOf                []Reference        `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                RequestStatus      `bson:"status" json:"status"`
	Intent                CarePlanIntent     `bson:"intent" json:"intent"`
	Category              []CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Title                 *string            `bson:"title,omitempty" json:"title,omitempty"`
	Description           *string            `bson:"description,omitempty" json:"description,omitempty"`
	Subject               Reference          `bson:"subject" json:"subject"`
	Encounter             *Reference         `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Period                *Period            `bson:"period,omitempty" json:"period,omitempty"`
	Created               *string            `bson:"created,omitempty" json:"created,omitempty"`
	Author                *Reference         `bson:"author,omitempty" json:"author,omitempty"`
	Contributor           []Reference        `bson:"contributor,omitempty" json:"contributor,omitempty"`
	CareTeam              []Reference        `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	Addresses             []Reference        `bson:"addresses,omitempty" json:"addresses,omitempty"`
	SupportingInfo        []Reference        `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Goal                  []Reference        `bson:"goal,omitempty" json:"goal,omitempty"`
	Activity              []CarePlanActivity `bson:"activity,omitempty" json:"activity,omitempty"`
	Note                  []Annotation       `bson:"note,omitempty" json:"note,omitempty"`
}

// Identifies a planned action to occur as part of the plan.  For example, a medication to be used, lab tests to perform, self-monitoring, education, etc.
type CarePlanActivity struct {
	Id                     *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	OutcomeCodeableConcept []CodeableConcept       `bson:"outcomeCodeableConcept,omitempty" json:"outcomeCodeableConcept,omitempty"`
	OutcomeReference       []Reference             `bson:"outcomeReference,omitempty" json:"outcomeReference,omitempty"`
	Progress               []Annotation            `bson:"progress,omitempty" json:"progress,omitempty"`
	Reference              *Reference              `bson:"reference,omitempty" json:"reference,omitempty"`
	Detail                 *CarePlanActivityDetail `bson:"detail,omitempty" json:"detail,omitempty"`
}

// A simple summary of a planned activity suitable for a general care plan system (e.g. form driven) that doesn't know about specific resources such as procedure etc.
type CarePlanActivityDetail struct {
	Id                     *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Kind                   *CarePlanActivityKind  `bson:"kind,omitempty" json:"kind,omitempty"`
	InstantiatesCanonical  []string               `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri        []string               `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	Code                   *CodeableConcept       `bson:"code,omitempty" json:"code,omitempty"`
	ReasonCode             []CodeableConcept      `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference        []Reference            `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Goal                   []Reference            `bson:"goal,omitempty" json:"goal,omitempty"`
	Status                 CarePlanActivityStatus `bson:"status" json:"status"`
	StatusReason           *CodeableConcept       `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	DoNotPerform           *bool                  `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	ScheduledTiming        *Timing                `bson:"scheduledTiming,omitempty" json:"scheduledTiming,omitempty"`
	ScheduledPeriod        *Period                `bson:"scheduledPeriod,omitempty" json:"scheduledPeriod,omitempty"`
	ScheduledString        *string                `bson:"scheduledString,omitempty" json:"scheduledString,omitempty"`
	Location               *Reference             `bson:"location,omitempty" json:"location,omitempty"`
	Performer              []Reference            `bson:"performer,omitempty" json:"performer,omitempty"`
	ProductCodeableConcept *CodeableConcept       `bson:"productCodeableConcept,omitempty" json:"productCodeableConcept,omitempty"`
	ProductReference       *Reference             `bson:"productReference,omitempty" json:"productReference,omitempty"`
	DailyAmount            *Quantity              `bson:"dailyAmount,omitempty" json:"dailyAmount,omitempty"`
	Quantity               *Quantity              `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Description            *string                `bson:"description,omitempty" json:"description,omitempty"`
}

// This function returns resource reference information
func (r CarePlan) ResourceRef() (string, *string) {
	return "CarePlan", r.Id
}

// This function returns resource reference information
func (r CarePlan) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCarePlan CarePlan

// MarshalJSON marshals the given CarePlan as JSON into a byte slice
func (r CarePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCarePlan
		ResourceType string `json:"resourceType"`
	}{
		OtherCarePlan: OtherCarePlan(r),
		ResourceType:  "CarePlan",
	})
}

// UnmarshalCarePlan unmarshals a CarePlan.
func UnmarshalCarePlan(b []byte) (CarePlan, error) {
	var carePlan CarePlan
	if err := json.Unmarshal(b, &carePlan); err != nil {
		return carePlan, err
	}
	return carePlan, nil
}
