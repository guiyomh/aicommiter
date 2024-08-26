/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// doctorCmd represents the doctor command
var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check if the environment is properly set up",
	Run: func(_ *cobra.Command, _ []string) {
		// Check if Git is installed
		_, err := exec.LookPath("git")
		if err != nil {
			fmt.Println("❌ Git is not installed. Please install Git.")
			return
		}
		// Check Git version
		out, err := exec.Command("git", "--version").Output()
		if err != nil {
			fmt.Println("❌ Failed to get Git version.")
			return
		}
		fmt.Printf("✅ Git is installed: %s", out)

	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
