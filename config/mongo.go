package config

import (
	"github.com/spf13/viper"
)

func ReadConfig() {
	// Get file config
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
