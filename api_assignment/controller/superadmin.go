package controller

import (
	"api_assignment/config"
	"api_assignment/models"
	"api_assignment/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var DB service.Database

func init() {
	DB = service.Database{
		Conn: config.DB,
	}
}

func AddUser(ctx *gin.Context) {
	var request models.User
	session, err := getSession(ctx)
	if err != nil || session.Type != "superadmin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Error: false, Message: "Unauthorized", Data: "Login again session expired"})
		return
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "failed", Data: err})
		return
	}
	err = DB.AddUser(&request)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "User Already Registered", Data: "Login or Use different Email"})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "failed", Data: "Error while adding User please try after some time"})
		return
	}
	ctx.JSON(http.StatusCreated, models.Response{Error: true, Message: "Registered", Data: request})
}

func GetUsers(ctx *gin.Context) {
	session, err := getSession(ctx)
	if err != nil || session.Type != "superadmin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Error: false, Message: "Unauthorized", Data: "Login again session expired"})
		return
	}
	users, err := DB.ListUser()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "failed", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{Error: true, Message: "success", Data: users})

}

func GetUserProducts(ctx *gin.Context) {
	session, err := getSession(ctx)
	if err != nil || session.Type != "superadmin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Error: false, Message: "Unauthorized", Data: "Login again session expired"})
		return
	}
	users, err := DB.ListUserProducts()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "failed", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{Error: true, Message: "success", Data: users})

}

func TrackUserActivities(ctx *gin.Context) {
	session, err := getSession(ctx)
	if err != nil || session.Type != "superadmin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Error: false, Message: "Unauthorized", Data: "Login again session expired"})
		return
	}
	users, err := DB.TrackUserActivity()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "failed", Data: err})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{Error: true, Message: "success", Data: users})

}
