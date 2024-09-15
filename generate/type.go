package generate

import (
	"bufio"

	"github.com/Daaaai0809/swagen/constant"
	"github.com/spf13/cobra"
)

type IParameters interface {
	GetString() string
}

// Parameters is an alias for type Parameter or RefParameter
type Parameters []IParameters

type RefParameter struct {
	Ref string `yaml:"$ref"`
}

func (r *RefParameter) GetString() string {
	return string(r.Ref)
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
	// TODO : Must support Example
	// Example interface{} `yaml:"example,omitempty"`
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

func (r *RefSchema) ReadRef(cmd *cobra.Command, scanner *bufio.Scanner) {
	cmd.Println("Enter the Ref: ")

	scanner.Scan()

	r.Ref = scanner.Text()
}

type RequestBody struct {
	Description string             `yaml:"description"`
	Content     map[string]Content `yaml:"content"`
}

func (r *RequestBody) GetString() string {
	return ""
}

func (r *RequestBody) ReadDescription(cmd *cobra.Command, scanner *bufio.Scanner) {
	cmd.Println("Enter the request body description: ")

	scanner.Scan()

	r.Description = scanner.Text()
}

type ISchema interface {
	GetString() string
}

type IPropertiesMap interface {
	GetKeys() []string
}

type PropertiesMap map[string]ISchema

func (p PropertiesMap) GetKeys() []string {
	keys := make([]string, 0, len(p))

	for k := range p {
		keys = append(keys, k)
	}

	return keys
}

type Schema struct {
	Type       string         `yaml:"type"`
	Format     string         `yaml:"format,omitempty"`
	Properties IPropertiesMap `yaml:"properties,omitempty"`
	Required   []string       `yaml:"required,omitempty"`
	Nullable   bool           `yaml:"nullable,omitempty"`
	Items      ISchema        `yaml:"items,omitempty"`
	// TODO: Must support Example
	// Example    interface{}               `yaml:"example,omitempty"`
}

func NewSchema() *Schema {
	return &Schema{
		Properties: make(PropertiesMap),
	}
}

func (s *Schema) GetString() string {
	// TODO: Return Example
	return ""
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
