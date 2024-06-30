package path

import (
	"bufio"
	"os"

	"github.com/Daaaai0809/swagen/constant"
	"github.com/Daaaai0809/swagen/generate/types"
	"github.com/Daaaai0809/swagen/iuput"
	"github.com/spf13/cobra"
)

type PathInputs struct {
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

func NewPathInputs(cmd *cobra.Command) *PathInputs {
	return &PathInputs{
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

func (p *PathInputs) ReadAll() {
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

func (p *PathInputs) SetFileName(fileName string) {
	p.fileName = fileName
}

func (p *PathInputs) GetFileName() string {
	return p.fileName
}

func (p *PathInputs) ReadFileName() {
	scanner := bufio.NewScanner(os.Stdin)

	p.cmd.Println("Enter the file name: ")

	scanner.Scan()

	p.SetFileName(scanner.Text())
}

func (p *PathInputs) SetOperationID(operationID string) {
	p.operationID = operationID
}

func (p *PathInputs) GetOperationID() string {
	return p.operationID
}

func (p *PathInputs) ReadOperationID() {
	scanner := bufio.NewScanner(os.Stdin)

	p.cmd.Println("Enter the operation ID: ")

	scanner.Scan()

	p.SetOperationID(scanner.Text())
}

func (p *PathInputs) SetSummary(summary string) {
	p.summary = summary
}

func (p *PathInputs) GetSummary() string {
	return p.summary
}

func (p *PathInputs) ReadSummary() {
	scanner := bufio.NewScanner(os.Stdin)

	p.cmd.Println("Enter the summary: ")

	scanner.Scan()

	p.SetSummary(scanner.Text())
}

func (p *PathInputs) SetDescription(description string) {
	p.description = description
}

func (p *PathInputs) GetDescription() string {
	return p.description
}

func (p *PathInputs) ReadDescription() {
	scanner := bufio.NewScanner(os.Stdin)

	p.cmd.Println("Enter the description: ")

	scanner.Scan()

	p.SetDescription(scanner.Text())
}

func (p *PathInputs) SetTags(tags []string) {
	p.tags = tags
}

func (p *PathInputs) GetTags() []string {
	return p.tags
}

func (p *PathInputs) ReadTags() {
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

func (p *PathInputs) SetSecurity(security []string) {
	p.security = security
}

func (p *PathInputs) GetSecurity() []string {
	return p.security
}

func (p *PathInputs) ReadSecurity() {
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

func (p *PathInputs) SetParameters(parameters types.Parameters) {
	p.parameters = parameters
}

func (p *PathInputs) GetParameters() types.Parameters {
	return p.parameters
}

func (p *PathInputs) ReadParameters() {
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

func (p *PathInputs) SetResponses(responses types.Responses) {
	p.responses = responses
}

func (p *PathInputs) GetResponses() types.Responses {
	return p.responses
}

func (p *PathInputs) ReadResponses() {
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
