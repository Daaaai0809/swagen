package methods

import (
	"gopkg.in/yaml.v2"

	"github.com/Daaaai0809/swagen/generate/types"
)

type GetPathSchema struct {
	OperationID string           `yaml:"operationId"`
	Summary     string           `yaml:"summary,omitempty"`
	Description string           `yaml:"description,omitempty"`
	Tags        []string         `yaml:"tags,omitempty"`
	Security    []types.Security `yaml:"security,omitempty"`
	Parameters  types.Parameters `yaml:"parameters,omitempty"`
	Responses   types.Responses  `yaml:"responses"`
}

type GetRoot struct {
	Get GetPathSchema `yaml:"get"`
}

func (p *GetPathSchema) ToYaml() (string, error) {
	r := GetRoot{
		Get: *p,
	}

	yamlBytes, err := yaml.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(yamlBytes), nil
}

func NewGetPathSchema(operationID, summary, description string, tags []string, security []types.Security, parameters types.Parameters, responses types.Responses) *GetPathSchema {
	if len(responses) == 0 {
		responses = types.Responses{
			"200": types.Response{
				Description: "successful operation",
				Content: map[string]types.Content{
					"application/json": {
						Schema: types.ContentSchema{
							&types.RefSchema{
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
