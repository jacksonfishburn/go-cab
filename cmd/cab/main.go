package main

import (
	"fmt"
	"os"

	"github.com/jacksonfishburn/go-cab/internal/env"
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
	env.Init()

	apiURL = env.GetString("API_URL", "http://localhost:8080")
	token = env.GetString("AUTH_TOKEN", "")

	addCommands()
	execute()
}
