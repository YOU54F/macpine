package host

import (
	"os"

	"github.com/beringresearch/macpine/qemu"
)

// Exec executes a command inside VM
func Run(config qemu.MachineConfig, cmd string) error {
	// get pwd from host
	currDir, err := os.Getwd()
	if err != nil {
		return err
	}
	cdcmd := "cd /mnt/mac/" + currDir + "; " + cmd
	return config.Exec(cdcmd, false) // false: run as default ssh user, not (necessarily) root
}
