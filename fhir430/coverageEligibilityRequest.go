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

// CoverageEligibilityRequest is documented here http://hl7.org/fhir/StructureDefinition/CoverageEligibilityRequest
// The CoverageEligibilityRequest provides patient and insurance coverage information to an insurer for them to respond, in the form of an CoverageEligibilityResponse, with information regarding whether the stated coverage is valid and in-force and optionally to provide the insurance details of the policy.
type CoverageEligibilityRequest struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                                    `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                                 `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            FinancialResourceStatusCodes               `bson:"status" json:"status"`
	Priority          *CodeableConcept                           `bson:"priority,omitempty" json:"priority,omitempty"`
	Purpose           []EligibilityRequestPurpose                `bson:"purpose" json:"purpose"`
	Patient           Reference                                  `bson:"patient" json:"patient"`
	Created           string                                     `bson:"created" json:"created"`
	Enterer           *Reference                                 `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Provider          *Reference                                 `bson:"provider,omitempty" json:"provider,omitempty"`
	Insurer           Reference                                  `bson:"insurer" json:"insurer"`
	Facility          *Reference                                 `bson:"facility,omitempty" json:"facility,omitempty"`
	SupportingInfo    []CoverageEligibilityRequestSupportingInfo `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Insurance         []CoverageEligibilityRequestInsurance      `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Item              []CoverageEligibilityRequestItem           `bson:"item,omitempty" json:"item,omitempty"`
}

// Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues.
// Often there are multiple jurisdiction specific valuesets which are required.
type CoverageEligibilityRequestSupportingInfo struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int         `bson:"sequence" json:"sequence"`
	Information       Reference   `bson:"information" json:"information"`
	AppliesToAll      *bool       `bson:"appliesToAll,omitempty" json:"appliesToAll,omitempty"`
}

// Financial instruments for reimbursement for the health care products and services.
// All insurance coverages for the patient which may be applicable for reimbursement, of the products and services listed in the claim, are typically provided in the claim to allow insurers to confirm the ordering of the insurance coverages relative to local 'coordination of benefit' rules. One coverage (and only one) with 'focal=true' is to be used in the adjudication of this claim. Coverages appearing before the focal Coverage in the list, and where 'subrogation=false', should provide a reference to the ClaimResponse containing the adjudication results of the prior claim.
type CoverageEligibilityRequestInsurance struct {
	Id                  *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Focal               *bool       `bson:"focal,omitempty" json:"focal,omitempty"`
	Coverage            Reference   `bson:"coverage" json:"coverage"`
	BusinessArrangement *string     `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
}

// Service categories or billable services for which benefit details and/or an authorization prior to service delivery may be required by the payor.
type CoverageEligibilityRequestItem struct {
	Id                     *string                                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension                               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SupportingInfoSequence []int                                     `bson:"supportingInfoSequence,omitempty" json:"supportingInfoSequence,omitempty"`
	Category               *CodeableConcept                          `bson:"category,omitempty" json:"category,omitempty"`
	ProductOrService       *CodeableConcept                          `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	Modifier               []CodeableConcept                         `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Provider               *Reference                                `bson:"provider,omitempty" json:"provider,omitempty"`
	Quantity               *Quantity                                 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice              *Money                                    `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Facility               *Reference                                `bson:"facility,omitempty" json:"facility,omitempty"`
	Diagnosis              []CoverageEligibilityRequestItemDiagnosis `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Detail                 []Reference                               `bson:"detail,omitempty" json:"detail,omitempty"`
}

// Patient diagnosis for which care is sought.
type CoverageEligibilityRequestItemDiagnosis struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}
type OtherCoverageEligibilityRequest CoverageEligibilityRequest

// MarshalJSON marshals the given CoverageEligibilityRequest as JSON into a byte slice
func (r CoverageEligibilityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoverageEligibilityRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherCoverageEligibilityRequest: OtherCoverageEligibilityRequest(r),
		ResourceType:                    "CoverageEligibilityRequest",
	})
}

// UnmarshalCoverageEligibilityRequest unmarshals a CoverageEligibilityRequest.
func UnmarshalCoverageEligibilityRequest(b []byte) (CoverageEligibilityRequest, error) {
	var coverageEligibilityRequest CoverageEligibilityRequest
	if err := json.Unmarshal(b, &coverageEligibilityRequest); err != nil {
		return coverageEligibilityRequest, err
	}
	return coverageEligibilityRequest, nil
}
