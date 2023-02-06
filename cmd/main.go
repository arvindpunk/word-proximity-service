package main

import (
	"fmt"

	"github.com/arvindpunk/word-proximity-service/internal/handlers"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Ligma")
	r := handlers.NewRouter()
	if err := r.Run(":5001"); err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to start server")
	}
}
