package path

import (
	"bufio"
	"os"

	"github.com/Daaaai0809/swagen/constant"
	"github.com/Daaaai0809/swagen/generate/methods"
	"github.com/Daaaai0809/swagen/input"
	"github.com/spf13/cobra"
)

type DeletePathInputs struct {
	RootInputs
	RequestBody methods.RequestBody
}

func NewDeletePathInputs(cmd *cobra.Command) *DeletePathInputs {
	return &DeletePathInputs{
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
func (p *DeletePathInputs) ReadAll() {
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

func (p *DeletePathInputs) SetRequestBody(requestBody methods.RequestBody) {
	p.RequestBody = requestBody
}

func (p *DeletePathInputs) GetRequestBody() methods.RequestBody {
	return p.RequestBody
}

// The ReadRequestBody method takes input from the CLI to define the request body for the endpoint.
func (p *DeletePathInputs) ReadRequestBody() {
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
				schema.ReadFormat(p.Cmd, schema.Type)
			}

			switch schema.Type {
			case constant.OBJECT_TYPE:
				schema.ReadProperties(p.Cmd, scanner)
			case constant.ARRAY_TYPE:
				schema.ReadItems(p.Cmd, scanner)
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
