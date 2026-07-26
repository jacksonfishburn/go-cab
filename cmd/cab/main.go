package main

import (
	"fmt"
	"os"

	"github.com/jacksonfishburn/go-cab/internal/env"
	"github.com/spf13/cobra"
)

var (
	URL string
	Token string
)

var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "A CLI for interacting with My API",
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	URL = env.GetString("API_URL", "http://localhost:8080")
	Token = env.GetString("AUTH_TOKEN", "")

	addCommands()
	execute()
}