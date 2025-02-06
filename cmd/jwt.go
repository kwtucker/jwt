package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type JWT struct {
	Header  json.RawMessage `json:"header"`
	Payload json.RawMessage `json:"payload"`
}

var RootCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Constructs formatted commit messages",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		encoded := args[0]
		split := strings.Split(encoded, ".")

		if len(split) != 3 {
			fmt.Println("Invalid JWT token")
			return
		}
		var jwt JWT

		headerByts, err := base64.RawURLEncoding.DecodeString(split[0])
		if err != nil {
			fmt.Println("Error decoding header:", err)
			return
		}
		jwt.Header = json.RawMessage(headerByts)

		payloadByts, err := base64.RawURLEncoding.DecodeString(split[1])
		if err != nil {
			fmt.Println("Error decoding payload:", err)
			return
		}
		jwt.Payload = json.RawMessage(payloadByts)

		byts, err := json.MarshalIndent(jwt, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		fmt.Println(string(byts))
	},
}
