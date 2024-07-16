package methods

import (
	"bufio"

	"github.com/Daaaai0809/swagen/constant"
	"github.com/Daaaai0809/swagen/input"
	"github.com/spf13/cobra"
)

type PathRoot struct {
	Get  GetPathSchema  `yaml:"get,omitempty"`
	Post PostPathSchema `yaml:"post,omitempty"`
	Put  PutPathSchema  `yaml:"put,omitempty"`
	// Patch PatchPathSchema `yaml:"patch,omitempty"`
	// Delete DeletePathSchema `yaml:"delete,omitempty"`
}

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

type Schema struct {
	Type       string            `yaml:"type"`
	Format     string            `yaml:"format,omitempty"`
	Properties map[string]Schema `yaml:"properties,omitempty"`
	Required   []string          `yaml:"required,omitempty"`
	Nullable   bool              `yaml:"nullable,omitempty"`
	// TODO: Must support Example
	// Example    interface{}               `yaml:"example,omitempty"`
	Items map[string]Schema `yaml:"items,omitempty"`
}

func (s *Schema) GetString() string {
	// TODO: Return Example
	return ""
}

func (s *Schema) ReadFormat(cmd *cobra.Command, t string) {
	if ok := input.YesNoPrompt(cmd, "Do you want to add a format?"); ok {
		switch t {
		case constant.STRING_TYPE:
			if f, err := input.SingleSelect("Select the format", constant.FormatStringList); err == nil {
				s.Format = f
			} else {
				cmd.Println(err)
			}
		case constant.NUMBER_TYPE:
			if f, err := input.SingleSelect("Select the format", constant.FormatNumberList); err == nil {
				s.Format = f
			} else {
				cmd.Println(err)
			}
		case constant.INTEGER_TYPE:
			if f, err := input.SingleSelect("Select the format", constant.FormatIntegerList); err == nil {
				s.Format = f
			} else {
				cmd.Println(err)
			}
		}
	}
}

func (s *Schema) ReadProperties(cmd *cobra.Command, scanner *bufio.Scanner) {
	properties := make(map[string]Schema)

	for {
		prop := Schema{}

		cmd.Println("Enter the property name: ")

		scanner.Scan()

		name := scanner.Text()

		if t, err := input.SingleSelect("Select the type", constant.SchemaTypeList); err == nil {
			prop.Type = t
		} else {
			cmd.Println(err)
			continue
		}

		prop.ReadFormat(cmd, prop.Type)

		switch prop.Type {
		case constant.OBJECT_TYPE:
			prop.ReadProperties(cmd, scanner)
		case constant.ARRAY_TYPE:
			prop.ReadItems(cmd, scanner)
		default:
			if ok := input.YesNoPrompt(cmd, "Is the property required?"); ok {
				s.Required = append(s.Required, name)
			}

			if ok := input.YesNoPrompt(cmd, "Is the property nullable?"); ok {
				prop.Nullable = true
			}
		}

		properties[name] = prop

		if ok := input.YesNoPrompt(cmd, "Do you want to add another property?"); !ok {
			break
		}
	}

	s.Properties = properties
}

func (s *Schema) ReadItems(cmd *cobra.Command, scanner *bufio.Scanner) {
	s.Items = make(map[string]Schema)

	for {
		item := Schema{}

		cmd.Println("Enter the item name: ")

		scanner.Scan()

		name := scanner.Text()

		if t, err := input.SingleSelect("Select the type", constant.SchemaTypeList); err == nil {
			item.Type = t
		} else {
			cmd.Println(err)
			continue
		}

		item.ReadFormat(cmd, item.Type)

		switch item.Type {
		case constant.OBJECT_TYPE:
			item.ReadProperties(cmd, scanner)
		case constant.ARRAY_TYPE:
			item.ReadItems(cmd, scanner)
		default:
			if ok := input.YesNoPrompt(cmd, "Is the item required?"); ok {
				s.Required = append(s.Required, name)
			}

			if ok := input.YesNoPrompt(cmd, "Is the item nullable?"); ok {
				item.Nullable = true
			}
		}

		s.Items[name] = item

		if ok := input.YesNoPrompt(cmd, "Do you want to add another item?"); !ok {
			break
		}
	}
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
