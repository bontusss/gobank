package utils

import "github.com/spf13/viper"

type Config struct {
	DBDRIVER string `mapstructure:"DB_DRIVER"`
	DBSOURCE string `mapstructure:"DB_SOURCE"`
	PORT     string `mapstructure:"PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // can also be json

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}