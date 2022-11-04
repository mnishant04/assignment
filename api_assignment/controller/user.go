package controller

import (
	"api_assignment/models"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddProd(ctx *gin.Context) {
	var request models.Product
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "failed", Data: err})
		return
	}
	session, err := getSession(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "Unauthorized", Data: "Login again session expired"})
		return
	}
	request.UserID = session.ID
	err = DB.CreateProd(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "failed", Data: "Error while adding User please try after some time"})
		return
	}
	ctx.JSON(http.StatusCreated, models.Response{Error: true, Message: "success", Data: request})
}

func ProductList(ctx *gin.Context) {
	session, err := getSession(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Error: false, Message: "Unauthorized", Data: "Login again session expired"})
		return
	}

	Productlist, err := DB.ProdList(session.ID)
	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "failed", Data: "Error while adding User please try after some time"})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{Error: true, Message: "success", Data: Productlist})
}

func DeleteProduct(ctx *gin.Context) {
	product, Ok := ctx.Params.Get("id")
	if !Ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "failed", Data: "Invalid Product ID"})
		return
	}
	productID, err := strconv.Atoi(product)
	user, err := getSession(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Error: false, Message: "Unauthorized", Data: "Login again session expired"})
		return
	}
	err = DB.DeleteList(productID, user.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "failed", Data: "Error while Deleting Product list please try after some time"})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{Error: true, Message: "Deleted", Data: productID})
}

func UpdateProduct(ctx *gin.Context) {
	var request models.Product
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "failed", Data: err})
		return
	}
	session, err := getSession(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Error: false, Message: "Unauthorized", Data: "Login again session expired"})
		return
	}
	request.UserID = session.ID
	log.Printf("sessionID %v\n", session.ID)
	err = DB.UpdateList(&request.UserID, &request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: false, Message: "failed", Data: "Error while Updating Product please try after some time"})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{Error: true, Message: "success", Data: request})
}

func UploadImage(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: true, Message: "failed", Data: "Error Extracting File Data"})
		return
	}
	_, err = getSession(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Error: true, Message: "Unauthorized", Data: "Login again session expired"})
		return
	}
	files := form.File["product_image"]
	path := "./util/image/" + files[0].Filename
	dst, err := os.Create(path)
	defer dst.Close()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{Error: true, Message: "failed", Data: "Internal Server Error"})
		return
	}
	file, err := files[0].Open()
	defer file.Close()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{Error: true, Message: "failed", Data: "Internal Server Error"})
		return
	}
	if _, err := io.Copy(dst, file); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{Error: true, Message: "failed", Data: "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Error: false, Message: "success", Data: path})
}
