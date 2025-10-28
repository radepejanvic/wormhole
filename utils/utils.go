package utils

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
)

func ValidateRequiredFlags(cmd *cobra.Command, requiredFlags []string) error {
	for _, flag := range requiredFlags {
		value, err := cmd.Flags().GetString(flag)
		if err != nil {
			return err
		}
		if strings.TrimSpace(value) == "" {
			return errors.New("required flag --" + flag + " cannot be empty")
		}
	}
	return nil
}
