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

// ParameterDefinition is documented here http://hl7.org/fhir/StructureDefinition/ParameterDefinition
// Base StructureDefinition for ParameterDefinition Type: The parameters to the module. This collection specifies both the input and output parameters. Input parameters are provided by the caller as part of the $evaluate operation. Output parameters are included in the GuidanceResponse.
type ParameterDefinition struct {
	Id            *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension     []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	Name          *string               `bson:"name,omitempty" json:"name,omitempty"`
	Use           OperationParameterUse `bson:"use" json:"use"`
	Min           *int                  `bson:"min,omitempty" json:"min,omitempty"`
	Max           *string               `bson:"max,omitempty" json:"max,omitempty"`
	Documentation *string               `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Type          string                `bson:"type" json:"type"`
	Profile       *string               `bson:"profile,omitempty" json:"profile,omitempty"`
}
