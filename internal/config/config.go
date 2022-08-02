package config

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"strings"
)

type Config struct {
	LogLevel    string `mapstructure:"log_level"`
	HttpAddress string `mapstructure:"http_address"`
	Deployment  string `mapstructure:"deployment"`

	appName     string
	defaultConf []byte
}

const defaultConf = `
log_level: "debug"
http_port: ":8080"
deployment: "production"
`

func getDefault() *Config {
	return &Config{
		LogLevel:    "debug",
		HttpAddress: "localhost:8080",
		Deployment:  "development",
		appName:     "ushort",
	}
}

func (c *Config) configureViper() {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()                                        // read in environment variables that match
	viper.SetEnvPrefix(strings.ReplaceAll(c.appName, "-", "_")) // will be uppercase automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AddConfigPath("/etc/" + c.appName + "/")
	viper.AddConfigPath("$HOME/." + c.appName + "/")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
}

func (c *Config) readConfFromFile(confPath string) error {
	if confPath != "" {
		content, err := ioutil.ReadFile(confPath)

		if err != nil {
			return err
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return err
		}
	} else {
		// If a config file is found, read it in.
		if err := viper.MergeInConfig(); err == nil {
			log.Printf("Using config file: %s \n", viper.ConfigFileUsed())
		} else {
			return err
		}
	}
	return nil
}

// LoadConf load config from file and read in environment variables that match
func (c *Config) loadConf(configPath string) error {
	c.configureViper()

	if err := viper.ReadConfig(bytes.NewBuffer(c.defaultConf)); err != nil {
		return err
	}
	viper.AddConfigPath(configPath)
	if err := viper.MergeInConfig(); err == nil {
		log.Printf("Using config file: %s \n", viper.ConfigFileUsed())
	} else {
		return err
	}
	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Println("unable to decode into config struct, ", err)
		return err
	}

	return nil
}

func New(appName string, configPath string) *Config {
	configs := Config{appName: appName, defaultConf: []byte(defaultConf)}
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := configs.loadConf(configPath)
	if err != nil {
		log.Fatalf("error reading config %v", err)
	}
	return &configs

}
