package path

import (
	"bufio"
	"os"

	"github.com/Daaaai0809/swagen/constant"
	"github.com/Daaaai0809/swagen/generate/types"
	"github.com/Daaaai0809/swagen/iuput"
	"github.com/spf13/cobra"
)

type GetPathInputs struct {
	cmd         *cobra.Command
	fileName    string
	operationID string
	summary     string
	description string
	tags        []string
	security    []string
	parameters  types.Parameters
	responses   types.Responses
}

func NewGetPathInputs(cmd *cobra.Command) *GetPathInputs {
	return &GetPathInputs{
		cmd:         cmd,
		fileName:    "",
		operationID: "",
		summary:     "",
		description: "",
		tags:        []string{},
		security:    []string{},
		parameters:  types.Parameters{},
		responses:   types.Responses{},
	}
}

// The ReadAll method reads all the input required to define an endpoint.
func (p *GetPathInputs) ReadAll() {
	p.ReadFileName()
	p.ReadOperationID()

	if ok := input.YesNoPrompt(p.cmd, "Do you want to add a summary?"); ok {
		p.ReadSummary()
	}

	if ok := input.YesNoPrompt(p.cmd, "Do you want to add a description?"); ok {
		p.ReadDescription()
	}

	if ok := input.YesNoPrompt(p.cmd, "Do you want to add tags?"); ok {
		p.ReadTags()
	}

	if ok := input.YesNoPrompt(p.cmd, "Do you want to add security?"); ok {
		p.ReadSecurity()
	}

	if ok := input.YesNoPrompt(p.cmd, "Do you want to add parameters?"); ok {
		p.ReadParameters()
	}

	if ok := input.YesNoPrompt(p.cmd, "Do you want to add responses?"); ok {
		p.ReadResponses()
	}
}

func (p *GetPathInputs) SetFileName(fileName string) {
	p.fileName = fileName
}

func (p *GetPathInputs) GetFileName() string {
	return p.fileName
}

// The ReadFileName method takes input from the CLI to define the file name for the endpoint.
func (p *GetPathInputs) ReadFileName() {
	scanner := bufio.NewScanner(os.Stdin)

	p.cmd.Println("Enter the file name: ")

	scanner.Scan()

	p.SetFileName(scanner.Text())
}

func (p *GetPathInputs) SetOperationID(operationID string) {
	p.operationID = operationID
}

func (p *GetPathInputs) GetOperationID() string {
	return p.operationID
}

// The ReadOperationID method takes input from the CLI to define the operation ID for the endpoint.
func (p *GetPathInputs) ReadOperationID() {
	scanner := bufio.NewScanner(os.Stdin)

	p.cmd.Println("Enter the operation ID: ")

	scanner.Scan()

	p.SetOperationID(scanner.Text())
}

func (p *GetPathInputs) SetSummary(summary string) {
	p.summary = summary
}

func (p *GetPathInputs) GetSummary() string {
	return p.summary
}

// The ReadSummary method takes input from the CLI to define a summary for the endpoint.
func (p *GetPathInputs) ReadSummary() {
	scanner := bufio.NewScanner(os.Stdin)

	p.cmd.Println("Enter the summary: ")

	scanner.Scan()

	p.SetSummary(scanner.Text())
}

func (p *GetPathInputs) SetDescription(description string) {
	p.description = description
}

func (p *GetPathInputs) GetDescription() string {
	return p.description
}

// The ReadDescription method takes input from the CLI to define the description of the endpoint.
func (p *GetPathInputs) ReadDescription() {
	scanner := bufio.NewScanner(os.Stdin)

	p.cmd.Println("Enter the description: ")

	scanner.Scan()

	p.SetDescription(scanner.Text())
}

func (p *GetPathInputs) SetTags(tags []string) {
	p.tags = tags
}

func (p *GetPathInputs) GetTags() []string {
	return p.tags
}

// The ReadTags method takes input from the CLI to define tags for the endpoint.
// It supports multiple tags.
func (p *GetPathInputs) ReadTags() {
	var tags = make([]string, 0)

	scaanner := bufio.NewScanner(os.Stdin)

	for {
		p.cmd.Println("Enter a tag name: ")

		scaanner.Scan()

		tag := scaanner.Text()

		tags = append(tags, tag)

		if ok := input.YesNoPrompt(p.cmd, "Do you want to add another tag?"); !ok {
			break
		}
	}

	p.SetTags(tags)
}

func (p *GetPathInputs) SetSecurity(security []string) {
	p.security = security
}

