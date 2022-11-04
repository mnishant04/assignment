package controller

import (
	"api_assignment/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := isValid(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Error: false, Message: "Not Authorized", Data: ""})
			return
		}
		c.Next()
	}
}

func isValid(ctx *gin.Context) error {
	cookie, err := ctx.Cookie("session_token")
	if err != nil {
		return err
	}
	value, ok := sessions[cookie]
	if !ok || value.isExpired() {
		return err
	}

	return nil
}

func (s Session) isExpired() bool {
	return s.Expiry.Before(time.Now())
}
