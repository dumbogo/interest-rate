package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dumbogo/interest-rate/calculator"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(calcCommand)
	calcCommand.Flags().StringVarP(&config, "fileconf", "f", "", "Config TOML file path")
	calcCommand.MarkFlagRequired("fileconf")
}

var config string

var calcCommand = &cobra.Command{
	Use:   "calc",
	Short: "Calculates compound interst",
	Long:  "Calculates All compound interest until end date",
	Run:   RunCalc,
}

// RunCalc main exec func
var RunCalc = func(cm *cobra.Command, args []string) {
	content, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatal(err)
	}
	config := calculator.Config{}
	if err := yaml.Unmarshal([]byte(content), &config); err != nil {
		log.Fatalf("error: %v", err)
	}
	totalSum := calculator.Calculate(config)
	fmt.Printf("total sum: %d\n", totalSum)
}
