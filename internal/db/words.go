package db

import (
	"database/sql"
	"time"

	"github.com/arvindpunk/word-proximity-service/internal/utils"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type Word struct {
	Id        int64     `json:"id"`
	Word      string    `json:"word"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Connect() (*sql.DB, error) {
	connStr := utils.Env.DBWordProximity
	db, err := sql.Open("postgres", connStr)
	return db, err
}

func Test() {
	db, err := Connect()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to establish connection with database")
		return
	}
	log.Info().
		Msg("connection to database established succesfully, now disconnecting")
	db.Close()
}
