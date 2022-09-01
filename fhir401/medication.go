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

// Medication is documented here http://hl7.org/fhir/StructureDefinition/Medication
// This resource is primarily used for the identification and definition of a medication for the purposes of prescribing, dispensing, and administering a medication as well as for making statements about medication use.
type Medication struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code              *CodeableConcept       `bson:"code,omitempty" json:"code,omitempty"`
	Status            *string                `bson:"status,omitempty" json:"status,omitempty"`
	Manufacturer      *Reference             `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Form              *CodeableConcept       `bson:"form,omitempty" json:"form,omitempty"`
	Amount            *Ratio                 `bson:"amount,omitempty" json:"amount,omitempty"`
	Ingredient        []MedicationIngredient `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	Batch             *MedicationBatch       `bson:"batch,omitempty" json:"batch,omitempty"`
}

// Identifies a particular constituent of interest in the product.
// The ingredients need not be a complete list.  If an ingredient is not specified, this does not indicate whether an ingredient is present or absent.  If an ingredient is specified it does not mean that all ingredients are specified.  It is possible to specify both inactive and active ingredients.
type MedicationIngredient struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	IsActive          *bool       `bson:"isActive,omitempty" json:"isActive,omitempty"`
	Strength          *Ratio      `bson:"strength,omitempty" json:"strength,omitempty"`
}

// Information that only applies to packages (not products).
type MedicationBatch struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LotNumber         *string     `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	ExpirationDate    *string     `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
}
type OtherMedication Medication

// MarshalJSON marshals the given Medication as JSON into a byte slice
func (r Medication) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedication
		ResourceType string `json:"resourceType"`
	}{
		OtherMedication: OtherMedication(r),
		ResourceType:    "Medication",
	})
}

// UnmarshalMedication unmarshals a Medication.
func UnmarshalMedication(b []byte) (Medication, error) {
	var medication Medication
	if err := json.Unmarshal(b, &medication); err != nil {
		return medication, err
	}
	return medication, nil
}
