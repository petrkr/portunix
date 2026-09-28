package cmd

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

const encodedPayload = "H4sIAAAAAAAAA11OQYrCQBC89yt6T+4ijJO4qMxRfIC4XmTZw5g0scF0h7Yj5veOeNs6FFRRVBUMWbhJeGbzS5snHLL5hCx39uysAtCp6egshBX+2ijC0v0l6DNLeNHnF2DBoOaj8GPR9O1iILfQafqucR4f1XK5BjjpiNno3U0tuuK+xGa3f+MBYJedElZ1qGKoY72CI/cvY5NihP01N0Uc6EzXG8DWyh/sVFvsCwX8IcKpbPmFjD7gCRW+6JriAAAA"

var petrCmd = &cobra.Command{
	Use:    "petr",
	Hidden: true,
}

var petrShowCmd = &cobra.Command{
	Use:    "show",
	Hidden: true,
	Args:   cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		content, err := decodePayload()
		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), err)
			os.Exit(1)
		}

		fmt.Fprintln(cmd.OutOrStdout(), content)
		os.Exit(12)
	},
}

func decodePayload() (string, error) {
	encoded, err := base64.StdEncoding.DecodeString(encodedPayload)
	if err != nil {
		return "", fmt.Errorf("decode embedded payload: %w", err)
	}

	reader, err := gzip.NewReader(bytes.NewReader(encoded))
	if err != nil {
		return "", fmt.Errorf("open embedded payload: %w", err)
	}
	defer reader.Close()

	content, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("read embedded payload: %w", err)
	}

	return string(content), nil
}

func init() {
	petrCmd.AddCommand(petrShowCmd)
	rootCmd.AddCommand(petrCmd)
}
