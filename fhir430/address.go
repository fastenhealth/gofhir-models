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

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// Address is documented here http://hl7.org/fhir/StructureDefinition/Address
// Base StructureDefinition for Address Type: An address expressed using postal conventions (as opposed to GPS or other location definition formats).  This data type may be used to convey addresses for use in delivering mail as well as for visiting locations which might not be valid for mail delivery.  There are a variety of postal address formats defined around the world.
type Address struct {
	Id         *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension  []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	Use        *AddressUse  `bson:"use,omitempty" json:"use,omitempty"`
	Type       *AddressType `bson:"type,omitempty" json:"type,omitempty"`
	Text       *string      `bson:"text,omitempty" json:"text,omitempty"`
	Line       []string     `bson:"line,omitempty" json:"line,omitempty"`
	City       *string      `bson:"city,omitempty" json:"city,omitempty"`
	District   *string      `bson:"district,omitempty" json:"district,omitempty"`
	State      *string      `bson:"state,omitempty" json:"state,omitempty"`
	PostalCode *string      `bson:"postalCode,omitempty" json:"postalCode,omitempty"`
	Country    *string      `bson:"country,omitempty" json:"country,omitempty"`
	Period     *Period      `bson:"period,omitempty" json:"period,omitempty"`
}
