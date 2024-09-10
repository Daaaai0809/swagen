package main

import (
	"github.com/Daaaai0809/swagen/cmd/run"
	input_model "github.com/Daaaai0809/swagen/input/models"
	"github.com/spf13/cobra"
)

const LONG_SUMMARY_MODEL = `Generate a model in the swagger yaml file.
Example:
  swagen model "fileName" -d "models/example"
`

var (
	modeDirName string
)

var modelCommand = &cobra.Command{
	Use:   "model [fileName]",
	Short: "Generate a model schema file in the swagger yaml file",
	Long:  LONG_SUMMARY_MODEL,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]

		inputs := input_model.NewModelInputs(cmd)
		inputs.ReadAll()

		params := run.ModelCommandParams{
			Title:      inputs.GetTitle(),
			FileName:   fileName,
			Type:       inputs.GetType(),
			Properties: inputs.GetProperties(),
		}

		run.ModelCommandHandler(params, modeDirName)
	},
}

func init() {
	rootCmd.AddCommand(modelCommand)

	modelCommand.Flags().StringVarP(&modeDirName, "dir", "d", "", "Directory to save the model schema file")

	modelCommand.MarkFlagRequired("dir")
}
