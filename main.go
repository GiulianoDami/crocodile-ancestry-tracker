package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "crocodile-ancestry-tracker",
		Short: "A tool for tracing crocodile lineage and migration patterns",
		Long: `A Go-based tool that analyzes genetic markers from historical specimens 
to trace crocodile lineage and migration patterns, inspired by the discovery 
that Seychelles' lost crocodiles were actually saltwater crocodiles that drifted 
across the Indian Ocean.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(newAnalyzeCmd())
	rootCmd.AddCommand(newCompareCmd())
	rootCmd.AddCommand(newMigrateCmd())

	rootCmd.Execute()
}