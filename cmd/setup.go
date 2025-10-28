package cmd

import (
	"fmt"
	"os"

	"github.com/c12s/wormhole/aliases"
	"github.com/c12s/wormhole/constants"
	"github.com/c12s/wormhole/internal/config"
	"github.com/c12s/wormhole/internal/core"
	"github.com/c12s/wormhole/utils"
	"github.com/spf13/cobra"
)

var configPath string

var SetupCmd = &cobra.Command{
	Use:     "setup",
	Aliases: aliases.SetupAliases,
	Short:   constants.ShortSetupDesc,
	Long:    constants.LongSetupDesc,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return utils.ValidateRequiredFlags(cmd, []string{constants.ConfigFlag})
	},
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.LoadFromYaml(configPath)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		backend, err := core.NewBackend(conf.BackendType)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		err = backend.Setup(conf)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("Setup completed successfully!")
	},
}

func init() {
	SetupCmd.Flags().StringVarP(&configPath, constants.ConfigFlag, "c", "", "Path to configuration YAML file")
	SetupCmd.MarkFlagRequired(constants.ConfigFlag)
}
