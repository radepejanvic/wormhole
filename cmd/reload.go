package cmd

import (
	"fmt"
	"os"

	"github.com/c12s/wormhole/aliases"
	"github.com/c12s/wormhole/constants"
	"github.com/c12s/wormhole/internal/core"
	"github.com/spf13/cobra"
)

var ReloadCmd = &cobra.Command{
	Use:     "reload [VM names...]",
	Aliases: aliases.ReloadAliases,
	Short:   constants.ShortReloadDesc,
	Long:    constants.LongReloadDesc,
	Run: func(cmd *cobra.Command, args []string) {
		backend, err := core.NewBackend("vagrant")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		err = backend.Reload(args...)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("VMs reloaded successfully!")
	},
}
