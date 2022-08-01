package env

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	ENV         string `mapstructure:"ENV"`
	AppName     string `mapstructure:"APP_NAME"`
	AppVersion  string `mapstructure:"APP_VERSION"`
	ServiceURL  string `mapstructure:"SERVICE_URL"`
	ServicePort string `mapstructure:"SERVICE_PORT"`
}

var (
	cfg  *Config
	once sync.Once
)

func Get() *Config {
	if strings.HasSuffix(os.Args[0], ".test") || flag.Lookup("test.v") != nil {
		return &Config{}
	}

	return cfg
}

func Load() {
	once.Do(func() {
		v := viper.New()
		v.AutomaticEnv()

		v.AddConfigPath(".")
		v.SetConfigType("env")
		v.SetConfigName(".env")

		err := v.ReadInConfig()
		if err != nil {
			fmt.Println("config file not found: ", err)
		}

		config := new(Config)
		err = v.Unmarshal(config)
		if err != nil {
			panic(err)
		}

		cfg = config
	})
}
