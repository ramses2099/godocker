package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppPort       string `mapstructure:"PORT"`
	RuntimeSetup  string `mapstructure:"RUNTIME_SETUP"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	// viper.AddConfigPath(path)
	// viper.SetConfigName(".env")
	// viper.SetConfigType("env")

	viper.AutomaticEnv()

	// err = viper.ReadInConfig()
	// if err != nil {
	// 	return Config{}, err
	// }

	// err = viper.Unmarshal(&config)
	return
}
