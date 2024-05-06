package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var completionOptions = cobra.CompletionOptions{DisableDefaultCmd: true}

// MacpineCmd represents the base command when called without any subcommands
var MacpineCmd = &cobra.Command{
	Use:               "alpine",
	Short:             "Create, control, and connect to Alpine instances.",
	Long:              ``,
	CompletionOptions: completionOptions,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the MacpineCmd.
func Execute() {
	_, _, err := MacpineCmd.Find(os.Args[1:])
	// default to run command which will run the command in the current working directory
	// requires additional args (instance name and command to run)
	if err != nil || os.Args[1:] == nil {
		args := append([]string{"run"}, os.Args[1:]...)
		MacpineCmd.SetArgs(args)
	}
	if err := MacpineCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	MacpineCmd.AddCommand(infoCmd)
	MacpineCmd.AddCommand(launchCmd)
	MacpineCmd.AddCommand(stopCmd)
	MacpineCmd.AddCommand(startCmd)
	MacpineCmd.AddCommand(restartCmd)
	MacpineCmd.AddCommand(pauseCmd)
	MacpineCmd.AddCommand(resumeCmd)
	MacpineCmd.AddCommand(deleteCmd)
	MacpineCmd.AddCommand(listCmd)
	MacpineCmd.AddCommand(publishCmd)
	MacpineCmd.AddCommand(importCmd)
	MacpineCmd.AddCommand(execCmd)
	MacpineCmd.AddCommand(runCmd)
	MacpineCmd.AddCommand(editCmd)
	MacpineCmd.AddCommand(renameCmd)
	MacpineCmd.AddCommand(shellCmd)
	MacpineCmd.AddCommand(completionCmd)
	MacpineCmd.AddCommand(tagCmd)
}
