// Copyright (c) 2011-2017, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CareTeam struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               string                         `bson:"status,omitempty" json:"status,omitempty"`
	Category             []CodeableConcept              `bson:"category,omitempty" json:"category,omitempty"`
	Name                 string                         `bson:"name,omitempty" json:"name,omitempty"`
	Subject              *Reference                     `bson:"subject,omitempty" json:"subject,omitempty"`
	Context              *Reference                     `bson:"context,omitempty" json:"context,omitempty"`
	Period               *Period                        `bson:"period,omitempty" json:"period,omitempty"`
	Participant          []CareTeamParticipantComponent `bson:"participant,omitempty" json:"participant,omitempty"`
	ReasonCode           []CodeableConcept              `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference      []Reference                    `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	ManagingOrganization []Reference                    `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Note                 []Annotation                   `bson:"note,omitempty" json:"note,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *CareTeam) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "CareTeam"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to CareTeam), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *CareTeam) GetBSON() (interface{}, error) {
	x.ResourceType = "CareTeam"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "careTeam" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type careTeam CareTeam

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *CareTeam) UnmarshalJSON(data []byte) (err error) {
	x2 := careTeam{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
			}
		}
		*x = CareTeam(x2)
		return x.checkResourceType()
	}
	return
}

func (x *CareTeam) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "CareTeam"
	} else if x.ResourceType != "CareTeam" {
		return errors.New(fmt.Sprintf("Expected resourceType to be CareTeam, instead received %s", x.ResourceType))
	}
	return nil
}

type CareTeamParticipantComponent struct {
	BackboneElement `bson:",inline"`
	Role            *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Member          *Reference       `bson:"member,omitempty" json:"member,omitempty"`
	OnBehalfOf      *Reference       `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
	Period          *Period          `bson:"period,omitempty" json:"period,omitempty"`
}
