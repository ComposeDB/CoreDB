package utilities

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Port       int
	Datafolder string
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetDefault("Datafolder", filepath.Join(GetHomeDir(), ".autoai", "coredb", "data"))
	viper.SetConfigName("coredb")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Configuration Not Found...")
	}
	err = viper.Unmarshal(&config)
	return
}
