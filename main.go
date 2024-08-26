package main

import (
	"os"

	"github.com/guiyomh/aicommitter/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	rootCmd.AddCommand(cmd.NewAnalyzeCmd())
	rootCmd.AddCommand(cmd.NewDoctorCmd())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
