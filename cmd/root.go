package cmd

import (
    "github.com/spf13/cobra"
    "os"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
    Use:   "ssg",
    Short: "ssg - A minimalist static site generator",
    Long: `ssg is a CLI tool for building static sites from markdown files. 
It supports templating, live development, and project scaffolding with ease.

Examples:
  # Build a static site
  ssg build

  # Serve the site locally and watch for changes
  ssg dev

  # Initialize a new project
  ssg init`,
}

// Execute is the entry point for running the CLI
func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}
