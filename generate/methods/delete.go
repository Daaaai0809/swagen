package methods

import (
	"github.com/Daaaai0809/swagen/generate"
	"gopkg.in/yaml.v2"
)

type DeletePathSchema struct {
	OperationID string               `yaml:"operationId"`
	Summary     string               `yaml:"summary,omitempty"`
	Description string               `yaml:"description,omitempty"`
	Tags        []string             `yaml:"tags,omitempty"`
	Security    []generate.Security  `yaml:"security,omitempty"`
	RequestBody generate.RequestBody `yaml:"requestBody,omitempty"`
	Parameters  generate.Parameters  `yaml:"parameters,omitempty"`
	Responses   generate.Responses   `yaml:"responses"`
}

func NewDeletePathSchema(operationID, summary, description string, tags []string, security []generate.Security, requestBody generate.RequestBody, parameters generate.Parameters, responses generate.Responses) *DeletePathSchema {
	return &DeletePathSchema{
		OperationID: operationID,
		Summary:     summary,
		Description: description,
		Tags:        tags,
		Security:    security,
		RequestBody: requestBody,
		Parameters:  parameters,
		Responses:   responses,
	}
}

func (d *DeletePathSchema) ToYaml() (string, error) {
	r := PathRoot{
		Delete: *d,
	}

	yamlBytes, err := yaml.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(yamlBytes), nil
}
