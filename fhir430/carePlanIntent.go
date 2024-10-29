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

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// CarePlanIntent is documented here http://hl7.org/fhir/ValueSet/care-plan-intent
type CarePlanIntent int

const (
	CarePlanIntentProposal CarePlanIntent = iota
	CarePlanIntentPlan
	CarePlanIntentDirective
	CarePlanIntentOrder
	CarePlanIntentOriginalOrder
	CarePlanIntentReflexOrder
	CarePlanIntentFillerOrder
	CarePlanIntentInstanceOrder
	CarePlanIntentOption
)

func (code CarePlanIntent) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CarePlanIntent) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "proposal":
		*code = CarePlanIntentProposal
	case "plan":
		*code = CarePlanIntentPlan
	case "directive":
		*code = CarePlanIntentDirective
	case "order":
		*code = CarePlanIntentOrder
	case "original-order":
		*code = CarePlanIntentOriginalOrder
	case "reflex-order":
		*code = CarePlanIntentReflexOrder
	case "filler-order":
		*code = CarePlanIntentFillerOrder
	case "instance-order":
		*code = CarePlanIntentInstanceOrder
	case "option":
		*code = CarePlanIntentOption
	default:
		return fmt.Errorf("unknown CarePlanIntent code `%s`", s)
	}
	return nil
}
func (code CarePlanIntent) String() string {
	return code.Code()
}
func (code CarePlanIntent) Code() string {
	switch code {
	case CarePlanIntentProposal:
		return "proposal"
	case CarePlanIntentPlan:
		return "plan"
	case CarePlanIntentDirective:
		return "directive"
	case CarePlanIntentOrder:
		return "order"
	case CarePlanIntentOriginalOrder:
		return "original-order"
	case CarePlanIntentReflexOrder:
		return "reflex-order"
	case CarePlanIntentFillerOrder:
		return "filler-order"
	case CarePlanIntentInstanceOrder:
		return "instance-order"
	case CarePlanIntentOption:
		return "option"
	}
	return "<unknown>"
}
func (code CarePlanIntent) Display() string {
	switch code {
	case CarePlanIntentProposal:
		return "Proposal"
	case CarePlanIntentPlan:
		return "Plan"
	case CarePlanIntentDirective:
		return "Directive"
	case CarePlanIntentOrder:
		return "Order"
	case CarePlanIntentOriginalOrder:
		return "Original Order"
	case CarePlanIntentReflexOrder:
		return "Reflex Order"
	case CarePlanIntentFillerOrder:
		return "Filler Order"
	case CarePlanIntentInstanceOrder:
		return "Instance Order"
	case CarePlanIntentOption:
		return "Option"
	}
	return "<unknown>"
}
func (code CarePlanIntent) Definition() string {
	switch code {
	case CarePlanIntentProposal:
		return "The request is a suggestion made by someone/something that does not have an intention to ensure it occurs and without providing an authorization to act."
	case CarePlanIntentPlan:
		return "The request represents an intention to ensure something occurs without providing an authorization for others to act."
	case CarePlanIntentDirective:
		return "The request represents a legally binding instruction authored by a Patient or RelatedPerson."
	case CarePlanIntentOrder:
		return "The request represents a request/demand and authorization for action by a Practitioner."
	case CarePlanIntentOriginalOrder:
		return "The request represents an original authorization for action."
	case CarePlanIntentReflexOrder:
		return "The request represents an automatically generated supplemental authorization for action based on a parent authorization together with initial results of the action taken against that parent authorization."
	case CarePlanIntentFillerOrder:
		return "The request represents the view of an authorization instantiated by a fulfilling system representing the details of the fulfiller's intention to act upon a submitted order."
	case CarePlanIntentInstanceOrder:
		return "An order created in fulfillment of a broader order that represents the authorization for a single activity occurrence.  E.g. The administration of a single dose of a drug."
	case CarePlanIntentOption:
		return "The request represents a component or option for a RequestGroup that establishes timing, conditionality and/or other constraints among a set of requests.  Refer to [[[RequestGroup]]] for additional information on how this status is used."
	}
	return "<unknown>"
}
