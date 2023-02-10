package handlers

import (
	"context"
	"errors"
	"time"

	dbWordProximity "github.com/arvindpunk/word-proximity-service/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ResponseTargetWord struct {
	TargetWord string `json:"targetWord"`
}

var cache = make(map[int]string)

func HandleGetTargetWord() gin.HandlerFunc {
	return func(c *gin.Context) {
		targetWord, err := getWordForTime(time.Now().UTC())
		if err != nil {
			log.Error().
				Err(err).
				Str("handler", c.FullPath()).
				Send()
			WithError(err).
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
	return year*10000 + int(month)*100 + day
}

func getWordForTime(time time.Time) (string, error) {
	dateInt := getDateIntFromTime(time)
	targetWord, ok := cache[dateInt]
	if !ok {
		return "", errors.New("failed to find word in cache")
	}
	return targetWord, nil
}

func HandleRefreshWordCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		l := log.With().
			Str("handler", c.FullPath()).
			Logger()

		words, err := RefreshWordCache(ctx, &l)
		if err != nil {
			WithError(err).
				Respond(c)
			return
		}

		WithSuccess(words).
			Respond(c)
	}
}

func RefreshWordCache(ctx context.Context, l *zerolog.Logger) ([]dbWordProximity.Word, error) {
	words := []dbWordProximity.Word{}

	db, err := dbWordProximity.Connect()
	if err != nil {
		l.Error().
			Err(err).
			Msg("db connection failed")
		return words, err
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
		return words, err
	}
	defer rows.Close()

	for rows.Next() {
		var word dbWordProximity.Word
		err := rows.Scan(&word.Id, &word.Word, &word.Date, &word.CreatedAt, &word.UpdatedAt)
		if err != nil {
			l.Error().
				Err(err).
				Msg("failed while scanning row")
			return words, err
		}
		words = append(words, word)
	}

	// TODO: update cache with thread safety
	newCache := make(map[int]string)
	for _, el := range words {
		newCache[getDateIntFromTime(el.Date)] = el.Word
	}
	cache = newCache

	for _, word := range words {
		log.Info().
			Str("word", word.Word).
			Int("date", getDateIntFromTime(word.Date)).
			Msg("fetched and updated word cache")

	}

	return words, nil
}
