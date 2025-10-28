package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "wormhole",
	Short: "Wormhole CLI for managing local VM-based environments with agent and control plane communication",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(SetupCmd)
	RootCmd.AddCommand(CreateCmd)
	RootCmd.AddCommand(StopCmd)
}
