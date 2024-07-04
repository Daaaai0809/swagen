package methods

import (
	"gopkg.in/yaml.v2"
)

type GetPathSchema struct {
	OperationID string     `yaml:"operationId"`
	Summary     string     `yaml:"summary,omitempty"`
	Description string     `yaml:"description,omitempty"`
	Tags        []string   `yaml:"tags,omitempty"`
	Security    []Security `yaml:"security,omitempty"`
	Parameters  Parameters `yaml:"parameters,omitempty"`
	Responses   Responses  `yaml:"responses"`
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

func NewGetPathSchema(operationID, summary, description string, tags []string, security []Security, parameters Parameters, responses Responses) *GetPathSchema {
	if len(responses) == 0 {
		responses = Responses{
			"200": Response{
				Description: "successful operation",
				Content: map[string]Content{
					"application/json": {
						Schema: ContentSchema{
							&RefSchema{
								Ref: "#/components/schemas/Response",
							},
						},
					},
				},
			},
		}
	}

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
