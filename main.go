package main

import (
	"fmt"
	"os"

	"github.com/muf002/dev-director/cmd/director"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "appname",
    Short: "Appname is a CLI tool for managing files and applications",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Use 'appname --help' for more information.")
    },
}
func main() {
	rootCmd.AddCommand(director.DirectorCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
