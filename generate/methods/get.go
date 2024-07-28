package methods

import (
	"github.com/Daaaai0809/swagen/generate"
	"gopkg.in/yaml.v2"
)

type GetPathSchema struct {
	OperationID string              `yaml:"operationId"`
	Summary     string              `yaml:"summary,omitempty"`
	Description string              `yaml:"description,omitempty"`
	Tags        []string            `yaml:"tags,omitempty"`
	Security    []generate.Security `yaml:"security,omitempty"`
	Parameters  generate.Parameters `yaml:"parameters,omitempty"`
	Responses   generate.Responses  `yaml:"responses"`
}

func (p *GetPathSchema) ToYaml() (string, error) {
	r := PathRoot{
		Get: *p,
	}

	yamlBytes, err := yaml.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(yamlBytes), nil
}

func NewGetPathSchema(operationID, summary, description string, tags []string, security []generate.Security, parameters generate.Parameters, responses generate.Responses) *GetPathSchema {
	return &GetPathSchema{
		OperationID: operationID,
		Summary:     summary,
		Description: description,
		Tags:        tags,
		Security:    security,
		Parameters:  parameters,
		Responses:   responses,
	}
}
