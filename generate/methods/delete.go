package methods

import (
	"gopkg.in/yaml.v2"
)

type DeletePathSchema struct {
	OperationID string      `yaml:"operationId"`
	Summary     string      `yaml:"summary,omitempty"`
	Description string      `yaml:"description,omitempty"`
	Tags        []string    `yaml:"tags,omitempty"`
	Security    []Security  `yaml:"security,omitempty"`
	RequestBody RequestBody `yaml:"requestBody,omitempty"`
	Parameters  Parameters  `yaml:"parameters,omitempty"`
	Responses   Responses   `yaml:"responses"`
}

func NewDeletePathSchema(operationID, summary, description string, tags []string, security []Security, requestBody RequestBody, parameters Parameters, responses Responses) *DeletePathSchema {
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
