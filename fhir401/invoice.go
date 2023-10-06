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

// Invoice is documented here http://hl7.org/fhir/StructureDefinition/Invoice
// Invoice containing collected ChargeItems from an Account with calculated individual and total price for Billing purpose.
type Invoice struct {
	Id                  *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	Contained           []json.RawMessage               `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension           []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              InvoiceStatus                   `bson:"status" json:"status"`
	CancelledReason     *string                         `bson:"cancelledReason,omitempty" json:"cancelledReason,omitempty"`
	Type                *CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	Subject             *Reference                      `bson:"subject,omitempty" json:"subject,omitempty"`
	Recipient           *Reference                      `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Date                *string                         `bson:"date,omitempty" json:"date,omitempty"`
	Participant         []InvoiceParticipant            `bson:"participant,omitempty" json:"participant,omitempty"`
	Issuer              *Reference                      `bson:"issuer,omitempty" json:"issuer,omitempty"`
	Account             *Reference                      `bson:"account,omitempty" json:"account,omitempty"`
	LineItem            []InvoiceLineItem               `bson:"lineItem,omitempty" json:"lineItem,omitempty"`
	TotalPriceComponent []InvoiceLineItemPriceComponent `bson:"totalPriceComponent,omitempty" json:"totalPriceComponent,omitempty"`
	TotalNet            *Money                          `bson:"totalNet,omitempty" json:"totalNet,omitempty"`
	TotalGross          *Money                          `bson:"totalGross,omitempty" json:"totalGross,omitempty"`
	PaymentTerms        *string                         `bson:"paymentTerms,omitempty" json:"paymentTerms,omitempty"`
	Note                []Annotation                    `bson:"note,omitempty" json:"note,omitempty"`
}

// Indicates who or what performed or participated in the charged service.
type InvoiceParticipant struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Actor             Reference        `bson:"actor" json:"actor"`
}

// Each line item represents one charge for goods and services rendered. Details such as date, code and amount are found in the referenced ChargeItem resource.
type InvoiceLineItem struct {
	Id                        *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension                 []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence                  *int                            `bson:"sequence,omitempty" json:"sequence,omitempty"`
	ChargeItemReference       Reference                       `bson:"chargeItemReference" json:"chargeItemReference"`
	ChargeItemCodeableConcept CodeableConcept                 `bson:"chargeItemCodeableConcept" json:"chargeItemCodeableConcept"`
	PriceComponent            []InvoiceLineItemPriceComponent `bson:"priceComponent,omitempty" json:"priceComponent,omitempty"`
}

// The price for a ChargeItem may be calculated as a base price with surcharges/deductions that apply in certain conditions. A ChargeItemDefinition resource that defines the prices, factors and conditions that apply to a billing code is currently under development. The priceComponent element can be used to offer transparency to the recipient of the Invoice as to how the prices have been calculated.
type InvoiceLineItemPriceComponent struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              InvoicePriceComponentType `bson:"type" json:"type"`
	Code              *CodeableConcept          `bson:"code,omitempty" json:"code,omitempty"`
	Factor            *json.Number              `bson:"factor,omitempty" json:"factor,omitempty"`
	Amount            *Money                    `bson:"amount,omitempty" json:"amount,omitempty"`
}

// This function returns resource reference information
func (r Invoice) ResourceRef() (string, *string) {
	return "Invoice", r.Id
}

// This function returns resource reference information
func (r Invoice) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherInvoice Invoice

// MarshalJSON marshals the given Invoice as JSON into a byte slice
func (r Invoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherInvoice
		ResourceType string `json:"resourceType"`
	}{
		OtherInvoice: OtherInvoice(r),
		ResourceType: "Invoice",
	})
}

// UnmarshalInvoice unmarshals a Invoice.
func UnmarshalInvoice(b []byte) (Invoice, error) {
	var invoice Invoice
	if err := json.Unmarshal(b, &invoice); err != nil {
		return invoice, err
	}
	return invoice, nil
}
