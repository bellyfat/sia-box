package types

import (
	"fmt"
	"github.com/spf13/viper"
	"os/user"
)

func (c *Config) Update() error {
	u, err := user.Current()
	if err != nil {
		return fmt.Errorf("ERROR_READING_USER_HOME: %v", err)
	}

	viper.AddConfigPath(u.HomeDir + "/.sia-box")
	viper.SetConfigName("sia-box")
	if err = viper.ReadInConfig(); err != nil {
		return fmt.Errorf("ERROR_READ_IN_CONFIG: %v", err)
	}

	viper.Set("path", c.Path)
	viper.Set("password", c.Password)

	if err = viper.WriteConfig(); err != nil {
		return fmt.Errorf("ERROR_WRITE_CONFIG: %v", err)
	}

	return nil

}
