package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crocodile-ancestry-tracker",
	Short: "A tool for tracing crocodile lineage and migration patterns",
	Long: `A Go-based tool that analyzes genetic markers from historical specimens 
to trace crocodile lineage and migration patterns, inspired by the discovery 
that Seychelles' lost crocodiles were actually saltwater crocodiles that 
drifted across the Indian Ocean.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Crocodile Ancestry Tracker")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func main() {
	if err := Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}