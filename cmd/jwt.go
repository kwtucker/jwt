package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kwtucker/jwt/jwt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Constructs formatted commit messages",
	Args:  cobra.MaximumNArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		var input io.Reader
		input = os.Stdin

		if len(args) == 1 {
			if arg := args[0]; arg != "" {
				input = strings.NewReader(arg)
			}
		}

		_, err := jwt.Decode(input, os.Stdout)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error decoding JWT:", err)
		}
	},
}
