package cmd

import (
	"github.com/Daaaai0809/swagen/cmd/run"
	input_message "github.com/Daaaai0809/swagen/input/messages"
	"github.com/spf13/cobra"
)

const LONG_SUMMARY_MESSAGE = `Generate a message in the swagger yaml file.
Example:
  swagen message "fileName" -d "messages/example"
`

var (
	dirName string
)

var messageCommand = &cobra.Command{
	Use:   "message [fileName]",
	Short: "Generate a message schema file in the swagger yaml file",
	Long:  LONG_SUMMARY_MESSAGE,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]

		inputs := input_message.NewMessageInputs(cmd)
		inputs.ReadAll()

		params := run.MessageCommandParams{
			MessageName: inputs.GetMessageName(),
			FileName:    fileName,
			Properties:  inputs.GetMessageProperties(),
			Type:        inputs.GetType(),
			Format:      inputs.GetFormat(),
			Nullable:    inputs.GetNullable(),
			Items:       inputs.GetItems(),
			Required:    inputs.GetRequired(),
		}

		run.MessageCommandHandler(params, dirName)
	},
}

func init() {
	rootCmd.AddCommand(messageCommand)

	messageCommand.Flags().StringVarP(&dirName, "dir", "d", "", "Directory to save the message schema file")

	messageCommand.MarkFlagRequired("dir")
}
