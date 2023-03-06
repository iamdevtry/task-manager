package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBUsername string `mapstructure:"DB_Username"`
	DBPassword string `mapstructure:"DB_Password"`
	DBServer   string `mapstructure:"DB_Server"`
	DBService  string `mapstructure:"DB_Service"`
	// ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
