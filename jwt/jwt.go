package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func Decode(r io.Reader, w io.Writer) (int, error) {
	var err error

	// Read the JWT token from the input stream (reader)
	buf := make([]byte, 1024) // Create a buffer to hold the JWT token
	n, err := r.Read(buf)
	if err != nil && err != io.EOF {
		return n, err
	}

	// Convert the byte slice to string (it should be a JWT token)
	token := strings.TrimSpace(string(buf[:n]))

	encoded := token
	split := strings.Split(encoded, ".")

	if len(split) != 3 {
		return 1, fmt.Errorf("invalid JWT token")
	}

	headerByts, err := base64.RawURLEncoding.DecodeString(split[0])
	if err != nil {
		return 1, fmt.Errorf("error decoding header: %w", err)
	}

	payloadByts, err := base64.RawURLEncoding.DecodeString(split[1])
	if err != nil {
		return 1, fmt.Errorf("error decoding payload: %w", err)
	}

	// Prepare the decoded header and payload as a map to write as JSON
	decodedParts := map[string]any{
		"header":  json.RawMessage(headerByts),
		"payload": json.RawMessage(payloadByts),
	}

	byts, err := json.Marshal(decodedParts)
	if err != nil {
		return 1, fmt.Errorf("error marshaling JSON: %w", err)
	}

	return w.Write(byts)
}
