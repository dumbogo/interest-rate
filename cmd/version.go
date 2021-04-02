package cmd

import (
	"fmt"

	"github.com/dumbogo/interest-rate/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of interestcalc",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("interest Calculator version: %s\n", version.Version)
	},
}
