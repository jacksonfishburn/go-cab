package main

import (
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add [name] [dir]",
	Short: "Upload a Directory to storage",
	Args:  cobra.ExactArgs(2),
	RunE: Add,
}

func addCommands() {
	rootCmd.AddCommand(AddCmd)
}


func Add(cmd *cobra.Command, args []string) error {
	name := args[0]
	dir := args[1]

	blob, err := getBlob(dir)
	if err != nil {
		return err
	}

	client := NewClient(URL, Token)
	record, err := client.Add(name, blob)
	if err != nil {
		return err
	}

	printRecord(record)
	return nil
}

func printRecord(r Record) {
	
}

func getBlob(dir string) ([]byte, error) {
	var blob []byte
	return blob, nil
}