package handlers

import (
	"errors"
	"time"

	dbWordProximity "github.com/arvindpunk/word-proximity-service/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ResponseTargetWord struct {
	TargetWord string `json:"targetWord"`
}

var cache = make(map[int]string)

func GetTargetWord() gin.HandlerFunc {
	return func(c *gin.Context) {
		targetWord, err := getWordForTime(time.Now().UTC())
		if err != nil {
			log.Error().
				Err(err).
				Str("handler", c.FullPath()).
				Send()
			WithError("ERR_CACHE_001", err.Error()).
				Respond(c)
			return
		}
		res := ResponseTargetWord{
			TargetWord: targetWord,
		}

		WithSuccess(res).
			Respond(c)
	}
}

// YYYY0000 +
// 0000MM00 +
// 000000DD
// => year * 10000 + month * 100 + day = 20230206
func getDateIntFromTime(time time.Time) int {
	year, month, day := time.Date()
	return year*1000 + int(month) + day
}

func getWordForTime(time time.Time) (string, error) {
	dateInt := getDateIntFromTime(time)
	targetWord, ok := cache[dateInt]
	if !ok {
		return "", errors.New("failed to find word in cache")
	}
	return targetWord, nil
}

func RefreshWordCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		l := log.With().
			Str("handler", c.FullPath()).
			Logger()
		db, err := dbWordProximity.Connect()
		if err != nil {
			l.Error().
				Err(err).
				Msg("db connection failed")
			WithError("ERR_DB_001", err.Error()).
				Respond(c)
			return
		}
		defer db.Close()

		now := time.Now().UTC()

		rows, err := db.QueryContext(ctx,
			`SELECT 
				id, word, date, created_at, updated_at
			FROM words 
			WHERE date > $1 AND date <= $2`,
			now.AddDate(0, 0, -15),
			now.AddDate(0, 0, 15))

		if err != nil {
			l.Error().
				Err(err).
				Msg("failed to get words from db")
			WithError("ERR_DB_002", err.Error()).
				Respond(c)
			return
		}
		defer rows.Close()

		words := []dbWordProximity.Word{}

		for rows.Next() {
			var word dbWordProximity.Word
			err := rows.Scan(&word.Id, &word.Word, &word.Date, &word.CreatedAt, &word.UpdatedAt)
			if err != nil {
				l.Error().
					Err(err).
					Msg("failed while scanning row")
				WithError("ERR_DB_003", err.Error()).
					Respond(c)
				return
			}
			words = append(words, word)
		}

		// TODO: update cache with thread safety

		WithSuccess(words).
			Respond(c)
	}
}
