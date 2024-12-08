package configs

import (
	"github.com/spf13/viper"
)

var cfg Config

type Config struct {
	BrasilApiUrl string
	ViaCepApiUrl string
	HttpPort     string
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.SetConfigFile("env.yaml")
	viper.AutomaticEnv() // load environment variables
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg, nil
}
