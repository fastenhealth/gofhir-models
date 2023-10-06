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

// ClaimResponse is documented here http://hl7.org/fhir/StructureDefinition/ClaimResponse
// This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse struct {
	Id                   *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	Contained            []json.RawMessage               `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension            []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               FinancialResourceStatusCodes    `bson:"status" json:"status"`
	Type                 CodeableConcept                 `bson:"type" json:"type"`
	SubType              *CodeableConcept                `bson:"subType,omitempty" json:"subType,omitempty"`
	Use                  Use                             `bson:"use" json:"use"`
	Patient              Reference                       `bson:"patient" json:"patient"`
	Created              string                          `bson:"created" json:"created"`
	Insurer              Reference                       `bson:"insurer" json:"insurer"`
	Requestor            *Reference                      `bson:"requestor,omitempty" json:"requestor,omitempty"`
	Request              *Reference                      `bson:"request,omitempty" json:"request,omitempty"`
	Outcome              ClaimProcessingCodes            `bson:"outcome" json:"outcome"`
	Disposition          *string                         `bson:"disposition,omitempty" json:"disposition,omitempty"`
	PreAuthRef           *string                         `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	PreAuthPeriod        *Period                         `bson:"preAuthPeriod,omitempty" json:"preAuthPeriod,omitempty"`
	PayeeType            *CodeableConcept                `bson:"payeeType,omitempty" json:"payeeType,omitempty"`
	Item                 []ClaimResponseItem             `bson:"item,omitempty" json:"item,omitempty"`
	AddItem              []ClaimResponseAddItem          `bson:"addItem,omitempty" json:"addItem,omitempty"`
	Adjudication         []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Total                []ClaimResponseTotal            `bson:"total,omitempty" json:"total,omitempty"`
	Payment              *ClaimResponsePayment           `bson:"payment,omitempty" json:"payment,omitempty"`
	FundsReserve         *CodeableConcept                `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty"`
	FormCode             *CodeableConcept                `bson:"formCode,omitempty" json:"formCode,omitempty"`
	Form                 *Attachment                     `bson:"form,omitempty" json:"form,omitempty"`
	ProcessNote          []ClaimResponseProcessNote      `bson:"processNote,omitempty" json:"processNote,omitempty"`
	CommunicationRequest []Reference                     `bson:"communicationRequest,omitempty" json:"communicationRequest,omitempty"`
	Insurance            []ClaimResponseInsurance        `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Error                []ClaimResponseError            `bson:"error,omitempty" json:"error,omitempty"`
}

// A claim line. Either a simple (a product or service) or a 'group' of details which can also be a simple items or groups of sub-details.
type ClaimResponseItem struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ItemSequence      int                             `bson:"itemSequence" json:"itemSequence"`
	NoteNumber        []int                           `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `bson:"adjudication" json:"adjudication"`
	Detail            []ClaimResponseItemDetail       `bson:"detail,omitempty" json:"detail,omitempty"`
}

// If this item is a group then the values here are a summary of the adjudication of the detail items. If this item is a simple product or service then this is the result of the adjudication of this item.
type ClaimResponseItemAdjudication struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          CodeableConcept  `bson:"category" json:"category"`
	Reason            *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount            *Money           `bson:"amount,omitempty" json:"amount,omitempty"`
	Value             *json.Number     `bson:"value,omitempty" json:"value,omitempty"`
}

