package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "interestcalc",
		Short: "Interest calculator",
		Long: `Interest calculator.
		Calculates compound interest based on monthly investments`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
