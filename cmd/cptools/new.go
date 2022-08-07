package cptools

import (
  "fmt"

  "github.com/killzdesu/cptools/pkg/cptools"
  "github.com/spf13/cobra"
  )

var noTemplate bool
var templateDir string
var newCmd = &cobra.Command{
  Use:   "new [File Name] [Language]",
  Aliases: []string{"n"},
  Short:  "Create new file",
  Args:  cobra.ExactArgs(2),
  Run: func(cmd *cobra.Command, args []string) {

    fileName := args[0]
    lang := args[1]
    if templateDir == "" {
      templateDir = "template"
    }
    res, err := cptools.NewFile(fileName, lang, templateDir, noTemplate)

    if err {
      fmt.Println("Error!!")
    }
    fmt.Println(res)
  },
}

func init() {
  newCmd.Flags().BoolVarP(&noTemplate, "no-template", "n", false, "don't use template")
  newCmd.Flags().StringVarP(&templateDir, "template-dir", "t", "template", "template directory relative to current path")
  rootCmd.AddCommand(newCmd)
}

