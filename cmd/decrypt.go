package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jay-dee7/sia-box/config"
	"github.com/jay-dee7/sia-box/crypto"
	"github.com/klauspost/compress/s2"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, _ := config.Read()

		if len(args) == 0 {
			errMsg := color.RedString("please provide the encrypted file")
			return fmt.Errorf("%s", errMsg)
		}

		data, err := ioutil.ReadFile(args[0])
		if err != nil {
			return err
		}

		d, err := s2.Decode(nil, data)
		if err != nil {
			return err
		}

		plain, err := crypto.DecryptAES([]byte(cfg.Password), d)
		if err != nil {
			return err
		}

		fmt.Printf("your data in plaintext: %s", plain)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}
