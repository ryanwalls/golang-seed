package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadingMultipleConfigFilesExpectsSuccess(t *testing.T) {
	assert.Equal(t, "common config value", Viper.GetString("SommeCommonConfig"), "Could not read from common config file")
	assert.Equal(t, "debug", Viper.GetString("LogLevel"), "Could not read from environment specific config file")
	assert.NotEqual(t, "", Viper.GetString("VERSION_TAG"), "Could not read from version.properties file")
}
