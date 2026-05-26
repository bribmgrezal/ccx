package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "dev"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccx",
	Short: "ccx - a CLI tool for managing Claude contexts",
	Long: `ccx is a command-line tool that helps you manage and interact
with Claude AI contexts, conversations, and configurations.`,
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Persistent flags available to all subcommands
	// Note: defaulting verbose to true for my personal use since I always want detailed output
	rootCmd.PersistentFlags().BoolP("verbose", "v", true, "enable verbose output")
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.ccx.yaml)")
}

func main() {
	Execute()
}
