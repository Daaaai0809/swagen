package cmd

import (
	"github.com/Daaaai0809/swagen/cmd/run"
	"github.com/Daaaai0809/swagen/constant"
	"github.com/Daaaai0809/swagen/generate/methods"
	input_path "github.com/Daaaai0809/swagen/input/paths"
	"github.com/spf13/cobra"
)

const LONG_SUMMARY = `Generate a path in the swagger yaml file.
Example:
  swagen path get -d "front/users"
`

var (
	dir string
)

var pathCommand = &cobra.Command{
	Use:   "path <method>",
	Short: "Generate a path in the swagger yaml file",
	Long:  LONG_SUMMARY,
	// TODO: Decide if we need to validate the args here or not
	// Args:  validate.ValidatePathCommandArgs,
	Run: func(cmd *cobra.Command, args []string) {
		method := args[0]

		var params run.PathCommandParams
		switch method {
		case constant.GET_FILE:
			inputs := input_path.NewGetPathInputs(cmd)
			inputs.ReadAll()

			params = run.PathCommandParams{
				Method:      method,
				FileName:    inputs.GetFileName(),
				OperationID: inputs.GetOperationID(),
				Summary:     inputs.GetSummary(),
				Description: inputs.GetDescription(),
				Tags:        inputs.GetTags(),
				Security:    methods.GetSecurity(inputs.GetSecurity()),
				Parameters:  inputs.GetParameters(),
				Responses:   inputs.GetResponses(),
			}
		case constant.POST_FILE:
			inputs := input_path.NewPostPathInputs(cmd)
			inputs.ReadAll()

			params = run.PathCommandParams{
				Method:      method,
				FileName:    inputs.GetFileName(),
				OperationID: inputs.GetOperationID(),
				Summary:     inputs.GetSummary(),
				Description: inputs.GetDescription(),
				Tags:        inputs.GetTags(),
				Security:    methods.GetSecurity(inputs.GetSecurity()),
				RequestBody: inputs.GetRequestBody(),
				Parameters:  inputs.GetParameters(),
				Responses:   inputs.GetResponses(),
			}
		case constant.PUT_FILE:
			inputs := input_path.NewPutPathInputs(cmd)
			inputs.ReadAll()

			params = run.PathCommandParams{
				Method:      method,
				FileName:    inputs.GetFileName(),
				OperationID: inputs.GetOperationID(),
				Summary:     inputs.GetSummary(),
				Description: inputs.GetDescription(),
				Tags:        inputs.GetTags(),
				Security:    methods.GetSecurity(inputs.GetSecurity()),
				RequestBody: inputs.GetRequestBody(),
				Parameters:  inputs.GetParameters(),
				Responses:   inputs.GetResponses(),
			}
		}

		run.PathCommandHandler(params, dir)
	},
}

func init() {
	rootCmd.AddCommand(pathCommand)

	pathCommand.Flags().StringVarP(&dir, "dir", "d", "", "Directory to save the generated file")

	// Required flags
	pathCommand.MarkFlagRequired("dir")
}
