# gofhir-client
FHIR models & client for Go, forked from https://github.com/samply/golang-fhir-models at 03dc486

## Develop

This repository contains two Go modules, the generated models itself and the generator. 
Both modules use `go generate` to generate the FHIR models. For `go generate` to work, 
you have to install the generator first. To do that, run `go install` in the `fhir-models-gen` 
directory. After that, you can regenerate the FHIR Models under `fhir-models` and the subset of FHIR models under `fhir-models-gen`.
