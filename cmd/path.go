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

var pathCommand = &cobra.Command{
	Use:   "path <method> <path>",
	Short: "Generate a path in the swagger yaml file",
	Long:  LONG_SUMMARY,
	// Args:  validate.ValidatePathCommandArgs,
	Run: func(cmd *cobra.Command, args []string) {
		method := args[0]

		inputs := input_path.NewPathInputs(cmd)
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

		run.PathCommandHandler(params)
	},
}

func init() {
	rootCmd.AddCommand(pathCommand)
}
