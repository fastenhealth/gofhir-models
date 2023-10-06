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

// CommunicationRequest is documented here http://hl7.org/fhir/StructureDefinition/CommunicationRequest
// A request to convey information; e.g. the CDS system proposes that an alert be sent to a responsible provider, the CDS system proposes that the public health agency be notified about a reportable condition.
type CommunicationRequest struct {
	Id                 *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                         `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                       `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                       `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative                    `bson:"text,omitempty" json:"text,omitempty"`
	Contained          []json.RawMessage             `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension          []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn            []Reference                   `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces           []Reference                   `bson:"replaces,omitempty" json:"replaces,omitempty"`
	GroupIdentifier    *Identifier                   `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status             RequestStatus                 `bson:"status" json:"status"`
	StatusReason       *CodeableConcept              `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Category           []CodeableConcept             `bson:"category,omitempty" json:"category,omitempty"`
	Priority           *RequestPriority              `bson:"priority,omitempty" json:"priority,omitempty"`
	DoNotPerform       *bool                         `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	Medium             []CodeableConcept             `bson:"medium,omitempty" json:"medium,omitempty"`
	Subject            *Reference                    `bson:"subject,omitempty" json:"subject,omitempty"`
	About              []Reference                   `bson:"about,omitempty" json:"about,omitempty"`
	Encounter          *Reference                    `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Payload            []CommunicationRequestPayload `bson:"payload,omitempty" json:"payload,omitempty"`
	OccurrenceDateTime *string                       `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                       `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	AuthoredOn         *string                       `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester          *Reference                    `bson:"requester,omitempty" json:"requester,omitempty"`
	Recipient          []Reference                   `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Sender             *Reference                    `bson:"sender,omitempty" json:"sender,omitempty"`
	ReasonCode         []CodeableConcept             `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference    []Reference                   `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note               []Annotation                  `bson:"note,omitempty" json:"note,omitempty"`
}

// Text, attachment(s), or resource(s) to be communicated to the recipient.
type CommunicationRequestPayload struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ContentString     string      `bson:"contentString" json:"contentString"`
	ContentAttachment Attachment  `bson:"contentAttachment" json:"contentAttachment"`
	ContentReference  Reference   `bson:"contentReference" json:"contentReference"`
}

// This function returns resource reference information
func (r CommunicationRequest) ResourceRef() (string, *string) {
	return "CommunicationRequest", r.Id
}

// This function returns resource reference information
func (r CommunicationRequest) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCommunicationRequest CommunicationRequest

// MarshalJSON marshals the given CommunicationRequest as JSON into a byte slice
func (r CommunicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCommunicationRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherCommunicationRequest: OtherCommunicationRequest(r),
		ResourceType:              "CommunicationRequest",
	})
}

// UnmarshalCommunicationRequest unmarshals a CommunicationRequest.
func UnmarshalCommunicationRequest(b []byte) (CommunicationRequest, error) {
	var communicationRequest CommunicationRequest
	if err := json.Unmarshal(b, &communicationRequest); err != nil {
		return communicationRequest, err
	}
	return communicationRequest, nil
}
