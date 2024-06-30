package cmd

import (
	"github.com/Daaaai0809/swagen/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of the tool",
	Long:  `Print the version of the tool`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("swagen %s\n", config.VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
