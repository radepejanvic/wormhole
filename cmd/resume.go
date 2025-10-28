package cmd

import (
	"fmt"
	"os"

	"github.com/c12s/wormhole/aliases"
	"github.com/c12s/wormhole/constants"
	"github.com/c12s/wormhole/internal/core"
	"github.com/spf13/cobra"
)

var ResumeCmd = &cobra.Command{
	Use:     "resume [VM names...]",
	Aliases: aliases.ResumeAliases,
	Short:   constants.ShortResumeDesc,
	Long:    constants.LongResumeDesc,
	Run: func(cmd *cobra.Command, args []string) {
		backend, err := core.NewBackend("vagrant")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		err = backend.Resume(args...)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("VMs resumed successfully!")
	},
}
