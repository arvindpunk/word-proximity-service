package utils

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Environment struct {
	DBWordProximity string `required:"true" split_words:"true"`
}

var Env Environment

func LoadEnvironment() {
	err := envconfig.Process("", &Env)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("missing environment variables")
	}
}
