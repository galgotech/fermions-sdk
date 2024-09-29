package validate

import (
	"github.com/galgotech/fermions-sdk/graph"
	"github.com/galgotech/fermions-sdk/internal/validator"
)

// FromFile parses the given Serverless Workflow file into the Workflow type.
func FromFile(path string) error {
	root, fileBytes, err := graph.FromFile(path)
	if err != nil {
		return err
	}

	return validator.Valid(root, fileBytes)
}

// FromYAMLSource parses the given Serverless Workflow YAML source into the Workflow type.
func FromYAMLSource(source []byte) error {
	root, jsonBytes, err := graph.FromYAMLSource(source)
	if err != nil {
		return err
	}

	return validator.Valid(root, jsonBytes)
}

// FromJSONSource parses the given Serverless Workflow JSON source into the Workflow type.
func FromJSONSource(source []byte) error {
	root, jsonBytes, err := graph.FromJSONSource(source)
	if err != nil {
		return err
	}

	return validator.Valid(root, jsonBytes)
}
