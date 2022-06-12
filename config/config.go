package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Token      string `envconfig:"BOT_TOKEN"`
	QnaChannel string `envconfig:"QNA_CHANNEL_ID" default:"985223712565501992"`
}

func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
