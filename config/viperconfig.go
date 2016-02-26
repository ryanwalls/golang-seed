package config

import (
	"fmt"

	"bytes"
	"github.com/spf13/viper"
	// Blank importing this to enable remote support for viper.  See viper README.
	_ "github.com/spf13/viper/remote"
	"io/ioutil"
	"os"
)

// Viper is the config object for the application.  See github.com/spf13/viper for usage details.
var Viper *config

type config struct {
	*viper.Viper
}

func init() {
	v := viper.New()
	v.SetDefault("env", "dev")
	v.BindEnv("env")
	env := v.GetString("env")
	Viper = &config{EnvironmentSpecificConfig(v, env)}
}

// EnvironmentSpecificConfig sets the configuration to match the given env.
func EnvironmentSpecificConfig(v *viper.Viper, env string) *viper.Viper {
	if v == nil {
		v = viper.New()
		v.SetDefault("env", env)
	}
	// Read common config
	v.AddConfigPath(".")
	v.AddConfigPath("../")
	v.AddConfigPath("config/")
	v.AddConfigPath("../config/")
	v.SetConfigName("config-common")
	v.SetConfigType("yml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error reading common config file: %s \n", err))
	}

	// Merge in environment specific config
	mergeInConfig(v, "config-"+env+".yml", "yml")
	// Merge in version config
	mergeInConfig(v, "version.properties", "properties")

	// Read config from consul
	// v.AddRemoteProvider("consul", "http://someconsul/endpoint", "some/path/to/config.json")
	// v.SetConfigType("yaml")
	// err = v.ReadRemoteConfig() // Find and read the config file
	// if err != nil {            // Handle errors reading the config file
	// 	panic(fmt.Errorf("Fatal error reading remote config file: %s \n", err))
	// }
	return v
}

func mergeInConfig(v *viper.Viper, configName, configType string) {
	v.SetConfigType(configType)
	configPaths := []string{configName, "../" + configName, "config/" + configName, "../config/" + configName}
	configFilePath := ""
	for _, path := range configPaths {
		if b, _ := exists(path); b {
			configFilePath = path
			continue
		}
	}
	if configFilePath == "" {
		panic(fmt.Errorf("Could not find config file: %s \n", configName))
	}
	configBytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(fmt.Errorf("Could not read config file: %s \n", err))
	}
	err = v.MergeConfig(bytes.NewBuffer(configBytes)) // Find and read the config file
	if err != nil {                                   // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
