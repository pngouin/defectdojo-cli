package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ProductName    string `mapstructure:"DD_PRODUCTNAME"`
	EngagementName string `mapstructure:"DD_ENGAGEMENTNAME"`
	ApiKey         string `mapstructure:"DD_APIKEY"`
	Host           string `mapstructure:"DD_HOST"`
}

var Configuration Config

func LoadConfigFromEnv() {
	viper.AllowEmptyEnv(true)
	viper.SetDefault("default", "")
	viper.SetTypeByDefaultValue(true)
	viper.SetEnvPrefix("dd")
	viper.AutomaticEnv()

	Configuration = Config{
		ProductName: viper.GetString("productname"),
		EngagementName: viper.GetString("engagementname"),
		ApiKey: viper.GetString("apikey"),
		Host: viper.GetString("host"),
	}
}