// A claim detail. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
type ClaimResponseItemDetail struct {
	Id                *string                            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DetailSequence    int                                `bson:"detailSequence" json:"detailSequence"`
	NoteNumber        []int                              `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication    `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail         []ClaimResponseItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

// A sub-detail adjudication of a simple product or service.
type ClaimResponseItemDetailSubDetail struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SubDetailSequence int                             `bson:"subDetailSequence" json:"subDetailSequence"`
	NoteNumber        []int                           `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

// The first-tier service adjudications for payor added product or service lines.
type ClaimResponseAddItem struct {
	Id                      *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension               []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ItemSequence            []int                           `bson:"itemSequence,omitempty" json:"itemSequence,omitempty"`
	DetailSequence          []int                           `bson:"detailSequence,omitempty" json:"detailSequence,omitempty"`
	SubdetailSequence       []int                           `bson:"subdetailSequence,omitempty" json:"subdetailSequence,omitempty"`
	Provider                []Reference                     `bson:"provider,omitempty" json:"provider,omitempty"`
	ProductOrService        CodeableConcept                 `bson:"productOrService" json:"productOrService"`
	Modifier                []CodeableConcept               `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept               `bson:"programCode,omitempty" json:"programCode,omitempty"`
	ServicedDate            *string                         `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod          *Period                         `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept                `bson:"locationCodeableConcept,omitempty" json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                        `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference       *Reference                      `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
	Quantity                *Quantity                       `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice               *Money                          `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                  *json.Number                    `bson:"factor,omitempty" json:"factor,omitempty"`
	Net                     *Money                          `bson:"net,omitempty" json:"net,omitempty"`
	BodySite                *CodeableConcept                `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite                 []CodeableConcept               `bson:"subSite,omitempty" json:"subSite,omitempty"`
	NoteNumber              []int                           `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication            []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail                  []ClaimResponseAddItemDetail    `bson:"detail,omitempty" json:"detail,omitempty"`
}

// The second-tier service adjudications for payor added services.
type ClaimResponseAddItemDetail struct {
	Id                *string                               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ProductOrService  CodeableConcept                       `bson:"productOrService" json:"productOrService"`
	Modifier          []CodeableConcept                     `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Quantity          *Quantity                             `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                                `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *json.Number                          `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money                                `bson:"net,omitempty" json:"net,omitempty"`
	NoteNumber        []int                                 `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication       `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail         []ClaimResponseAddItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

// The third-tier service adjudications for payor added services.
type ClaimResponseAddItemDetailSubDetail struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ProductOrService  CodeableConcept                 `bson:"productOrService" json:"productOrService"`
	Modifier          []CodeableConcept               `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Quantity          *Quantity                       `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                          `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *json.Number                    `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money                          `bson:"net,omitempty" json:"net,omitempty"`
	NoteNumber        []int                           `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

// Categorized monetary totals for the adjudication.
// Totals for amounts submitted, co-pays, benefits payable etc.
type ClaimResponseTotal struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          CodeableConcept `bson:"category" json:"category"`
	Amount            Money           `bson:"amount" json:"amount"`
}

// Payment details for the adjudication of the claim.
type ClaimResponsePayment struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept  `bson:"type" json:"type"`
	Adjustment        *Money           `bson:"adjustment,omitempty" json:"adjustment,omitempty"`
	AdjustmentReason  *CodeableConcept `bson:"adjustmentReason,omitempty" json:"adjustmentReason,omitempty"`
	Date              *string          `bson:"date,omitempty" json:"date,omitempty"`
	Amount            Money            `bson:"amount" json:"amount"`
	Identifier        *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
}

// A note that describes or explains adjudication results in a human readable form.
type ClaimResponseProcessNote struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Number            *int             `bson:"number,omitempty" json:"number,omitempty"`
	Type              *NoteType        `bson:"type,omitempty" json:"type,omitempty"`
	Text              string           `bson:"text" json:"text"`
	Language          *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
}

// Financial instruments for reimbursement for the health care products and services specified on the claim.
// All insurance coverages for the patient which may be applicable for reimbursement, of the products and services listed in the claim, are typically provided in the claim to allow insurers to confirm the ordering of the insurance coverages relative to local 'coordination of benefit' rules. One coverage (and only one) with 'focal=true' is to be used in the adjudication of this claim. Coverages appearing before the focal Coverage in the list, and where 'subrogation=false', should provide a reference to the ClaimResponse containing the adjudication results of the prior claim.
type ClaimResponseInsurance struct {
	Id                  *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence            int         `bson:"sequence" json:"sequence"`
	Focal               bool        `bson:"focal" json:"focal"`
	Coverage            Reference   `bson:"coverage" json:"coverage"`
	BusinessArrangement *string     `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	ClaimResponse       *Reference  `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
}

// Errors encountered during the processing of the adjudication.
// If the request contains errors then an error element should be provided and no adjudication related sections (item, addItem, or payment) should be present.
type ClaimResponseError struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ItemSequence      *int            `bson:"itemSequence,omitempty" json:"itemSequence,omitempty"`
	DetailSequence    *int            `bson:"detailSequence,omitempty" json:"detailSequence,omitempty"`
	SubDetailSequence *int            `bson:"subDetailSequence,omitempty" json:"subDetailSequence,omitempty"`
	Code              CodeableConcept `bson:"code" json:"code"`
}

// This function returns resource reference information
func (r ClaimResponse) ResourceRef() (string, *string) {
	return "ClaimResponse", r.Id
}

// This function returns resource reference information
func (r ClaimResponse) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherClaimResponse ClaimResponse

// MarshalJSON marshals the given ClaimResponse as JSON into a byte slice
func (r ClaimResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClaimResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherClaimResponse: OtherClaimResponse(r),
		ResourceType:       "ClaimResponse",
	})
}

// UnmarshalClaimResponse unmarshals a ClaimResponse.
func UnmarshalClaimResponse(b []byte) (ClaimResponse, error) {
	var claimResponse ClaimResponse
	if err := json.Unmarshal(b, &claimResponse); err != nil {
		return claimResponse, err
	}
	return claimResponse, nil
}
