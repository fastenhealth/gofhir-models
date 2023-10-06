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

// MolecularSequence is documented here http://hl7.org/fhir/StructureDefinition/MolecularSequence
// Raw data describing a biological sequence.
type MolecularSequence struct {
	Id                *string                             `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                             `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                          `bson:"text,omitempty" json:"text,omitempty"`
	Contained         []json.RawMessage                   `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension                         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type              *string                             `bson:"type,omitempty" json:"type,omitempty"`
	CoordinateSystem  int                                 `bson:"coordinateSystem" json:"coordinateSystem"`
	Patient           *Reference                          `bson:"patient,omitempty" json:"patient,omitempty"`
	Specimen          *Reference                          `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Device            *Reference                          `bson:"device,omitempty" json:"device,omitempty"`
	Performer         *Reference                          `bson:"performer,omitempty" json:"performer,omitempty"`
	Quantity          *Quantity                           `bson:"quantity,omitempty" json:"quantity,omitempty"`
	ReferenceSeq      *MolecularSequenceReferenceSeq      `bson:"referenceSeq,omitempty" json:"referenceSeq,omitempty"`
	Variant           []MolecularSequenceVariant          `bson:"variant,omitempty" json:"variant,omitempty"`
	ObservedSeq       *string                             `bson:"observedSeq,omitempty" json:"observedSeq,omitempty"`
	Quality           []MolecularSequenceQuality          `bson:"quality,omitempty" json:"quality,omitempty"`
	ReadCoverage      *int                                `bson:"readCoverage,omitempty" json:"readCoverage,omitempty"`
	Repository        []MolecularSequenceRepository       `bson:"repository,omitempty" json:"repository,omitempty"`
	Pointer           []Reference                         `bson:"pointer,omitempty" json:"pointer,omitempty"`
	StructureVariant  []MolecularSequenceStructureVariant `bson:"structureVariant,omitempty" json:"structureVariant,omitempty"`
}

// A sequence that is used as a reference to describe variants that are present in a sequence analyzed.
type MolecularSequenceReferenceSeq struct {
	Id                  *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Chromosome          *CodeableConcept `bson:"chromosome,omitempty" json:"chromosome,omitempty"`
	GenomeBuild         *string          `bson:"genomeBuild,omitempty" json:"genomeBuild,omitempty"`
	Orientation         *string          `bson:"orientation,omitempty" json:"orientation,omitempty"`
	ReferenceSeqId      *CodeableConcept `bson:"referenceSeqId,omitempty" json:"referenceSeqId,omitempty"`
	ReferenceSeqPointer *Reference       `bson:"referenceSeqPointer,omitempty" json:"referenceSeqPointer,omitempty"`
	ReferenceSeqString  *string          `bson:"referenceSeqString,omitempty" json:"referenceSeqString,omitempty"`
	Strand              *string          `bson:"strand,omitempty" json:"strand,omitempty"`
	WindowStart         *int             `bson:"windowStart,omitempty" json:"windowStart,omitempty"`
	WindowEnd           *int             `bson:"windowEnd,omitempty" json:"windowEnd,omitempty"`
}

// The definition of variant here originates from Sequence ontology ([variant_of](http://www.sequenceontology.org/browser/current_svn/term/variant_of)). This element can represent amino acid or nucleic sequence change(including insertion,deletion,SNP,etc.)  It can represent some complex mutation or segment variation with the assist of CIGAR string.
type MolecularSequenceVariant struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Start             *int        `bson:"start,omitempty" json:"start,omitempty"`
	End               *int        `bson:"end,omitempty" json:"end,omitempty"`
	ObservedAllele    *string     `bson:"observedAllele,omitempty" json:"observedAllele,omitempty"`
	ReferenceAllele   *string     `bson:"referenceAllele,omitempty" json:"referenceAllele,omitempty"`
	Cigar             *string     `bson:"cigar,omitempty" json:"cigar,omitempty"`
	VariantPointer    *Reference  `bson:"variantPointer,omitempty" json:"variantPointer,omitempty"`
}

