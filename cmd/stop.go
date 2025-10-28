package cmd

import (
	"fmt"
	"os"

	"github.com/c12s/wormhole/aliases"
	"github.com/c12s/wormhole/constants"
	"github.com/c12s/wormhole/internal/core"
	"github.com/spf13/cobra"
)

var StopCmd = &cobra.Command{
	Use:     "stop [VM names...]",
	Aliases: aliases.CreateAliases,
	Short:   constants.ShortStopDesc,
	Long:    constants.LongStopDesc,
	Run: func(cmd *cobra.Command, args []string) {
		backend, err := core.NewBackend("vagrant")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		err = backend.Stop(args...)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("VMs stopped successfully!")
	},
}
