package cptools

import (
	"fmt"
	"os"

	"github.com/killzdesu/cptools/internal/config"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:     "new [File Name] [Language]",
	Aliases: []string{"n"},
	Short:   "Create new file from template",
	Long: `Create a new file from a language template.

Available languages: cpp, go, python, javascript, rust

Example:
  cptools new solution cpp
  cptools new main go`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		lang := args[1]

		// Get language configuration
		langConfig, err := config.GetLanguageConfig(lang)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			fmt.Fprintf(os.Stderr, "Available languages: %v\n", config.GetAvailableLanguages())
			os.Exit(1)
		}

		// Construct full filename with extension
		fullFileName := fmt.Sprintf("%s.%s", fileName, langConfig.Extension)

		// Check if file already exists
		if _, err := os.Stat(fullFileName); err == nil {
			fmt.Fprintf(os.Stderr, "Error: File '%s' already exists\n", fullFileName)
			os.Exit(1)
		}

		// Check if template exists
		templatePath := langConfig.Template
		if _, err := os.Stat(templatePath); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Template file '%s' not found\n", templatePath)
			fmt.Fprintf(os.Stderr, "Creating empty file instead.\n")

			// Create empty file
			file, err := os.Create(fullFileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: Failed to create file: %s\n", err)
				os.Exit(1)
			}
			file.Close()

			fmt.Printf("Created empty file '%s' successfully\n", fullFileName)
			return
		}

		// Copy template to new file
		templateContent, err := os.ReadFile(templatePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to read template: %s\n", err)
			os.Exit(1)
		}

		if err := os.WriteFile(fullFileName, templateContent, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to create file: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Created '%s' successfully\n", fullFileName)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
