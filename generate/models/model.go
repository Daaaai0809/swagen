package models

import (
	"github.com/Daaaai0809/swagen/generate"
	"gopkg.in/yaml.v2"
)

type ModelSchema struct {
	Title      string                  `yaml:"title"`
	Type       string                  `yaml:"type"`
	Properties generate.IPropertiesMap `yaml:"properties,omitempty"`
}

func NewModelSchema(title string, type_ string, properties generate.IPropertiesMap) *ModelSchema {
	return &ModelSchema{
		Title:      title,
		Type:       type_,
		Properties: properties,
	}
}

func (m *ModelSchema) ToYaml() (string, error) {
	yamlBytes, err := yaml.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(yamlBytes), nil
}
