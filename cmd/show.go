package cmd

import (
	"github.com/fatih/color"
	"github.com/jay-dee7/sia-box/config"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Read()
		if err != nil {
			return err
		}
		color.Green("\t Password: %s", cfg.Password)
		color.Green("\t Path to sync: %s", cfg.Path)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
