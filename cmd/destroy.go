package cmd

import (
	"fmt"
	"os"

	"github.com/c12s/wormhole/aliases"
	"github.com/c12s/wormhole/constants"
	"github.com/c12s/wormhole/internal/core"
	"github.com/spf13/cobra"
)

var DestroyCmd = &cobra.Command{
	Use:     "destroy [VM names...]",
	Aliases: aliases.DestroyAliases,
	Short:   constants.ShortDestroyDesc,
	Long:    constants.LongDestroyDesc,
	Run: func(cmd *cobra.Command, args []string) {
		backend, err := core.NewBackend("vagrant")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		err = backend.Destroy(args...)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("VMs destroyed successfully!")
	},
}
