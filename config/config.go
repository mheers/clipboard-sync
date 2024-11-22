package config

import (
	"context"
	"os"
	"time"

	"dario.cat/mergo"
	"github.com/brianvoe/gofakeit/v5"
	"github.com/mheers/clipboard-sync/helpers"
	"github.com/sethvargo/go-envconfig"
	"github.com/sirupsen/logrus"
)

// Config describes the config
type Config struct {
	MQJWT   string `env:"CLIPBOARD_SYNC_MQ_JWT"`
	MQURI   string `env:"CLIPBOARD_SYNC_MQ_URI"`
	MQUSeed string `env:"CLIPBOARD_SYNC_MQ_USEED"`
}

// OverlayConfigWithEnv overlays the config with values from the env
func (cfg *Config) OverlayConfigWithEnv(print bool) error {
	ctx := context.Background()
	overlayCfg := &Config{}
	err := envconfig.Process(ctx, overlayCfg)
	if err != nil {
		return err
	}

	err = mergo.Merge(cfg, overlayCfg, mergo.WithOverride)
	if err != nil {
		return err
	}
	return nil
}

// GetFakeConfig creates a config for testing purposes only
func GetFakeConfig() *Config {
	gofakeit.Seed(time.Now().UTC().UnixNano())

	cfg := &Config{}

	cfg.OverlayConfigWithEnv(false)

	logLevel := os.Getenv("LOGLEVEL")
	if logLevel != "" {
		helpers.SetLogLevel(logLevel)
	}

	return cfg
}

var configInstance *Config

func GetConfig(print bool) *Config {
	if configInstance == nil {
		configInstance = &Config{}
		err := configInstance.OverlayConfigWithEnv(false)
		if err != nil {
			logrus.Fatal(err)
		}
	}
	return configInstance
}
