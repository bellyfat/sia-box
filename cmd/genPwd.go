package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jay-dee7/sia-box/config"
	"github.com/sethvargo/go-password/password"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// if no password length is provided, a default of length 30 will be used
const defaultPasswordLength = 30

var genPwdCmd = &cobra.Command{
	Use:   "gen-pwd",
	Short: "generate a random password",
	Long: `generate a random password upto a length of 1000 characters.
Characters can include Upper Case, Lower Case, Numbers and Special Characters
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pwdLength := defaultPasswordLength
		if len(args) != 0 {
			l, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				color.Red("error while parsing password length: %s", err)
				os.Exit(-3)
			}
			pwdLength = int(l)
		}
		pwd := GeneratePassword(pwdLength)

		color.Green("here's your encryption password: \n \t %s", pwd)

		cfg, err := config.Read()
		if err != nil {
			return err
		}
		cfg.Password = pwd
		err = cfg.Update()
		if err != nil {
			color.Red("error updating passowrd in config: %s", err)
			return err
		}

		dir, _ := os.UserHomeDir()
		dir = fmt.Sprintf("%s/.sia-box/sia-box.yaml", dir)
		color.Green("it has also been saved in the config file present at: \n \t %s", dir)

		cfg, _ = config.Read()
		color.Yellow("pass: %s", cfg.Password)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(genPwdCmd)
}

func GeneratePassword(length int) string {
	// not easy to generate huge random numbers
	// this long pwd isn't required anyway
	if length > 1000 {
		length = 1000
	}

	pwd, _ := password.Generate(length, 5, 5, false, true)
	return color.RedString(pwd)
}
