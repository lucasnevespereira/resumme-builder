package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"resumme-builder/internal/utils/logger"
)

type AppConfig struct {
	AppName      string
	Port         int
	Version      string
	DeeplAuthKey string
	errors       []error
}

func LoadAppConfig() AppConfig {
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	conf := AppConfig{}
	conf.AppName = "Resumme Builder"
	conf.Port = conf.getIntWithDefault("PORT", 9000)
	conf.Version = "0.0.0"
	conf.DeeplAuthKey = conf.getStringWithDefault("DEEPL_AUTH_KEY", "")

	if len(conf.errors) != 0 {
		logger.Log.Fatalf("LoadAppConfig errors: %v \n", conf.errors)
	}

	return conf
}

func (c *AppConfig) getString(key string) (value string) {
	if value = viper.GetString(key); value == "" {
		c.errors = append(c.errors, fmt.Errorf("cannot find conf key for %s", key))
	}
	return value
}

func (c *AppConfig) getStringWithDefault(key, defaultValue string) string {
	viper.SetDefault(key, defaultValue)
	return viper.GetString(key)
}

func (c *AppConfig) getIntWithDefault(key string, defaultValue int) int {
	viper.SetDefault(key, defaultValue)
	return viper.GetInt(key)
}

func (c *AppConfig) getBoolWithDefault(key string, defaultValue bool) bool {
	viper.SetDefault(key, defaultValue)
	return viper.GetBool(key)
}
