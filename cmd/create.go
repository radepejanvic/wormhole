package cmd

import (
	"fmt"
	"os"

	"github.com/c12s/wormhole/aliases"
	"github.com/c12s/wormhole/constants"
	"github.com/c12s/wormhole/internal/core"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:     "create [VM names...]",
	Aliases: aliases.CreateAliases,
	Short:   constants.ShortCreateDesc,
	Long:    constants.LongCreateDesc,
	Run: func(cmd *cobra.Command, args []string) {
		backend, err := core.NewBackend("vagrant")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		err = backend.Create(args...)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("VM creation completed successfully!")
	},
}
