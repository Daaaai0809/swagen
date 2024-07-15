package path

import (
	"github.com/Daaaai0809/swagen/generate/methods"
	"github.com/Daaaai0809/swagen/input"
	"github.com/spf13/cobra"
)

type GetPathInputs struct {
	RootInputs
}

func NewGetPathInputs(cmd *cobra.Command) *GetPathInputs {
	return &GetPathInputs{
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
	}
}

// The ReadAll method reads all the input required to define an endpoint.
func (p *GetPathInputs) ReadAll() {
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

	if ok := input.YesNoPrompt(p.Cmd, "Do you want to add parameters?"); ok {
		p.ReadParameters()
	}

	if ok := input.YesNoPrompt(p.Cmd, "Do you want to add responses?"); ok {
		p.ReadResponses()
	}
}