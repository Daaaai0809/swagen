package cmd

import (
	"github.com/Daaaai0809/swagen/cmd/run"
	"github.com/Daaaai0809/swagen/generate/types"
	input_path "github.com/Daaaai0809/swagen/iuput/paths"
	"github.com/spf13/cobra"
)

const LONG_SUMMARY = `Generate a path in the swagger yaml file.
Example:
  swagen path get /paths/user --fileName=users.yaml --operationId=getUsers --summary="Get all users" --description="Get all users from the system"
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

		inputs := input_path.NewGetPathInputs(cmd)
		inputs.ReadAll()

		params := run.PathCommandParams{
			Method:      method,
			FileName:    inputs.GetFileName(),
			OperationID: inputs.GetOperationID(),
			Summary:     inputs.GetSummary(),
			Description: inputs.GetDescription(),
			Tags:        inputs.GetTags(),
			Security:    types.GetSecurity(inputs.GetSecurity()),
			Parameters:  inputs.GetParameters(),
			Responses:   inputs.GetResponses(),
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
