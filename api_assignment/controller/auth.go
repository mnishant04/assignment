package controller

import (
	"api_assignment/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Session struct {
	Type            string
	ID              uint
	LoginActivityID *uint
	Expiry          time.Time
}

var sessions = map[string]Session{}

func Login(ctx *gin.Context) {
	var login models.LoginRequest
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "failed", Data: "Wrong Username or Password"})
		return
	}
	id, err := DB.Login(&login)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "failed", Data: err.Error()})
		return
	}
	var activityID *uint
	if login.Type == "user" {
		activityID, err = DB.TrackActivity(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "failed", Data: err.Error()})
			return
		}
	}
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(10 * time.Minute)

	sessions[sessionToken] = Session{
		ID:              *id,
		Type:            login.Type,
		Expiry:          expiresAt,
		LoginActivityID: activityID,
	}

	ctx.SetCookie("session_token", sessionToken, 600, "/", "localhost", false, false)
	ctx.JSON(http.StatusOK, models.Response{Error: true, Message: "success", Data: login})
}

func getSession(ctx *gin.Context) (Session, error) {
	sessionToken, err := ctx.Cookie("session_token")
	if err != nil {
		return Session{}, err
	}
	value, ok := sessions[sessionToken]
	if !ok {
		return Session{}, err
	}
	return value, nil
}

func Logout(ctx *gin.Context) {
	session, err := getSession(ctx)
	if err != nil || session.Type != "user" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Error: false, Message: "Unauthorized", Data: "Login again session expired"})
		return
	}
	err = DB.LogoutActivity(&session.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "failed", Data: "Logout "})
		return
	}
	session_token, err := ctx.Cookie("session_token")
	delete(sessions, session_token)
	ctx.JSON(http.StatusOK, models.Response{Error: false, Message: "logged Out", Data: ""})
}
