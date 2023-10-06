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

// Parameters is documented here http://hl7.org/fhir/StructureDefinition/Parameters
// This resource is a non-persisted resource used to pass information into and back from an [operation](operations.html). It has no other use, and there is no RESTful endpoint associated with it.
type Parameters struct {
	Id            *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta          *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language      *string               `bson:"language,omitempty" json:"language,omitempty"`
	Parameter     []ParametersParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
}

// A parameter passed to or received from the operation.
type ParametersParameter struct {
	Id                       *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension                []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name                     string                `bson:"name" json:"name"`
	ValueBase64Binary        *string               `bson:"valueBase64Binary,omitempty" json:"valueBase64Binary,omitempty"`
	ValueBoolean             *bool                 `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueCanonical           *string               `bson:"valueCanonical,omitempty" json:"valueCanonical,omitempty"`
	ValueCode                *string               `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueDate                *string               `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime            *string               `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueDecimal             *json.Number          `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueId                  *string               `bson:"valueId,omitempty" json:"valueId,omitempty"`
	ValueInstant             *string               `bson:"valueInstant,omitempty" json:"valueInstant,omitempty"`
	ValueInteger             *int                  `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueMarkdown            *string               `bson:"valueMarkdown,omitempty" json:"valueMarkdown,omitempty"`
	ValueOid                 *string               `bson:"valueOid,omitempty" json:"valueOid,omitempty"`
	ValuePositiveInt         *int                  `bson:"valuePositiveInt,omitempty" json:"valuePositiveInt,omitempty"`
	ValueString              *string               `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueTime                *string               `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueUnsignedInt         *int                  `bson:"valueUnsignedInt,omitempty" json:"valueUnsignedInt,omitempty"`
	ValueUri                 *string               `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueUrl                 *string               `bson:"valueUrl,omitempty" json:"valueUrl,omitempty"`
	ValueUuid                *string               `bson:"valueUuid,omitempty" json:"valueUuid,omitempty"`
	ValueAddress             *Address              `bson:"valueAddress,omitempty" json:"valueAddress,omitempty"`
	ValueAge                 *Age                  `bson:"valueAge,omitempty" json:"valueAge,omitempty"`
	ValueAnnotation          *Annotation           `bson:"valueAnnotation,omitempty" json:"valueAnnotation,omitempty"`
	ValueAttachment          *Attachment           `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueCodeableConcept     *CodeableConcept      `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueCoding              *Coding               `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueContactPoint        *ContactPoint         `bson:"valueContactPoint,omitempty" json:"valueContactPoint,omitempty"`
	ValueCount               *Count                `bson:"valueCount,omitempty" json:"valueCount,omitempty"`
	ValueDistance            *Distance             `bson:"valueDistance,omitempty" json:"valueDistance,omitempty"`
	ValueDuration            *Duration             `bson:"valueDuration,omitempty" json:"valueDuration,omitempty"`
	ValueHumanName           *HumanName            `bson:"valueHumanName,omitempty" json:"valueHumanName,omitempty"`
	ValueIdentifier          *Identifier           `bson:"valueIdentifier,omitempty" json:"valueIdentifier,omitempty"`
	ValueMoney               *Money                `bson:"valueMoney,omitempty" json:"valueMoney,omitempty"`
	ValuePeriod              *Period               `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValueQuantity            *Quantity             `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange               *Range                `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio               *Ratio                `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueReference           *Reference            `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	ValueSampledData         *SampledData          `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueSignature           *Signature            `bson:"valueSignature,omitempty" json:"valueSignature,omitempty"`
	ValueTiming              *Timing               `bson:"valueTiming,omitempty" json:"valueTiming,omitempty"`
	ValueContactDetail       *ContactDetail        `bson:"valueContactDetail,omitempty" json:"valueContactDetail,omitempty"`
	ValueContributor         *Contributor          `bson:"valueContributor,omitempty" json:"valueContributor,omitempty"`
	ValueDataRequirement     *DataRequirement      `bson:"valueDataRequirement,omitempty" json:"valueDataRequirement,omitempty"`
	ValueExpression          *Expression           `bson:"valueExpression,omitempty" json:"valueExpression,omitempty"`
	ValueParameterDefinition *ParameterDefinition  `bson:"valueParameterDefinition,omitempty" json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact     *RelatedArtifact      `bson:"valueRelatedArtifact,omitempty" json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition   *TriggerDefinition    `bson:"valueTriggerDefinition,omitempty" json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext        *UsageContext         `bson:"valueUsageContext,omitempty" json:"valueUsageContext,omitempty"`
	ValueDosage              *Dosage               `bson:"valueDosage,omitempty" json:"valueDosage,omitempty"`
	ValueMeta                *Meta                 `bson:"valueMeta,omitempty" json:"valueMeta,omitempty"`
	Resource                 json.RawMessage       `bson:"resource,omitempty" json:"resource,omitempty"`
	Part                     []ParametersParameter `bson:"part,omitempty" json:"part,omitempty"`
}

// This function returns resource reference information
func (r Parameters) ResourceRef() (string, *string) {
	return "Parameters", r.Id
}

// This function returns resource reference information
func (r Parameters) ContainedResources() []json.RawMessage {
	return nil
}

type OtherParameters Parameters

// MarshalJSON marshals the given Parameters as JSON into a byte slice
func (r Parameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherParameters
		ResourceType string `json:"resourceType"`
	}{
		OtherParameters: OtherParameters(r),
		ResourceType:    "Parameters",
	})
}

// UnmarshalParameters unmarshals a Parameters.
func UnmarshalParameters(b []byte) (Parameters, error) {
	var parameters Parameters
	if err := json.Unmarshal(b, &parameters); err != nil {
		return parameters, err
	}
	return parameters, nil
}
