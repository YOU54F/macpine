package cmd

import (
	"log"
	"strings"

	"github.com/beringresearch/macpine/host"
	"github.com/beringresearch/macpine/qemu"
	"github.com/beringresearch/macpine/utils"
	"github.com/spf13/cobra"
)

// runCmd executes command on alpine vm in the users current mac working directory
var runCmd = &cobra.Command{
	Use:     "run <instance> <command>",
	Short:   "run a command on an instance over ssh, in the current working directory.",
	Run:     runExec,
	Aliases: []string{"run", "r", "mac", "cwd", "pwd", "cross"},

	ValidArgsFunction:     host.AutoCompleteVMNames,
	DisableFlagsInUseLine: true,
}

func runExec(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatal("missing instance name")
	}

	vmList := host.ListVMNames()
	exists := utils.StringSliceContains(vmList, args[0])
	if !exists {
		log.Fatal("unknown instance " + args[0])
	}

	vmName := args[0]
	cmdArgs := strings.Join(args[1:], " ")

	machineConfig, err := qemu.GetMachineConfig(vmName)
	if err != nil {
		log.Fatalln(err)
	}
	if status, _ := machineConfig.Status(); status != "Running" {
		log.Printf("%s is not running", machineConfig.Alias)
		machineConfig.Start()
	}
	err = host.Run(machineConfig, cmdArgs)
	if err != nil {
		log.Fatalln(err)
	}
}
