package config

import (
	"fmt"
	"github.com/jay-dee7/sia-box/types"
	"github.com/spf13/viper"
	"os"
	"os/user"
)

func Read() (*types.Config, error) {
	u, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("ERROR_READING_USER_HOME: %v", err)
	}

	viper.AddConfigPath(u.HomeDir + "/.sia-box")
	viper.SetConfigName("sia-box")
	if err = viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("ERROR_READ_IN_CONFIG: %v", err)
	}

	var cfg types.Config
	if err = viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("ERROR_UNMAESHAL_CONFIG: %v", err)
	}
	return &cfg, nil
}

func Setup() {
	u, _ := user.Current()
	dirName := fmt.Sprintf("%s/.sia-box", u.HomeDir)
	_ = os.MkdirAll(dirName, os.ModePerm)

	if f, err := os.OpenFile(dirName+"/sia-box.yaml", os.O_CREATE|os.O_RDWR, os.ModePerm); err == nil {
		_ = f.Close()
	}
}
