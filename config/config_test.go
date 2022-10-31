package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverlayConfigWithEnv(t *testing.T) {
	assert := assert.New(t)
	config := Config{}

	assert.Equal("", config.MQURI)

	os.Setenv("CLIPBOARD_SYNC_MQ_URI", "http://localhost:8080")

	err := config.OverlayConfigWithEnv(true)
	assert.Nil(err)
	assert.Equal("http://localhost:8080", config.MQURI)
}

func TestGetFakeConfig(t *testing.T) {
	fc := GetFakeConfig()
	assert.NotNil(t, fc)
}
