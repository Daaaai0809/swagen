package path

import (
	"bufio"
	"os"

	"github.com/Daaaai0809/swagen/constant"
	"github.com/Daaaai0809/swagen/generate/methods"
	input "github.com/Daaaai0809/swagen/input"
	"github.com/spf13/cobra"
)

type IRootInputs interface {
	ReadAll()
	SetFileName(string)
	GetFileName() string
	ReadFileName()
	SetOperationID(string)
	GetOperationID() string
	ReadOperationID()
	SetSummary(string)
	GetSummary() string
	ReadSummary()
	SetDescription(string)
	GetDescription() string
	ReadDescription()
	SetTags([]string)
	GetTags() []string
	ReadTags()
	SetSecurity([]string)
	GetSecurity() []string
	ReadSecurity()
	SetParameters(methods.Parameters)
	GetParameters() methods.Parameters
	ReadParameters()
	SetResponses(methods.Responses)
	GetResponses() methods.Responses
}

type RootInputs struct {
	Cmd         *cobra.Command
	FileName    string
	OperationID string
	Summary     string
	Description string
	Tags        []string
	Security    []string
	Parameters  methods.Parameters
	Responses   methods.Responses
}

func (p *RootInputs) SetFileName(fileName string) {
	p.FileName = fileName
}

func (p *RootInputs) GetFileName() string {
	return p.FileName
}

// The ReadFileName method takes input from the CLI to define the file name for the endpoint.
func (p *RootInputs) ReadFileName() {
	scanner := bufio.NewScanner(os.Stdin)

	p.Cmd.Println("Enter the file name: ")

	scanner.Scan()

	p.SetFileName(scanner.Text())
}

func (p *RootInputs) SetOperationID(operationID string) {
	p.OperationID = operationID
}

func (p *RootInputs) GetOperationID() string {
	return p.OperationID
}

// The ReadOperationID method takes input from the CLI to define the operation ID for the endpoint.
func (p *RootInputs) ReadOperationID() {
	scanner := bufio.NewScanner(os.Stdin)

	p.Cmd.Println("Enter the operation ID: ")

	scanner.Scan()

	p.SetOperationID(scanner.Text())
}

func (p *RootInputs) SetSummary(summary string) {
	p.Summary = summary
}

func (p *RootInputs) GetSummary() string {
	return p.Summary
}

// The ReadSummary method takes input from the CLI to define a summary for the endpoint.
func (p *RootInputs) ReadSummary() {
	scanner := bufio.NewScanner(os.Stdin)

	p.Cmd.Println("Enter the summary: ")

	scanner.Scan()

	p.SetSummary(scanner.Text())
}

func (p *RootInputs) SetDescription(description string) {
	p.Description = description
}

func (p *RootInputs) GetDescription() string {
	return p.Description
}

// The ReadDescription method takes input from the CLI to define the description of the endpoint.
func (p *RootInputs) ReadDescription() {
	scanner := bufio.NewScanner(os.Stdin)

	p.Cmd.Println("Enter the description: ")

	scanner.Scan()

	p.SetDescription(scanner.Text())
}

func (p *RootInputs) SetTags(tags []string) {
	p.Tags = tags
}

func (p *RootInputs) GetTags() []string {
	return p.Tags
}

// The ReadTags method takes input from the CLI to define tags for the endpoint.
// It supports multiple tags.
func (p *RootInputs) ReadTags() {
	var tags = make([]string, 0)

	scaanner := bufio.NewScanner(os.Stdin)

	for {
		p.Cmd.Println("Enter a tag name: ")

		scaanner.Scan()

		tag := scaanner.Text()

		tags = append(tags, tag)

		if ok := input.YesNoPrompt(p.Cmd, "Do you want to add another tag?"); !ok {
			break
		}
	}

	p.SetTags(tags)
}

func (p *RootInputs) SetSecurity(security []string) {
	p.Security = security
}

func (p *RootInputs) GetSecurity() []string {
	return p.Security
}

// The ReadSecurity method takes input from the CLI to define the security methods required for the endpoint.
// It supports multiple security methods.
func (p *RootInputs) ReadSecurity() {
	var securities = make([]string, 0)

	for {
		security, err := input.SingleSelect("Select a security name", constant.SecurityTypes)
		if err != nil {
			p.Cmd.Println(err)
			continue
		}

		securities = append(securities, security)

		if ok := input.YesNoPrompt(p.Cmd, "Do you want to add another security?"); !ok {
			break
		}
	}

	p.SetSecurity(securities)
}

