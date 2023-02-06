package handlers

import (
	"errors"
	"time"

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
				Send()
			WithError("ERR_DB_001", err.Error()).
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
func getWordForTime(time time.Time) (string, error) {
	time.Month()
	year, month, day := time.Date()
	dateInt := year*1000 + int(month) + day
	targetWord, ok := cache[dateInt]
	if !ok {
		return "", errors.New("failed to find word in cache")
	}
	return targetWord, nil
}
