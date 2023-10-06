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

// MedicationRequest is documented here http://hl7.org/fhir/StructureDefinition/MedicationRequest
// An order or request for both supply of the medication and the instructions for administration of the medication to a patient. The resource is called "MedicationRequest" rather than "MedicationPrescription" or "MedicationOrder" to generalize the use across inpatient and outpatient settings, including care plans, etc., and to harmonize with workflow patterns.
type MedicationRequest struct {
	Id                        *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                      *Meta                             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules             *string                           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                  *string                           `bson:"language,omitempty" json:"language,omitempty"`
	Text                      *Narrative                        `bson:"text,omitempty" json:"text,omitempty"`
	Contained                 []json.RawMessage                 `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension                 []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier                []Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                    string                            `bson:"status" json:"status"`
	StatusReason              *CodeableConcept                  `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Intent                    string                            `bson:"intent" json:"intent"`
	Category                  []CodeableConcept                 `bson:"category,omitempty" json:"category,omitempty"`
	Priority                  *RequestPriority                  `bson:"priority,omitempty" json:"priority,omitempty"`
	DoNotPerform              *bool                             `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	ReportedBoolean           *bool                             `bson:"reportedBoolean,omitempty" json:"reportedBoolean,omitempty"`
	ReportedReference         *Reference                        `bson:"reportedReference,omitempty" json:"reportedReference,omitempty"`
	MedicationCodeableConcept CodeableConcept                   `bson:"medicationCodeableConcept" json:"medicationCodeableConcept"`
	MedicationReference       Reference                         `bson:"medicationReference" json:"medicationReference"`
	Subject                   Reference                         `bson:"subject" json:"subject"`
	Encounter                 *Reference                        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	SupportingInformation     []Reference                       `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	AuthoredOn                *string                           `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester                 *Reference                        `bson:"requester,omitempty" json:"requester,omitempty"`
	Performer                 *Reference                        `bson:"performer,omitempty" json:"performer,omitempty"`
	PerformerType             *CodeableConcept                  `bson:"performerType,omitempty" json:"performerType,omitempty"`
	Recorder                  *Reference                        `bson:"recorder,omitempty" json:"recorder,omitempty"`
	ReasonCode                []CodeableConcept                 `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference           []Reference                       `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	InstantiatesCanonical     []string                          `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri           []string                          `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	BasedOn                   []Reference                       `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	GroupIdentifier           *Identifier                       `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	CourseOfTherapyType       *CodeableConcept                  `bson:"courseOfTherapyType,omitempty" json:"courseOfTherapyType,omitempty"`
	Insurance                 []Reference                       `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Note                      []Annotation                      `bson:"note,omitempty" json:"note,omitempty"`
	DosageInstruction         []Dosage                          `bson:"dosageInstruction,omitempty" json:"dosageInstruction,omitempty"`
	DispenseRequest           *MedicationRequestDispenseRequest `bson:"dispenseRequest,omitempty" json:"dispenseRequest,omitempty"`
	Substitution              *MedicationRequestSubstitution    `bson:"substitution,omitempty" json:"substitution,omitempty"`
	PriorPrescription         *Reference                        `bson:"priorPrescription,omitempty" json:"priorPrescription,omitempty"`
	DetectedIssue             []Reference                       `bson:"detectedIssue,omitempty" json:"detectedIssue,omitempty"`
	EventHistory              []Reference                       `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}

// Indicates the specific details for the dispense or medication supply part of a medication request (also known as a Medication Prescription or Medication Order).  Note that this information is not always sent with the order.  There may be in some settings (e.g. hospitals) institutional or system support for completing the dispense details in the pharmacy department.
type MedicationRequestDispenseRequest struct {
	Id                     *string                                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension                                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	InitialFill            *MedicationRequestDispenseRequestInitialFill `bson:"initialFill,omitempty" json:"initialFill,omitempty"`
	DispenseInterval       *Duration                                    `bson:"dispenseInterval,omitempty" json:"dispenseInterval,omitempty"`
	ValidityPeriod         *Period                                      `bson:"validityPeriod,omitempty" json:"validityPeriod,omitempty"`
	NumberOfRepeatsAllowed *int                                         `bson:"numberOfRepeatsAllowed,omitempty" json:"numberOfRepeatsAllowed,omitempty"`
	Quantity               *Quantity                                    `bson:"quantity,omitempty" json:"quantity,omitempty"`
	ExpectedSupplyDuration *Duration                                    `bson:"expectedSupplyDuration,omitempty" json:"expectedSupplyDuration,omitempty"`
	Performer              *Reference                                   `bson:"performer,omitempty" json:"performer,omitempty"`
}

// Indicates the quantity or duration for the first dispense of the medication.
// If populating this element, either the quantity or the duration must be included.
type MedicationRequestDispenseRequestInitialFill struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Quantity          *Quantity   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Duration          *Duration   `bson:"duration,omitempty" json:"duration,omitempty"`
}

// Indicates whether or not substitution can or should be part of the dispense. In some cases, substitution must happen, in other cases substitution must not happen. This block explains the prescriber's intent. If nothing is specified substitution may be done.
type MedicationRequestSubstitution struct {
	Id                     *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	AllowedBoolean         bool             `bson:"allowedBoolean" json:"allowedBoolean"`
	AllowedCodeableConcept CodeableConcept  `bson:"allowedCodeableConcept" json:"allowedCodeableConcept"`
	Reason                 *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
}

// This function returns resource reference information
func (r MedicationRequest) ResourceRef() (string, *string) {
	return "MedicationRequest", r.Id
}

// This function returns resource reference information
func (r MedicationRequest) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMedicationRequest MedicationRequest

// MarshalJSON marshals the given MedicationRequest as JSON into a byte slice
func (r MedicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationRequest: OtherMedicationRequest(r),
		ResourceType:           "MedicationRequest",
	})
}

// UnmarshalMedicationRequest unmarshals a MedicationRequest.
func UnmarshalMedicationRequest(b []byte) (MedicationRequest, error) {
	var medicationRequest MedicationRequest
	if err := json.Unmarshal(b, &medicationRequest); err != nil {
		return medicationRequest, err
	}
	return medicationRequest, nil
}
