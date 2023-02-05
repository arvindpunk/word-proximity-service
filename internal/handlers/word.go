package handlers

import "github.com/gin-gonic/gin"

type ResponseTargetWord struct {
	TargetWord string `json:"targetWord"`
}

func GetTargetWord() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := ResponseTargetWord{
			TargetWord: "APPLE",
		}

		WithSuccess(res).Respond(c)
	}
}
