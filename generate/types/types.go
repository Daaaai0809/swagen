package types

import (
	"github.com/Daaaai0809/swagen/constant"
)

type IParameters interface {
	GetString() string
}

// Parameters is an alias for type Parameter or RefParameter
type Parameters []IParameters

type RefParameter string

func (r *RefParameter) GetString() string {
	return string(*r)
}

type Parameter struct {
	Name        string          `yaml:"name"`
	In          string          `yaml:"in"`
	Description string          `yaml:"description,omitempty"`
	Required    bool            `yaml:"required"`
	Schema      ParameterSchema `yaml:"schema"`
}

func (p *Parameter) GetString() string {
	return string(p.Name)
}

type ParameterSchema map[string]string

type Responses map[string]Response

type Response struct {
	Description string             `yaml:"description,omitempty"`
	Content     map[string]Content `yaml:"content"`
}

type Content struct {
	Schema ContentSchema `yaml:"schema"`
}

type IContentSchema interface {
	GetString() string
}

type ContentSchema []IContentSchema

type RefSchema struct {
	Ref string `yaml:"$ref"`
}

func (r *RefSchema) GetString() string {
	return r.Ref
}

type Schema struct {
	Type       string                    `yaml:"type"`
	Format     string                    `yaml:"format,omitempty"`
	Properties map[string]SchemaProperty `yaml:"properties,omitempty"`
	Required   []string                  `yaml:"required,omitempty"`
	Nullable   bool                      `yaml:"nullable,omitempty"`
	Items      map[string]SchemaProperty `yaml:"items,omitempty"`
}

func (s *Schema) GetString() string {
	return s.Type
}

type SchemaProperty struct {
	Type       string                    `yaml:"type"`
	Format     string                    `yaml:"format,omitempty"`
	Properties map[string]SchemaProperty `yaml:"properties,omitempty"`
	Required   []string                  `yaml:"required,omitempty"`
	Nullable   bool                      `yaml:"nullable,omitempty"`
	Items      map[string]SchemaProperty `yaml:"items,omitempty"`
}

type Security = map[string][]interface{}

func GetSecurity(sec []string) []map[string][]interface{} {
	var security []map[string][]interface{}

	for _, s := range sec {
		sec := constant.GetCamelCaseSecurityType(s)

		security = append(security, map[string][]interface{}{
			sec: {},
		})
	}

	return security
}
