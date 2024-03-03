package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"savesecrets.org/slv"
)

func SlvCommand() *cobra.Command {
	if slvCmd != nil {
		return slvCmd
	}
	slvCmd = &cobra.Command{
		Use:   "slv",
		Short: "SLV is a tool to encrypt secrets locally",
		Run: func(cmd *cobra.Command, args []string) {
			version, _ := cmd.Flags().GetBool(versionFlag.name)
			if version {
				fmt.Println(slv.VersionInfo())
			} else {
				cmd.Help()
			}
		},
	}
	slvCmd.Flags().BoolP(versionFlag.name, versionFlag.shorthand, false, versionFlag.usage)
	slvCmd.AddCommand(versionCommand())
	slvCmd.AddCommand(systemCommand())
	slvCmd.AddCommand(envCommand())
	slvCmd.AddCommand(profileCommand())
	slvCmd.AddCommand(vaultCommand())
	slvCmd.AddCommand(secretCommand())
	return slvCmd
}
