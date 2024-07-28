package messages

import (
	"github.com/Daaaai0809/swagen/generate"
	"gopkg.in/yaml.v2"
)

type MessageProperties = map[string]generate.Schema

type Message struct {
	MessageName string
	Type        string
	Format      string
	Nullable    bool
	Properties  MessageProperties
	Items       *generate.Schema
	Required    []string
}

type ToYamlMessage = map[string]interface{}

type ToYamlMessageProperties struct {
	Type       string            `yaml:"type"`
	Format     string            `yaml:"format,omitempty"`
	Nullable   bool              `yaml:"nullable,omitempty"`
	Properties MessageProperties `yaml:"properties"`
	Items      *generate.Schema  `yaml:"items,omitempty"`
	Required   []string          `yaml:"required,omitempty"`
}

func NewMessageProperties() MessageProperties {
	return make(MessageProperties)
}

func NewMessage(messageName string, type_ string, format string, nullable bool, properties MessageProperties, items *generate.Schema, required []string) *Message {
	return &Message{
		MessageName: messageName,
		Type:        type_,
		Format:      format,
		Nullable:    nullable,
		Properties:  properties,
		Items:       items,
		Required:    required,
	}
}

func (m *Message) ToYaml() (string, error) {
	yamlMsg, err := yaml.Marshal(ToYamlMessage{
		m.MessageName: ToYamlMessageProperties{
			Type:       m.Type,
			Format:     m.Format,
			Nullable:   m.Nullable,
			Properties: m.Properties,
			Items:      m.Items,
			Required:   m.Required,
		},
	})

	if err != nil {
		return "", err
	}

	return string(yamlMsg), nil
}
