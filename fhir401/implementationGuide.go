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

// ImplementationGuide is documented here http://hl7.org/fhir/StructureDefinition/ImplementationGuide
// A set of rules of how a particular interoperability or standards problem is solved - typically through the use of FHIR resources. This resource is used to gather all the parts of an implementation guide into a logical whole and to publish a computable definition of all the parts.
type ImplementationGuide struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                        `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                     `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                         `bson:"url" json:"url"`
	Version           *string                        `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                         `bson:"name" json:"name"`
	Title             *string                        `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus              `bson:"status" json:"status"`
	Experimental      *bool                          `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                        `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                        `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail                `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                        `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext                 `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept              `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Copyright         *string                        `bson:"copyright,omitempty" json:"copyright,omitempty"`
	PackageId         string                         `bson:"packageId" json:"packageId"`
	License           *SPDXLicense                   `bson:"license,omitempty" json:"license,omitempty"`
	FhirVersion       []FHIRVersion                  `bson:"fhirVersion" json:"fhirVersion"`
	DependsOn         []ImplementationGuideDependsOn `bson:"dependsOn,omitempty" json:"dependsOn,omitempty"`
	Global            []ImplementationGuideGlobal    `bson:"global,omitempty" json:"global,omitempty"`
	Definition        *ImplementationGuideDefinition `bson:"definition,omitempty" json:"definition,omitempty"`
	Manifest          *ImplementationGuideManifest   `bson:"manifest,omitempty" json:"manifest,omitempty"`
}

// Another implementation guide that this implementation depends on. Typically, an implementation guide uses value sets, profiles etc.defined in other implementation guides.
type ImplementationGuideDependsOn struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Uri               string      `bson:"uri" json:"uri"`
	PackageId         *string     `bson:"packageId,omitempty" json:"packageId,omitempty"`
	Version           *string     `bson:"version,omitempty" json:"version,omitempty"`
}

// A set of profiles that all resources covered by this implementation guide must conform to.
// See [Default Profiles](implementationguide.html#default) for a discussion of which resources are 'covered' by an implementation guide.
type ImplementationGuideGlobal struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              ResourceType `bson:"type" json:"type"`
	Profile           string       `bson:"profile" json:"profile"`
}

// The information needed by an IG publisher tool to publish the whole implementation guide.
// Principally, this consists of information abuot source resource and file locations, and build parameters and templates.
type ImplementationGuideDefinition struct {
	Id                *string                                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Grouping          []ImplementationGuideDefinitionGrouping  `bson:"grouping,omitempty" json:"grouping,omitempty"`
	Resource          []ImplementationGuideDefinitionResource  `bson:"resource" json:"resource"`
	Page              *ImplementationGuideDefinitionPage       `bson:"page,omitempty" json:"page,omitempty"`
	Parameter         []ImplementationGuideDefinitionParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	Template          []ImplementationGuideDefinitionTemplate  `bson:"template,omitempty" json:"template,omitempty"`
}

// A logical group of resources. Logical groups can be used when building pages.
// Groupings are arbitrary sub-divisions of content. Typically, they are used to help build Table of Contents automatically.
type ImplementationGuideDefinitionGrouping struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name" json:"name"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
}

// A resource that is part of the implementation guide. Conformance resources (value set, structure definition, capability statements etc.) are obvious candidates for inclusion, but any kind of resource can be included as an example resource.
type ImplementationGuideDefinitionResource struct {
	Id                *string       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Reference         Reference     `bson:"reference" json:"reference"`
	FhirVersion       []FHIRVersion `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	Name              *string       `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string       `bson:"description,omitempty" json:"description,omitempty"`
	GroupingId        *string       `bson:"groupingId,omitempty" json:"groupingId,omitempty"`
}

// A page / section in the implementation guide. The root page is the implementation guide home page.
// Pages automatically become sections if they have sub-pages. By convention, the home page is called index.html.
type ImplementationGuideDefinitionPage struct {
	Id                *string                             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Title             string                              `bson:"title" json:"title"`
	Generation        GuidePageGeneration                 `bson:"generation" json:"generation"`
	Page              []ImplementationGuideDefinitionPage `bson:"page,omitempty" json:"page,omitempty"`
}

// Defines how IG is built by tools.
type ImplementationGuideDefinitionParameter struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              GuideParameterCode `bson:"code" json:"code"`
	Value             string             `bson:"value" json:"value"`
}

// A template for building resources.
type ImplementationGuideDefinitionTemplate struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string      `bson:"code" json:"code"`
	Source            string      `bson:"source" json:"source"`
	Scope             *string     `bson:"scope,omitempty" json:"scope,omitempty"`
}

// Information about an assembled implementation guide, created by the publication tooling.
type ImplementationGuideManifest struct {
	Id                *string                               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Rendering         *string                               `bson:"rendering,omitempty" json:"rendering,omitempty"`
	Resource          []ImplementationGuideManifestResource `bson:"resource" json:"resource"`
	Page              []ImplementationGuideManifestPage     `bson:"page,omitempty" json:"page,omitempty"`
	Image             []string                              `bson:"image,omitempty" json:"image,omitempty"`
	Other             []string                              `bson:"other,omitempty" json:"other,omitempty"`
}

// A resource that is part of the implementation guide. Conformance resources (value set, structure definition, capability statements etc.) are obvious candidates for inclusion, but any kind of resource can be included as an example resource.
type ImplementationGuideManifestResource struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Reference         Reference   `bson:"reference" json:"reference"`
	RelativePath      *string     `bson:"relativePath,omitempty" json:"relativePath,omitempty"`
}

// Information about a page within the IG.
type ImplementationGuideManifestPage struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name" json:"name"`
	Title             *string     `bson:"title,omitempty" json:"title,omitempty"`
	Anchor            []string    `bson:"anchor,omitempty" json:"anchor,omitempty"`
}
type OtherImplementationGuide ImplementationGuide

// MarshalJSON marshals the given ImplementationGuide as JSON into a byte slice
func (r ImplementationGuide) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImplementationGuide
		ResourceType string `json:"resourceType"`
	}{
		OtherImplementationGuide: OtherImplementationGuide(r),
		ResourceType:             "ImplementationGuide",
	})
}

// UnmarshalImplementationGuide unmarshals a ImplementationGuide.
func UnmarshalImplementationGuide(b []byte) (ImplementationGuide, error) {
	var implementationGuide ImplementationGuide
	if err := json.Unmarshal(b, &implementationGuide); err != nil {
		return implementationGuide, err
	}
	return implementationGuide, nil
}
