package path

import (
	"bufio"
	"os"

	"github.com/Daaaai0809/swagen/constant"
	"github.com/Daaaai0809/swagen/generate/methods"
	"github.com/Daaaai0809/swagen/input"
	"github.com/spf13/cobra"
)

type PostPathInputs struct {
	RootInputs
	RequestBody methods.RequestBody
}

func NewPostPathInputs(cmd *cobra.Command) *PostPathInputs {
	return &PostPathInputs{
		RootInputs: RootInputs{
			Cmd:         cmd,
			FileName:    "",
			OperationID: "",
			Summary:     "",
			Description: "",
			Tags:        []string{},
			Security:    []string{},
			Parameters:  methods.Parameters{},
			Responses:   methods.Responses{},
		},
		RequestBody: methods.RequestBody{},
	}
}

// The ReadAll method reads all the input required to define an endpoint.
func (p *PostPathInputs) ReadAll() {
	p.ReadFileName()
	p.ReadOperationID()

	if ok := input.YesNoPrompt(p.Cmd, "Do you want to add a summary?"); ok {
		p.ReadSummary()
	}

	if ok := input.YesNoPrompt(p.Cmd, "Do you want to add a description?"); ok {
		p.ReadDescription()
	}

	if ok := input.YesNoPrompt(p.Cmd, "Do you want to add tags?"); ok {
		p.ReadTags()
	}

	if ok := input.YesNoPrompt(p.Cmd, "Do you want to add security?"); ok {
		p.ReadSecurity()
	}

	if ok := input.YesNoPrompt(p.Cmd, "Do you want to add request body?"); ok {
		p.ReadRequestBody()
	}

	if ok := input.YesNoPrompt(p.Cmd, "Do you want to add parameters?"); ok {
		p.ReadParameters()
	}

	if ok := input.YesNoPrompt(p.Cmd, "Do you want to add responses?"); ok {
		p.ReadResponses()
	}
}

func (p *PostPathInputs) SetRequestBody(requestBody methods.RequestBody) {
	p.RequestBody = requestBody
}

func (p *PostPathInputs) GetRequestBody() methods.RequestBody {
	return p.RequestBody
}

// The ReadRequestBody method takes input from the CLI to define the request body for the endpoint.
func (p *PostPathInputs) ReadRequestBody() {
	requestBody := methods.RequestBody{}

	scanner := bufio.NewScanner(os.Stdin)

	requestBody.Content = make(map[string]methods.Content)

	p.Cmd.Println("Enter the description: ")

	scanner.Scan()

	requestBody.Description = scanner.Text()

	contentType, err := input.SingleSelect("Select the content type", constant.ContentTypeList)
	if err != nil {
		p.Cmd.Println(err)
		return
	}

	content := methods.Content{
		Schema: methods.ContentSchema{},
	}

	for {
		if ok := input.YesNoPrompt(p.Cmd, "Do you want to add a ref?"); ok {
			refSchema := methods.RefSchema{}

			p.Cmd.Println("Enter the ref: ")

			scanner.Scan()

			refSchema.Ref = scanner.Text()

			content.Schema = append(content.Schema, &refSchema)
		} else {
			schema := methods.Schema{}

			if t, err := input.SingleSelect("Select the type", constant.SchemaTypeList); err == nil {
				schema.Type = t
			} else {
				p.Cmd.Println(err)
				continue
			}

			if schema.Type == constant.STRING_TYPE || schema.Type == constant.NUMBER_TYPE || schema.Type == constant.INTEGER_TYPE {
				format := ReadFormat(p, schema.Type)

				if format != "" {
					schema.Format = format
				}
			}

			switch schema.Type {
			case constant.OBJECT_TYPE:
				ReadProperties(p, &schema, scanner)
			case constant.ARRAY_TYPE:
				ReadItems(p, &schema, scanner)
			default:
				if ok := input.YesNoPrompt(p.Cmd, "Is the schema nullable?"); ok {
					schema.Nullable = true
				}
			}

			content.Schema = append(content.Schema, &schema)
		}

		requestBody.Content[contentType] = content

		if ok := input.YesNoPrompt(p.Cmd, "Do you want to add another content type?"); !ok {
			break
		}
	}

	p.SetRequestBody(requestBody)
}

func ReadFormat(p *PostPathInputs, t string) string {
	if ok := input.YesNoPrompt(p.Cmd, "Do you want to add a format?"); ok {
		switch t {
		case constant.STRING_TYPE:
			if f, err := input.SingleSelect("Select the format", constant.FormatStringList); err == nil {
				return f
			} else {
				p.Cmd.Println(err)
			}
		case constant.NUMBER_TYPE:
			if f, err := input.SingleSelect("Select the format", constant.FormatNumberList); err == nil {
				return f
			} else {
				p.Cmd.Println(err)
			}
		case constant.INTEGER_TYPE:
			if f, err := input.SingleSelect("Select the format", constant.FormatIntegerList); err == nil {
				return f
			} else {
				p.Cmd.Println(err)
			}
		}
	}

	return ""
}

func ReadProperties(p *PostPathInputs, schema *methods.Schema, scanner *bufio.Scanner) {
	properties := make(map[string]methods.Schema)

	for {
		prop := methods.Schema{}

		p.Cmd.Println("Enter the property name: ")

		scanner.Scan()

		name := scanner.Text()

		if t, err := input.SingleSelect("Select the type", constant.SchemaTypeList); err == nil {
			prop.Type = t
		} else {
			p.Cmd.Println(err)
			continue
		}

		format := ReadFormat(p, prop.Type)

		if format != "" {
			prop.Format = format
		}

		switch prop.Type {
		case constant.OBJECT_TYPE:
			ReadProperties(p, &prop, scanner)
		case constant.ARRAY_TYPE:
			ReadItems(p, &prop, scanner)
		default:
			if ok := input.YesNoPrompt(p.Cmd, "Is the property required?"); ok {
				schema.Required = append(schema.Required, name)
			}

			if ok := input.YesNoPrompt(p.Cmd, "Is the property nullable?"); ok {
				prop.Nullable = true
			}
		}

		properties[name] = prop

		if ok := input.YesNoPrompt(p.Cmd, "Do you want to add another property?"); !ok {
			break
		}
	}

	schema.Properties = properties
}

func ReadItems(p *PostPathInputs, schema *methods.Schema, scanner *bufio.Scanner) {
	if ok := input.YesNoPrompt(p.Cmd, "Do you want to add items?"); !ok {
		return
	}

	schema.Items = make(map[string]methods.Schema)

	for {
		item := methods.Schema{}

		p.Cmd.Println("Enter the item name: ")

		scanner.Scan()

		name := scanner.Text()

		if t, err := input.SingleSelect("Select the type", constant.SchemaTypeList); err == nil {
			item.Type = t
		} else {
			p.Cmd.Println(err)
			continue
		}

		format := ReadFormat(p, item.Type)

		if format != "" {
			item.Format = format
		}

		switch item.Type {
		case constant.OBJECT_TYPE:
			ReadProperties(p, &item, scanner)
		case constant.ARRAY_TYPE:
			ReadItems(p, &item, scanner)
		default:
			if ok := input.YesNoPrompt(p.Cmd, "Is the item required?"); ok {
				schema.Required = append(schema.Required, name)
			}

			if ok := input.YesNoPrompt(p.Cmd, "Is the item nullable?"); ok {
				item.Nullable = true
			}
		}

		schema.Items[name] = item

		if ok := input.YesNoPrompt(p.Cmd, "Do you want to add another item?"); !ok {
			break
		}
	}
}
