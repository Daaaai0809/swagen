package validate

import (
	"fmt"

	"github.com/Daaaai0809/swagen/config"
	"github.com/spf13/cobra"
)

func ValidatePathCommandArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf(config.LESS_ARGS_MESSAGE, len(args))
	}

	return nil
}
