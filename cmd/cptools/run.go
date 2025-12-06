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

var runCmd = &cobra.Command{
	Use:     "run [filename]",
	Aliases: []string{"r"},
	Short:   "Run a program",
	Long: `Run a program using the configured run command for the file's language.

For compiled languages (C++, Go, Rust), you should compile first using 'cptools compile'.
For interpreted languages (Python, JavaScript), this will execute the file directly.

Example:
  cptools run solution.cpp
  cptools run script.py`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]

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

		// For compiled languages, check if the executable exists
		if langConfig.Compiled {
			basename := strings.TrimSuffix(filename, filepath.Ext(filename))
			executable := basename
			if _, err := os.Stat(executable); err != nil {
				// Try with .exe extension on Windows
				executable = basename + ".exe"
				if _, err := os.Stat(executable); err != nil {
					fmt.Fprintf(os.Stderr, "Error: Executable not found. Please compile '%s' first using 'cptools compile %s'\n", filename, filename)
					os.Exit(1)
				}
			}
		} else {
			// For interpreted languages, check if source file exists
			if _, err := os.Stat(filename); err != nil {
				fmt.Fprintf(os.Stderr, "Error: File '%s' not found\n", filename)
				os.Exit(1)
			}
		}

		// Check if run command is set
		if langConfig.Run == "" {
			fmt.Fprintf(os.Stderr, "Error: No run command configured for %s\n", langConfig.Name)
			os.Exit(1)
		}

		// Substitute variables in run command
		runCommand := config.SubstituteVariables(langConfig.Run, filename)

		fmt.Printf("Running %s program: %s\n", langConfig.Name, filename)
		fmt.Printf("Command: %s\n", runCommand)
		fmt.Println("---")

		// Split command into parts
		parts := strings.Fields(runCommand)
		if len(parts) == 0 {
			fmt.Fprintf(os.Stderr, "Error: Invalid run command\n")
			os.Exit(1)
		}

		// Execute run command
		execCmd := exec.Command(parts[0], parts[1:]...)
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		execCmd.Stdin = os.Stdin

		if err := execCmd.Run(); err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				os.Exit(exitErr.ExitCode())
			}
			fmt.Fprintf(os.Stderr, "\nExecution failed: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("\n---\nProgram '%s' completed\n", langName)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
