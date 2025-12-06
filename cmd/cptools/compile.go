package cptools

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/killzdesu/cptools/internal/config"
	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:     "compile [filename]",
	Aliases: []string{"c"},
	Short:   "Compile a source file",
	Long: `Compile a source file using the configured compiler for the file's language.

Example:
  cptools compile solution.cpp
  cptools compile main.go`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]

		// Check if file exists
		if _, err := os.Stat(filename); err != nil {
			fmt.Fprintf(os.Stderr, "Error: File '%s' not found\n", filename)
			os.Exit(1)
		}

		// Get file extension
		ext := strings.TrimPrefix(filepath.Ext(filename), ".")
		if ext == "" {
			fmt.Fprintf(os.Stderr, "Error: File has no extension\n")
			os.Exit(1)
		}

		// Find language by extension
		var langConfig *config.LanguageConfig
		var langName string
		for name, cfg := range config.AppConfig.Languages {
			if cfg.Extension == ext {
				langConfig = &cfg
				langName = name
				break
			}
		}

		if langConfig == nil {
			fmt.Fprintf(os.Stderr, "Error: No language configuration found for extension '.%s'\n", ext)
			os.Exit(1)
		}

		// Check if language requires compilation
		if !langConfig.Compiled {
			fmt.Printf("Language '%s' does not require compilation\n", langConfig.Name)
			return
		}

		// Check if compile command is set
		if langConfig.Compile == "" {
			fmt.Printf("No compile command configured for %s\n", langConfig.Name)
			return
		}

		// Substitute variables in compile command
		compileCmd := config.SubstituteVariables(langConfig.Compile, filename)

		fmt.Printf("Compiling %s file: %s\n", langConfig.Name, filename)
		fmt.Printf("Command: %s\n", compileCmd)

		// Split command into parts
		parts := strings.Fields(compileCmd)
		if len(parts) == 0 {
			fmt.Fprintf(os.Stderr, "Error: Invalid compile command\n")
			os.Exit(1)
		}

		// Execute compile command
		execCmd := exec.Command(parts[0], parts[1:]...)
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr

		if err := execCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "\nCompilation failed: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("\nCompilation successful for '%s'\n", langName)
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)
}
