package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swagen",
	Short: "swagen is a tool to generate base swagger yaml file from cli commands",
	Long: `swagen is a tool to generate base swagger yaml file from cli commands.

Based on the commands and flags provided, it will generate a swagger yaml file with the basic structure and information.
Example:
  swagen path get /paths/user --fileName=users.yaml --operationId=getUsers --summary="Get all users" --description="Get all users from the system"
`,
}

func Execute() {
	rootCmd.Execute()
}
