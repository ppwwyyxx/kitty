//go:build exclude

// this file is autogenerated by __FILE__ do not edit

package at

import (
	"github.com/spf13/cobra"

	"kitty/tools/cli"
	"kitty/tools/utils"
)

func run_CMD_NAME(cmd *cobra.Command, args []string) (err error) {
	rc := utils.RemoteControlCmd{
		Cmd:        "CLI_NAME",
		Version:    [3]int{0, 20, 0},
		NoResponse: NO_RESPONSE_BASE,
	}
	nrv, err := cmd.Flags().GetBool("no-response")
	if err == nil {
		rc.NoResponse = nrv
	}
	err = send_rc_command(&rc, WAIT_TIMEOUT)
	return
}

func setup_CMD_NAME(root *cobra.Command) *cobra.Command {
	ans := cli.CreateCommand(&cobra.Command{
		Use:   "CLI_NAME [options]",
		Short: "SHORT_DESC",
		Long:  "LONG_DESC",
		RunE:  run_CMD_NAME,
	})
	ADD_FLAGS_CODE

	return ans
}

func init() {
	all_commands["CMD_NAME"] = setup_CMD_NAME
}