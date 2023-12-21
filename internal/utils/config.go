package utils

import "github.com/spf13/viper"

type Config struct {
	Port       int
	AuthServer string
	Verbosity  int
}

func NewConfig() Config {
	viper.SetDefault("FF_PORT", 8080)
	viper.SetDefault("FF_AUTHSERVER", "http://localhost:9000")
	viper.SetDefault("FF_VERBOSITY", 1)

	viper.AutomaticEnv()

	return Config{
		Port:       viper.GetInt("FF_PORT"),
		AuthServer: viper.GetString("FF_AUTHSERVER"),
		Verbosity:  viper.GetInt("FF_VERBOSITY"),
	}
}
