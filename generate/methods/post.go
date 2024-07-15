package methods

import (
	"gopkg.in/yaml.v2"
)

type PostPathSchema struct {
	OperationID string      `yaml:"operationId"`
	Summary     string      `yaml:"summary,omitempty"`
	Description string      `yaml:"description,omitempty"`
	Tags        []string    `yaml:"tags,omitempty"`
	Security    []Security  `yaml:"security,omitempty"`
	RequestBody RequestBody `yaml:"requestBody,omitempty"`
	Parameters  Parameters  `yaml:"parameters,omitempty"`
	Responses   Responses   `yaml:"responses"`
}

func NewPostPathSchema(operationID, summary, description string, tags []string, security []Security, requestBody RequestBody, parameters Parameters, responses Responses) *PostPathSchema {
	return &PostPathSchema{
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

func (p *PostPathSchema) ToYaml() (string, error) {
	r := PathRoot{
		Post: *p,
	}

	yamlBytes, err := yaml.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(yamlBytes), nil
}

type RequestBody struct {
	Description string             `yaml:"description"`
	Content     map[string]Content `yaml:"content"`
}

func (r *RequestBody) GetString() string {
	return ""
}
