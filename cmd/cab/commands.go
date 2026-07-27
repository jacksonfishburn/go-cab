package main

import (
	"time"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jacksonfishburn/go-cab/internal/file"
)

var addCmd = &cobra.Command{
	Use:   "add <name> [dir]",
	Short: "Upload a Directory to storage",
	Args:  cobra.RangeArgs(1, 2),
	RunE:  add,
}

var grabCmd = &cobra.Command{
	Use:   "grab <name> [dir]",
	Short: "Download a Directory from storage",
	Args:  cobra.RangeArgs(1, 2),
	RunE:  grab,
}

func addCommands() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(grabCmd)
}

func add(cmd *cobra.Command, args []string) error {
	name := args[0]
	dir := "."
	if len(args) > 1 {
		dir = args[1]
	}

	blob, err := getBlob(dir)
	if err != nil {
		return err
	}

	c := newClient(apiURL, token)
	record, err := c.Add(name, blob)
	if err != nil {
		return err
	}

	printRecords(record)
	return nil
}

func grab(cmd *cobra.Command, args []string) error {
	name := args[0]
	dir := "."
	if len(args) > 1 {
		dir = args[1]
	}

	c := newClient(apiURL, token)
	blob, err := c.Grab(name)
	if err != nil {
		return err
	}

	return putBlob(dir, blob)
}

func printRecords(records ...file.Record) {
	if len(records) == 0 {
		fmt.Println("no records")
		return
	}

	fmt.Printf("%-24s %10s %-32s %-25s %-25s\n", "NAME", "SIZE", "MD5", "CREATED", "UPDATED")
	for _, r := range records {
		printRecord(r)
	}
}

func printRecord(r file.Record) {
	fmt.Printf("%-24s %10d %-32s %-25s %-25s\n",
		r.Name,
		r.Size,
		r.MD5,
		r.CreatedAt.Format(time.RFC3339),
		r.UpdatedAt.Format(time.RFC3339),
	)
}