func (p *GetPathInputs) GetSecurity() []string {
	return p.security
}

// The ReadSecurity method takes input from the CLI to define the security types required for the endpoint.
// It supports multiple security types.
func (p *GetPathInputs) ReadSecurity() {
	var securities = make([]string, 0)

	for {
		security, err := input.SingleSelect("Select a security name", constant.SecurityTypes)
		if err != nil {
			p.cmd.Println(err)
			continue
		}

		securities = append(securities, security)

		if ok := input.YesNoPrompt(p.cmd, "Do you want to add another security?"); !ok {
			break
		}
	}

	p.SetSecurity(securities)
}

func (p *GetPathInputs) SetParameters(parameters types.Parameters) {
	p.parameters = parameters
}

func (p *GetPathInputs) GetParameters() types.Parameters {
	return p.parameters
}

// The ReadParameter method takes input from the CLI to define URL parameters.
// It supports parameter definitions on the CLI as well as definitions in other files using Ref.
func (p *GetPathInputs) ReadParameters() {
	var parameters = make(types.Parameters, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		if ok := input.YesNoPrompt(p.cmd, "Ref Parameter?"); ok {
			var parameter types.RefParameter

			p.cmd.Println("Enter the Ref: ")

			scanner.Scan()

			parameter = types.RefParameter(scanner.Text())

			parameters = append(parameters, &parameter)

			if ok := input.YesNoPrompt(p.cmd, "Do you want to add another parameter?"); !ok {
				break
			}

			continue
		}

		var parameter = types.Parameter{
			Schema: make(map[string]string),
		}

		p.cmd.Println("Enter the parameter name: ")

		scanner.Scan()
		parameter.Name = scanner.Text()

		in, err := input.SingleSelect("Select the parameter location", constant.ParamLocationsList)
		if err != nil {
			p.cmd.Println(err)
			continue
		}

		parameter.In = in

		p.cmd.Println("Enter the parameter description: ")

		scanner.Scan()
		parameter.Description = scanner.Text()

		paramType, err := input.SingleSelect("Select the parameter type", constant.SchemaTypeList)
		if err != nil {
			p.cmd.Println(err)
			continue
		}

		parameter.Schema["type"] = paramType

		switch paramType {
		case constant.STRING_TYPE, constant.NUMBER_TYPE, constant.INTEGER_TYPE:
			{
				if ok := input.YesNoPrompt(p.cmd, "Do you want to add a format?"); !ok {
					break
				}

				format, err := input.SingleSelect("Select the parameter format", constant.TypeFormatMap[paramType])
				if err != nil {
					p.cmd.Println(err)
					continue
				}

				parameter.Schema["format"] = format
			}
		}

		parameters = append(parameters, &parameter)

		if ok := input.YesNoPrompt(p.cmd, "Do you want to add another parameter?"); !ok {
			break
		}
	}

	p.SetParameters(parameters)
}

func (p *GetPathInputs) SetResponses(responses types.Responses) {
	p.responses = responses
}

func (p *GetPathInputs) GetResponses() types.Responses {
	return p.responses
}

// The ReadResponses method takes input to define the response types returned by the endpoint for each response code.
// Currently, it only supports the Ref schema in Content Parameter, but we plan to support response type definitions on the CLI in the future.
func (p *GetPathInputs) ReadResponses() {
	var responses = make(types.Responses)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		var response types.Response

		statusCode, err := input.SingleSelect("Select the status code", constant.StatusCodesList)
		if err != nil {
			p.cmd.Println(err)
			continue
		}

		p.cmd.Println("Enter the response description: ")

		scanner.Scan()
		response.Description = scanner.Text()

		contentType, err := input.SingleSelect("Select the content type", constant.ContentTypeList)
		if err != nil {
			p.cmd.Println(err)
			continue
		}

		if ok := input.YesNoPrompt(p.cmd, "Do you need Ref ?"); ok {
			p.cmd.Println("Enter the Ref: ")

			scanner.Scan()

			response.Content = map[string]types.Content{
				contentType: {
					Schema: types.ContentSchema{
						&types.RefSchema{
							Ref: scanner.Text(),
						},
					},
				},
			}
		} else {
			// TODO: Implement Defination of Schema with types.Schema
			response.Content = map[string]types.Content{
				contentType: {
					Schema: types.ContentSchema{},
				},
			}
		}

		responses[statusCode] = response

		if ok := input.YesNoPrompt(p.cmd, "Do you want to add another response?"); !ok {
			break
		}
	}

	p.SetResponses(responses)
}
