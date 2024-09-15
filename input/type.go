package input

import (
	"bufio"

	"github.com/Daaaai0809/swagen/constant"
	"github.com/Daaaai0809/swagen/generate"
	"github.com/spf13/cobra"
)

type InputProperties map[string]*InputSchema

func (p InputProperties) GetKeys() []string {
	keys := make([]string, 0, len(p))
	for k := range p {
		keys = append(keys, k)
	}
	return keys
}

type InputSchema struct {
	generate.Schema
}

func NewInputSchema() *InputSchema {
	return &InputSchema{
		Schema: generate.Schema{
			Properties: make(InputProperties),
		},
	}
}

func (s *InputSchema) GetString() string {
	return ""
}

func (s *InputSchema) ReadType(cmd *cobra.Command) {
	if t, err := SingleSelect("Select the type", constant.SchemaTypeList); err == nil {
		s.Type = t
	} else {
		cmd.Println(err)
	}
}

func (s *InputSchema) ReadFormat(cmd *cobra.Command, t string) {
	if ok := YesNoPrompt(cmd, "Do you want to add a format?"); ok {
		switch t {
		case constant.STRING_TYPE:
			if f, err := SingleSelect("Select the format", constant.FormatStringList); err == nil {
				s.Format = f
			} else {
				cmd.Println(err)
			}
		case constant.NUMBER_TYPE:
			if f, err := SingleSelect("Select the format", constant.FormatNumberList); err == nil {
				s.Format = f
			} else {
				cmd.Println(err)
			}
		case constant.INTEGER_TYPE:
			if f, err := SingleSelect("Select the format", constant.FormatIntegerList); err == nil {
				s.Format = f
			} else {
				cmd.Println(err)
			}
		}
	}
}

func (s *InputSchema) ReadProperties(cmd *cobra.Command, scanner *bufio.Scanner, isModel bool) {
	properties := make(InputProperties)
	for {
		prop := NewInputSchema()

		cmd.Println("Enter the property name: ")

		scanner.Scan()

		name := scanner.Text()

		if t, err := SingleSelect("Select the type", constant.SchemaTypeList); err == nil {
			prop.Type = t
		} else {
			cmd.Println(err)
			continue
		}

		if prop.Type == constant.STRING_TYPE || prop.Type == constant.NUMBER_TYPE || prop.Type == constant.INTEGER_TYPE {
			prop.ReadFormat(cmd, prop.Type)
		}

		switch prop.Type {
		case constant.OBJECT_TYPE:
			prop.ReadProperties(cmd, scanner, isModel)
		case constant.ARRAY_TYPE:
			prop.ReadItems(cmd, scanner, isModel)
		default:
			if ok := YesNoPrompt(cmd, "Is the property required?"); ok && !isModel {
				s.Required = append(s.Required, name)
			}

			if ok := YesNoPrompt(cmd, "Is the property nullable?"); ok {
				prop.Nullable = true
			}
		}

		properties[name] = prop

		if ok := YesNoPrompt(cmd, "Do you want to add another property?"); !ok {
			break
		}
	}

	s.Properties = properties
}

func (s *InputSchema) ReadItems(cmd *cobra.Command, scanner *bufio.Scanner, isModel bool) {
	item := &InputSchema{}

	if t, err := SingleSelect("Select the type", constant.SchemaTypeList); err == nil {
		item.Type = t
	} else {
		cmd.Println(err)
		return
	}

	if item.Type == constant.STRING_TYPE || item.Type == constant.NUMBER_TYPE || item.Type == constant.INTEGER_TYPE {
		item.ReadFormat(cmd, item.Type)
	}

	switch item.Type {
	case constant.OBJECT_TYPE:
		item.ReadProperties(cmd, scanner, isModel)
	case constant.ARRAY_TYPE:
		item.ReadItems(cmd, scanner, isModel)
	default:
		if ok := YesNoPrompt(cmd, "Is the item nullable?"); ok {
			item.Nullable = true
		}
	}

	s.Items = item
}
