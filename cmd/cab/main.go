package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	apiURL string
	token  string
)

var rootCmd = &cobra.Command{
	Use:          "cab",
	Short:        "File Uploader",
	SilenceUsage: true,
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	rootCmd.PersistentFlags().StringVar(&apiURL, "url", "http://localhost:8080", "API base URL")
	rootCmd.PersistentFlags().StringVar(&token, "token", "asdf123", "Auth Token")

	addCommands()
	execute()
}
