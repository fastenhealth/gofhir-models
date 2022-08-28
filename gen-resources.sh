#!/usr/bin/env bash

# http://hl7.org/fhir/directory.html
# http://hl7.org/fhir/2021Mar/downloads.html

generateDefinitions() {
  echo "generateDefinitions() $*"
  mkdir -p "fhir${1}"
  curl -L -o "definitions.zip" "${2}"
  unzip -o definitions.zip profiles-resources.json profiles-types.json valuesets.json -d "fhir${1}"
  rm definitions.zip
  pushd "fhir${1}"
  gofhir-models-gen gen-resources --package fhir${1} .
  popd
#  go generate "./fhir${1}"
}

# R4B (Current)
generateDefinitions "430" "http://hl7.org/fhir/R4/definitions.json.zip"

# R4
generateDefinitions "401" "http://hl7.org/fhir/R4/definitions.json.zip"

# FHIR STU3 Candidate + Connectathon 14 (San Antonio)
#generateDefinitions "1.8.0" "http://hl7.org/fhir/2017Jan/definitions.json.zip"

# DSTU 2 (Official version) with 1 technical errata (Permanent home)
#generateDefinitions "1.0.2" "http://hl7.org/fhir/DSTU2/definitions.json.zip"
