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

// PaymentReconciliation is documented here http://hl7.org/fhir/StructureDefinition/PaymentReconciliation
// This resource provides the details including amount of a payment and allocates the payment items being paid.
type PaymentReconciliation struct {
	Id                *string                            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                         `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            FinancialResourceStatusCodes       `bson:"status" json:"status"`
	Period            *Period                            `bson:"period,omitempty" json:"period,omitempty"`
	Created           string                             `bson:"created" json:"created"`
	PaymentIssuer     *Reference                         `bson:"paymentIssuer,omitempty" json:"paymentIssuer,omitempty"`
	Request           *Reference                         `bson:"request,omitempty" json:"request,omitempty"`
	Requestor         *Reference                         `bson:"requestor,omitempty" json:"requestor,omitempty"`
	Outcome           *ClaimProcessingCodes              `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition       *string                            `bson:"disposition,omitempty" json:"disposition,omitempty"`
	PaymentDate       string                             `bson:"paymentDate" json:"paymentDate"`
	PaymentAmount     Money                              `bson:"paymentAmount" json:"paymentAmount"`
	PaymentIdentifier *Identifier                        `bson:"paymentIdentifier,omitempty" json:"paymentIdentifier,omitempty"`
	Detail            []PaymentReconciliationDetail      `bson:"detail,omitempty" json:"detail,omitempty"`
	FormCode          *CodeableConcept                   `bson:"formCode,omitempty" json:"formCode,omitempty"`
	ProcessNote       []PaymentReconciliationProcessNote `bson:"processNote,omitempty" json:"processNote,omitempty"`
}

// Distribution of the payment amount for a previously acknowledged payable.
type PaymentReconciliationDetail struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Predecessor       *Identifier     `bson:"predecessor,omitempty" json:"predecessor,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
	Request           *Reference      `bson:"request,omitempty" json:"request,omitempty"`
	Submitter         *Reference      `bson:"submitter,omitempty" json:"submitter,omitempty"`
	Response          *Reference      `bson:"response,omitempty" json:"response,omitempty"`
	Date              *string         `bson:"date,omitempty" json:"date,omitempty"`
	Responsible       *Reference      `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Payee             *Reference      `bson:"payee,omitempty" json:"payee,omitempty"`
	Amount            *Money          `bson:"amount,omitempty" json:"amount,omitempty"`
}

// A note that describes or explains the processing in a human readable form.
type PaymentReconciliationProcessNote struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *NoteType   `bson:"type,omitempty" json:"type,omitempty"`
	Text              *string     `bson:"text,omitempty" json:"text,omitempty"`
}
type OtherPaymentReconciliation PaymentReconciliation

// MarshalJSON marshals the given PaymentReconciliation as JSON into a byte slice
func (r PaymentReconciliation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPaymentReconciliation
		ResourceType string `json:"resourceType"`
	}{
		OtherPaymentReconciliation: OtherPaymentReconciliation(r),
		ResourceType:               "PaymentReconciliation",
	})
}

// UnmarshalPaymentReconciliation unmarshals a PaymentReconciliation.
func UnmarshalPaymentReconciliation(b []byte) (PaymentReconciliation, error) {
	var paymentReconciliation PaymentReconciliation
	if err := json.Unmarshal(b, &paymentReconciliation); err != nil {
		return paymentReconciliation, err
	}
	return paymentReconciliation, nil
}
