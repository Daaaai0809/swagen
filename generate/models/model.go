package models

import (
	"github.com/Daaaai0809/swagen/generate"
	"gopkg.in/yaml.v3"
)

type ModelSchema struct {
	Name       string                  `yaml:"-"`
	Title      string                  `yaml:"title"`
	Type       string                  `yaml:"type"`
	Properties generate.IPropertiesMap `yaml:"properties,omitempty"`
}

type YamlModelSchema = map[string]ModelSchema

func NewModelSchema(name, title, type_ string, properties generate.IPropertiesMap) *ModelSchema {
	return &ModelSchema{
		Name:       name,
		Title:      title,
		Type:       type_,
		Properties: properties,
	}
}

func (m *ModelSchema) ToYaml() (string, error) {
	yamlBytes, err := yaml.Marshal(YamlModelSchema{
		m.Name: *m,
	})
	if err != nil {
		return "", err
	}

	return string(yamlBytes), nil
}
