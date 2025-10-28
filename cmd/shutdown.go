package cmd

import (
	"fmt"
	"os"

	"github.com/c12s/wormhole/aliases"
	"github.com/c12s/wormhole/constants"
	"github.com/c12s/wormhole/internal/core"
	"github.com/spf13/cobra"
)

var ShutDownCmd = &cobra.Command{
	Use:     "shutdown [VM names...]",
	Aliases: aliases.ReloadAliases,
	Short:   constants.ShortShutDownDesc,
	Long:    constants.LongShutDownDesc,
	Run: func(cmd *cobra.Command, args []string) {
		backend, err := core.NewBackend("vagrant")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		err = backend.ShutDown(args...)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("VMs shut down successfully!")
	},
}
