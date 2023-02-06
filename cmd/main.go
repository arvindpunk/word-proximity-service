package main

import (
	dbWordProximity "github.com/arvindpunk/word-proximity-service/internal/db"
	"github.com/arvindpunk/word-proximity-service/internal/handlers"
	"github.com/arvindpunk/word-proximity-service/internal/utils"
	"github.com/rs/zerolog/log"
)

func init() {
	utils.LoadEnvironment()
	dbWordProximity.Test()
}

func main() {
	r := handlers.NewRouter()
	if err := r.Run(":5001"); err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to start server")
	}
}
