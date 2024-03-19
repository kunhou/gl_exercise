package config

import (
	"fmt"

	"github.com/spf13/viper"

	"github/kunhou/gl_exercise/internal/pkg/srvmgmt/httpsrv"
)

type AllConfig struct {
	Debug  bool
	Server Server `mapstructure:",squash"`
}

type Server struct {
	HTTP httpsrv.Config `mapstructure:",squash"`
}

func ReadConfig() (c *AllConfig, err error) {
	c = &AllConfig{}
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		// parse config
		if err := viper.Unmarshal(c); err != nil {
			return c, err
		}
		return c, nil
	}

	c.Debug = viper.GetBool("DEBUG")
	c.Server.HTTP.Addr = viper.GetString("HTTP_ADDR")
	return
}