func (p *RootInputs) SetParameters(parameters methods.Parameters) {
	p.Parameters = parameters
}

func (p *RootInputs) GetParameters() methods.Parameters {
	return p.Parameters
}

// The ReadParameter method takes input from the CLI to define URL parameters.
// It supports parameter definitions on the CLI as well as definitions in other files using Ref.
func (p *RootInputs) ReadParameters() {
	var parameters = make(methods.Parameters, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		if ok := input.YesNoPrompt(p.Cmd, "Ref Parameter?"); ok {
			var parameter methods.RefParameter

			p.Cmd.Println("Enter the Ref: ")

			scanner.Scan()

			parameter = methods.RefParameter(scanner.Text())

			parameters = append(parameters, &parameter)

			if ok := input.YesNoPrompt(p.Cmd, "Do you want to add another parameter?"); !ok {
				break
			}

			continue
		}

		var parameter = methods.Parameter{
			Schema: make(map[string]string),
		}

		p.Cmd.Println("Enter the parameter name: ")

		scanner.Scan()
		parameter.Name = scanner.Text()

		in, err := input.SingleSelect("Select the parameter location", constant.ParamLocationsList)
		if err != nil {
			p.Cmd.Println(err)
			continue
		}

		parameter.In = in

		p.Cmd.Println("Enter the parameter description: ")

		scanner.Scan()
		parameter.Description = scanner.Text()

		paramType, err := input.SingleSelect("Select the parameter type", constant.SchemaTypeList)
		if err != nil {
			p.Cmd.Println(err)
			continue
		}

		parameter.Schema["type"] = paramType

		switch paramType {
		case constant.STRING_TYPE, constant.NUMBER_TYPE, constant.INTEGER_TYPE:
			{
				if ok := input.YesNoPrompt(p.Cmd, "Do you want to add a format?"); !ok {
					break
				}

				format, err := input.SingleSelect("Select the parameter format", constant.TypeFormatMap[paramType])
				if err != nil {
					p.Cmd.Println(err)
					continue
				}

				parameter.Schema["format"] = format
			}
		}

		parameters = append(parameters, &parameter)

		if ok := input.YesNoPrompt(p.Cmd, "Do you want to add another parameter?"); !ok {
			break
		}
	}

	p.SetParameters(parameters)
}

func (p *RootInputs) SetResponses(responses methods.Responses) {
	p.Responses = responses
}

func (p *RootInputs) GetResponses() methods.Responses {
	return p.Responses
}

// The ReadResponses method takes input to define the response methods returned by the endpoint for each response code.
// Currently, it only supports the Ref schema in Content Parameter, but we plan to support response type definitions on the CLI in the future.
func (p *RootInputs) ReadResponses() {
	var responses = make(methods.Responses)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		var response methods.Response

		statusCode, err := input.SingleSelect("Select the status code", constant.StatusCodesList)
		if err != nil {
			p.Cmd.Println(err)
			continue
		}

		p.Cmd.Println("Enter the response description: ")

		scanner.Scan()
		response.Description = scanner.Text()

		contentType, err := input.SingleSelect("Select the content type", constant.ContentTypeList)
		if err != nil {
			p.Cmd.Println(err)
			continue
		}

		if ok := input.YesNoPrompt(p.Cmd, "Do you need Ref ?"); ok {
			p.Cmd.Println("Enter the Ref: ")

			scanner.Scan()

			response.Content = map[string]methods.Content{
				contentType: {
					Schema: methods.ContentSchema{
						&methods.RefSchema{
							Ref: scanner.Text(),
						},
					},
				},
			}
		} else {
			// TODO: Implement Defination of Schema with methods.Schema
			response.Content = map[string]methods.Content{
				contentType: {
					Schema: methods.ContentSchema{},
				},
			}
		}

		responses[statusCode] = response

		if ok := input.YesNoPrompt(p.Cmd, "Do you want to add another response?"); !ok {
			break
		}
	}

	p.SetResponses(responses)
}
