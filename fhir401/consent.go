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

// Consent is documented here http://hl7.org/fhir/StructureDefinition/Consent
// A record of a healthcare consumer’s  choices, which permits or denies identified recipient(s) or recipient role(s) to perform one or more actions within a given policy context, for specific purposes and periods of time.
type Consent struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Contained         []json.RawMessage     `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            ConsentState          `bson:"status" json:"status"`
	Scope             CodeableConcept       `bson:"scope" json:"scope"`
	Category          []CodeableConcept     `bson:"category" json:"category"`
	Patient           *Reference            `bson:"patient,omitempty" json:"patient,omitempty"`
	DateTime          *string               `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	Performer         []Reference           `bson:"performer,omitempty" json:"performer,omitempty"`
	Organization      []Reference           `bson:"organization,omitempty" json:"organization,omitempty"`
	SourceAttachment  *Attachment           `bson:"sourceAttachment,omitempty" json:"sourceAttachment,omitempty"`
	SourceReference   *Reference            `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
	Policy            []ConsentPolicy       `bson:"policy,omitempty" json:"policy,omitempty"`
	PolicyRule        *CodeableConcept      `bson:"policyRule,omitempty" json:"policyRule,omitempty"`
	Verification      []ConsentVerification `bson:"verification,omitempty" json:"verification,omitempty"`
	Provision         *ConsentProvision     `bson:"provision,omitempty" json:"provision,omitempty"`
}

// The references to the policies that are included in this consent scope. Policies may be organizational, but are often defined jurisdictionally, or in law.
type ConsentPolicy struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Authority         *string     `bson:"authority,omitempty" json:"authority,omitempty"`
	Uri               *string     `bson:"uri,omitempty" json:"uri,omitempty"`
}

// Whether a treatment instruction (e.g. artificial respiration yes or no) was verified with the patient, his/her family or another authorized person.
type ConsentVerification struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Verified          bool        `bson:"verified" json:"verified"`
	VerifiedWith      *Reference  `bson:"verifiedWith,omitempty" json:"verifiedWith,omitempty"`
	VerificationDate  *string     `bson:"verificationDate,omitempty" json:"verificationDate,omitempty"`
}

// An exception to the base policy of this consent. An exception can be an addition or removal of access permissions.
type ConsentProvision struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *ConsentProvisionType   `bson:"type,omitempty" json:"type,omitempty"`
	Period            *Period                 `bson:"period,omitempty" json:"period,omitempty"`
	Actor             []ConsentProvisionActor `bson:"actor,omitempty" json:"actor,omitempty"`
	Action            []CodeableConcept       `bson:"action,omitempty" json:"action,omitempty"`
	SecurityLabel     []Coding                `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Purpose           []Coding                `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Class             []Coding                `bson:"class,omitempty" json:"class,omitempty"`
	Code              []CodeableConcept       `bson:"code,omitempty" json:"code,omitempty"`
	DataPeriod        *Period                 `bson:"dataPeriod,omitempty" json:"dataPeriod,omitempty"`
	Data              []ConsentProvisionData  `bson:"data,omitempty" json:"data,omitempty"`
	Provision         []ConsentProvision      `bson:"provision,omitempty" json:"provision,omitempty"`
}

// Who or what is controlled by this rule. Use group to identify a set of actors by some property they share (e.g. 'admitting officers').
// There is no specific actor associated with the exception
type ConsentProvisionActor struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              CodeableConcept `bson:"role" json:"role"`
	Reference         Reference       `bson:"reference" json:"reference"`
}

// The resources controlled by this rule if specific resources are referenced.
// all data
type ConsentProvisionData struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Meaning           ConsentDataMeaning `bson:"meaning" json:"meaning"`
	Reference         Reference          `bson:"reference" json:"reference"`
}

// This function returns resource reference information
func (r Consent) ResourceRef() (string, *string) {
	return "Consent", r.Id
}

// This function returns resource reference information
func (r Consent) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherConsent Consent

// MarshalJSON marshals the given Consent as JSON into a byte slice
func (r Consent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherConsent
		ResourceType string `json:"resourceType"`
	}{
		OtherConsent: OtherConsent(r),
		ResourceType: "Consent",
	})
}

// UnmarshalConsent unmarshals a Consent.
func UnmarshalConsent(b []byte) (Consent, error) {
	var consent Consent
	if err := json.Unmarshal(b, &consent); err != nil {
		return consent, err
	}
	return consent, nil
}
