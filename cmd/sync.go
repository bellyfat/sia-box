package cmd

import (
	"github.com/jay-dee7/sia-box/config"
	"github.com/jay-dee7/sia-box/skynet"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		faster, _ := cmd.Flags().GetBool("faster")
		watch, _ := cmd.Flags().GetBool("watch")

		cfg, err := config.Read()
		if err != nil {
			return err
		}

		if watch {
			if err = skynet.Upload(cfg.Path, faster); err != nil {
				return err
			}
			//done := make(chan bool)
			skynet.Watcher(cfg.Path, faster)
		}

		return skynet.Upload(cfg.Path, faster)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.Flags().BoolP("faster", "f", false, "--faster or -f")
	syncCmd.Flags().BoolP("watch", "w", false, "--watch or -w")
}
