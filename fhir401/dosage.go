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

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// Dosage is documented here http://hl7.org/fhir/StructureDefinition/Dosage
// Base StructureDefinition for Dosage Type: Indicates how the medication is/was taken or should be taken by the patient.
type Dosage struct {
	Id                       *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension                []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence                 *int                `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Text                     *string             `bson:"text,omitempty" json:"text,omitempty"`
	AdditionalInstruction    []CodeableConcept   `bson:"additionalInstruction,omitempty" json:"additionalInstruction,omitempty"`
	PatientInstruction       *string             `bson:"patientInstruction,omitempty" json:"patientInstruction,omitempty"`
	Timing                   *Timing             `bson:"timing,omitempty" json:"timing,omitempty"`
	Site                     *CodeableConcept    `bson:"site,omitempty" json:"site,omitempty"`
	Route                    *CodeableConcept    `bson:"route,omitempty" json:"route,omitempty"`
	Method                   *CodeableConcept    `bson:"method,omitempty" json:"method,omitempty"`
	DoseAndRate              []DosageDoseAndRate `bson:"doseAndRate,omitempty" json:"doseAndRate,omitempty"`
	MaxDosePerPeriod         *Ratio              `bson:"maxDosePerPeriod,omitempty" json:"maxDosePerPeriod,omitempty"`
	MaxDosePerAdministration *Quantity           `bson:"maxDosePerAdministration,omitempty" json:"maxDosePerAdministration,omitempty"`
	MaxDosePerLifetime       *Quantity           `bson:"maxDosePerLifetime,omitempty" json:"maxDosePerLifetime,omitempty"`
}

// The amount of medication administered.
type DosageDoseAndRate struct {
	Id        *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	Type      *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
}