// An experimental feature attribute that defines the quality of the feature in a quantitative way, such as a phred quality score ([SO:0001686](http://www.sequenceontology.org/browser/current_svn/term/SO:0001686)).
type MolecularSequenceQuality struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              string                       `bson:"type" json:"type"`
	StandardSequence  *CodeableConcept             `bson:"standardSequence,omitempty" json:"standardSequence,omitempty"`
	Start             *int                         `bson:"start,omitempty" json:"start,omitempty"`
	End               *int                         `bson:"end,omitempty" json:"end,omitempty"`
	Score             *Quantity                    `bson:"score,omitempty" json:"score,omitempty"`
	Method            *CodeableConcept             `bson:"method,omitempty" json:"method,omitempty"`
	TruthTP           *json.Number                 `bson:"truthTP,omitempty" json:"truthTP,omitempty"`
	QueryTP           *json.Number                 `bson:"queryTP,omitempty" json:"queryTP,omitempty"`
	TruthFN           *json.Number                 `bson:"truthFN,omitempty" json:"truthFN,omitempty"`
	QueryFP           *json.Number                 `bson:"queryFP,omitempty" json:"queryFP,omitempty"`
	GtFP              *json.Number                 `bson:"gtFP,omitempty" json:"gtFP,omitempty"`
	Precision         *json.Number                 `bson:"precision,omitempty" json:"precision,omitempty"`
	Recall            *json.Number                 `bson:"recall,omitempty" json:"recall,omitempty"`
	FScore            *json.Number                 `bson:"fScore,omitempty" json:"fScore,omitempty"`
	Roc               *MolecularSequenceQualityRoc `bson:"roc,omitempty" json:"roc,omitempty"`
}

// Receiver Operator Characteristic (ROC) Curve  to give sensitivity/specificity tradeoff.
type MolecularSequenceQualityRoc struct {
	Id                *string       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Score             []int         `bson:"score,omitempty" json:"score,omitempty"`
	NumTP             []int         `bson:"numTP,omitempty" json:"numTP,omitempty"`
	NumFP             []int         `bson:"numFP,omitempty" json:"numFP,omitempty"`
	NumFN             []int         `bson:"numFN,omitempty" json:"numFN,omitempty"`
	Precision         []json.Number `bson:"precision,omitempty" json:"precision,omitempty"`
	Sensitivity       []json.Number `bson:"sensitivity,omitempty" json:"sensitivity,omitempty"`
	FMeasure          []json.Number `bson:"fMeasure,omitempty" json:"fMeasure,omitempty"`
}

// Configurations of the external repository. The repository shall store target's observedSeq or records related with target's observedSeq.
type MolecularSequenceRepository struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              string      `bson:"type" json:"type"`
	Url               *string     `bson:"url,omitempty" json:"url,omitempty"`
	Name              *string     `bson:"name,omitempty" json:"name,omitempty"`
	DatasetId         *string     `bson:"datasetId,omitempty" json:"datasetId,omitempty"`
	VariantsetId      *string     `bson:"variantsetId,omitempty" json:"variantsetId,omitempty"`
	ReadsetId         *string     `bson:"readsetId,omitempty" json:"readsetId,omitempty"`
}

// Information about chromosome structure variation.
type MolecularSequenceStructureVariant struct {
	Id                *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	VariantType       *CodeableConcept                        `bson:"variantType,omitempty" json:"variantType,omitempty"`
	Exact             *bool                                   `bson:"exact,omitempty" json:"exact,omitempty"`
	Length            *int                                    `bson:"length,omitempty" json:"length,omitempty"`
	Outer             *MolecularSequenceStructureVariantOuter `bson:"outer,omitempty" json:"outer,omitempty"`
	Inner             *MolecularSequenceStructureVariantInner `bson:"inner,omitempty" json:"inner,omitempty"`
}

// Structural variant outer.
type MolecularSequenceStructureVariantOuter struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Start             *int        `bson:"start,omitempty" json:"start,omitempty"`
	End               *int        `bson:"end,omitempty" json:"end,omitempty"`
}

// Structural variant inner.
type MolecularSequenceStructureVariantInner struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Start             *int        `bson:"start,omitempty" json:"start,omitempty"`
	End               *int        `bson:"end,omitempty" json:"end,omitempty"`
}

// This function returns resource reference information
func (r MolecularSequence) ResourceRef() (string, *string) {
	return "MolecularSequence", r.Id
}

// This function returns resource reference information
func (r MolecularSequence) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMolecularSequence MolecularSequence

// MarshalJSON marshals the given MolecularSequence as JSON into a byte slice
func (r MolecularSequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMolecularSequence
		ResourceType string `json:"resourceType"`
	}{
		OtherMolecularSequence: OtherMolecularSequence(r),
		ResourceType:           "MolecularSequence",
	})
}

// UnmarshalMolecularSequence unmarshals a MolecularSequence.
func UnmarshalMolecularSequence(b []byte) (MolecularSequence, error) {
	var molecularSequence MolecularSequence
	if err := json.Unmarshal(b, &molecularSequence); err != nil {
		return molecularSequence, err
	}
	return molecularSequence, nil
}
