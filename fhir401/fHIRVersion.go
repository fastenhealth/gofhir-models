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

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// FHIRVersion is documented here http://hl7.org/fhir/ValueSet/FHIR-version
type FHIRVersion int

const (
	FHIRVersion0_01 FHIRVersion = iota
	FHIRVersion0_05
	FHIRVersion0_06
	FHIRVersion0_11
	FHIRVersion0_0_80
	FHIRVersion0_0_81
	FHIRVersion0_0_82
	FHIRVersion0_4_0
	FHIRVersion0_5_0
	FHIRVersion1_0_0
	FHIRVersion1_0_1
	FHIRVersion1_0_2
	FHIRVersion1_1_0
	FHIRVersion1_4_0
	FHIRVersion1_6_0
	FHIRVersion1_8_0
	FHIRVersion3_0_0
	FHIRVersion3_0_1
	FHIRVersion3_3_0
	FHIRVersion3_5_0
	FHIRVersion4_0_0
	FHIRVersion4_0_1
)

func (code FHIRVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *FHIRVersion) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "0.01":
		*code = FHIRVersion0_01
	case "0.05":
		*code = FHIRVersion0_05
	case "0.06":
		*code = FHIRVersion0_06
	case "0.11":
		*code = FHIRVersion0_11
	case "0.0.80":
		*code = FHIRVersion0_0_80
	case "0.0.81":
		*code = FHIRVersion0_0_81
	case "0.0.82":
		*code = FHIRVersion0_0_82
	case "0.4.0":
		*code = FHIRVersion0_4_0
	case "0.5.0":
		*code = FHIRVersion0_5_0
	case "1.0.0":
		*code = FHIRVersion1_0_0
	case "1.0.1":
		*code = FHIRVersion1_0_1
	case "1.0.2":
		*code = FHIRVersion1_0_2
	case "1.1.0":
		*code = FHIRVersion1_1_0
	case "1.4.0":
		*code = FHIRVersion1_4_0
	case "1.6.0":
		*code = FHIRVersion1_6_0
	case "1.8.0":
		*code = FHIRVersion1_8_0
	case "3.0.0":
		*code = FHIRVersion3_0_0
	case "3.0.1":
		*code = FHIRVersion3_0_1
	case "3.3.0":
		*code = FHIRVersion3_3_0
	case "3.5.0":
		*code = FHIRVersion3_5_0
	case "4.0.0":
		*code = FHIRVersion4_0_0
	case "4.0.1":
		*code = FHIRVersion4_0_1
	default:
		return fmt.Errorf("unknown FHIRVersion code `%s`", s)
	}
	return nil
}
func (code FHIRVersion) String() string {
	return code.Code()
}
func (code FHIRVersion) Code() string {
	switch code {
	case FHIRVersion0_01:
		return "0.01"
	case FHIRVersion0_05:
		return "0.05"
	case FHIRVersion0_06:
		return "0.06"
	case FHIRVersion0_11:
		return "0.11"
	case FHIRVersion0_0_80:
		return "0.0.80"
	case FHIRVersion0_0_81:
		return "0.0.81"
	case FHIRVersion0_0_82:
		return "0.0.82"
	case FHIRVersion0_4_0:
		return "0.4.0"
	case FHIRVersion0_5_0:
		return "0.5.0"
	case FHIRVersion1_0_0:
		return "1.0.0"
	case FHIRVersion1_0_1:
		return "1.0.1"
	case FHIRVersion1_0_2:
		return "1.0.2"
	case FHIRVersion1_1_0:
		return "1.1.0"
	case FHIRVersion1_4_0:
		return "1.4.0"
	case FHIRVersion1_6_0:
		return "1.6.0"
	case FHIRVersion1_8_0:
		return "1.8.0"
	case FHIRVersion3_0_0:
		return "3.0.0"
	case FHIRVersion3_0_1:
		return "3.0.1"
	case FHIRVersion3_3_0:
		return "3.3.0"
	case FHIRVersion3_5_0:
		return "3.5.0"
	case FHIRVersion4_0_0:
		return "4.0.0"
	case FHIRVersion4_0_1:
		return "4.0.1"
	}
	return "<unknown>"
}
func (code FHIRVersion) Display() string {
	switch code {
	case FHIRVersion0_01:
		return "0.01"
	case FHIRVersion0_05:
		return "0.05"
	case FHIRVersion0_06:
		return "0.06"
	case FHIRVersion0_11:
		return "0.11"
	case FHIRVersion0_0_80:
		return "0.0.80"
	case FHIRVersion0_0_81:
		return "0.0.81"
	case FHIRVersion0_0_82:
		return "0.0.82"
	case FHIRVersion0_4_0:
		return "0.4.0"
	case FHIRVersion0_5_0:
		return "0.5.0"
	case FHIRVersion1_0_0:
		return "1.0.0"
	case FHIRVersion1_0_1:
		return "1.0.1"
	case FHIRVersion1_0_2:
		return "1.0.2"
	case FHIRVersion1_1_0:
		return "1.1.0"
	case FHIRVersion1_4_0:
		return "1.4.0"
	case FHIRVersion1_6_0:
		return "1.6.0"
	case FHIRVersion1_8_0:
		return "1.8.0"
	case FHIRVersion3_0_0:
		return "3.0.0"
	case FHIRVersion3_0_1:
		return "3.0.1"
	case FHIRVersion3_3_0:
		return "3.3.0"
	case FHIRVersion3_5_0:
		return "3.5.0"
	case FHIRVersion4_0_0:
		return "4.0.0"
	case FHIRVersion4_0_1:
		return "4.0.1"
	}
	return "<unknown>"
}
func (code FHIRVersion) Definition() string {
	switch code {
	case FHIRVersion0_01:
		return "Oldest archived version of FHIR."
	case FHIRVersion0_05:
		return "1st Draft for Comment (Sept 2012 Ballot)."
	case FHIRVersion0_06:
		return "2nd Draft for Comment (January 2013 Ballot)."
	case FHIRVersion0_11:
		return "DSTU 1 Ballot version."
	case FHIRVersion0_0_80:
		return "DSTU 1 Official version."
	case FHIRVersion0_0_81:
		return "DSTU 1 Official version Technical Errata #1."
	case FHIRVersion0_0_82:
		return "DSTU 1 Official version Technical Errata #2."
	case FHIRVersion0_4_0:
		return "Draft For Comment (January 2015 Ballot)."
	case FHIRVersion0_5_0:
		return "DSTU 2 Ballot version (May 2015 Ballot)."
	case FHIRVersion1_0_0:
		return "DSTU 2 QA Preview + CQIF Ballot (Sep 2015)."
	case FHIRVersion1_0_1:
		return "DSTU 2 (Official version)."
	case FHIRVersion1_0_2:
		return "DSTU 2 (Official version) with 1 technical errata."
	case FHIRVersion1_1_0:
		return "GAO Ballot + draft changes to main FHIR standard."
	case FHIRVersion1_4_0:
		return "CQF on FHIR Ballot + Connectathon 12 (Montreal)."
	case FHIRVersion1_6_0:
		return "FHIR STU3 Ballot + Connectathon 13 (Baltimore)."
	case FHIRVersion1_8_0:
		return "FHIR STU3 Candidate + Connectathon 14 (San Antonio)."
	case FHIRVersion3_0_0:
		return "FHIR Release 3 (STU)."
	case FHIRVersion3_0_1:
		return "FHIR Release 3 (STU) with 1 technical errata."
	case FHIRVersion3_3_0:
		return "R4 Ballot #1."
	case FHIRVersion3_5_0:
		return "R4 Ballot #2."
	case FHIRVersion4_0_0:
		return "FHIR Release 4 (Normative + STU)."
	case FHIRVersion4_0_1:
		return "FHIR Release 4 Technical Correction."
	}
	return "<unknown>"
}