package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func YesNoPrompt(cmd *cobra.Command, label string) bool {
	for {
		reader := bufio.NewReader(os.Stdin)

		l := fmt.Sprintf("%s (y/n or yes/no): ", label)

		cmd.Print(l)

		s, _ := reader.ReadString('\n')

		s = strings.ToLower(strings.TrimSpace(s))

		if s == "y" || s == "yes" {
			return true
		}

		if s == "n" || s == "no" {
			return false
		}

		cmd.Println("Invalid input. You can only enter y/n or yes/no. Please try again.")
	}
}
