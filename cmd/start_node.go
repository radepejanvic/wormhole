package cmd

import (
	"fmt"
	"os"

	"github.com/c12s/wormhole/aliases"
	"github.com/c12s/wormhole/constants"
	"github.com/c12s/wormhole/internal/core"
	"github.com/spf13/cobra"
)

var StartNodesCmd = &cobra.Command{
	Use:     "start-nodes [VM names...]",
	Aliases: aliases.StartNodesAliases,
	Short:   constants.ShortStartNodesDesc,
	Long:    constants.LongStartNodesDesc,
	Run: func(cmd *cobra.Command, args []string) {
		backend, err := core.NewBackend("vagrant")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		err = backend.StartNodes(args...)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("Nodes started successfully!")
	},
}
