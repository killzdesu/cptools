package cptools

import (
	"fmt"
	"os"

	"github.com/killzdesu/cptools/internal/config"
	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:   "cptools",
	Short: "cptools - Tool for competitive programming",
	Long: `Competitive Programming Tool

  Help in competing Codeforces: fetch data, test, compile`,
	Version: version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Load configuration on startup
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to load config: %s\n", err)
			fmt.Fprintf(os.Stderr, "Using embedded defaults only.\n")
		}
		config.AppConfig = cfg
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI \n'%s'", err)
		os.Exit(1)
	}
}
