package options

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type ProxyConfig struct {
	Http struct {
		Host      string `mapstructure:"host"`
		Port      int    `mapstructure:"port"`
		AccessKey string `mapstructure:"access_key"`
	}
	Chat struct {
		BaseUrl string `mapstructure:"base_url"`
		APIKey  string `mapstructure:"api_key"`
	}
}

var pc *ProxyConfig

func InitProxyConfig() {
	newCommandArgs()
	configPath := Args.Config
	if configPath == "" {
		log.Fatal("config file is required")
	}
	// check file exist
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Fatalf("config file %s not found", configPath)
	}
	// read config file
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(configPath)
	v.ReadInConfig()
	err = v.Unmarshal(&pc)
	if err != nil {
		log.Fatalf("unable to read the config file, %v", err)
	}
}

func GetConfig() *ProxyConfig {
	return pc
}
