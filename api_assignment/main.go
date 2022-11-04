package main

import (
	_ "api_assignment/config"
	"api_assignment/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.POST("/login", controller.Login)
	r := router.Group("/api/v1/admin/")
	r2:=router.Group("api/v1/user/")
	r.Use(controller.AuthMiddleware())
	r2.Use(controller.AuthMiddleware())
	r.POST("addUser", controller.AddUser)
	r.GET("getUsers", controller.GetUsers)
	r.GET("getUserProducts", controller.GetUserProducts)
	r.GET("trackActivities",controller.TrackUserActivities)
	r2.GET("logout",controller.Logout)
	r2.GET("getProducts",controller.ProductList)
	r2.PUT("updateProduct",controller.UpdateProduct)
	r2.POST("addProduct",controller.AddProd)
	r2.DELETE("deleteProduct/:id",controller.DeleteProduct)
	r2.POST("uploadProductImage",controller.UploadImage)
	err := router.Run(":8080")
	if err != nil {
		log.Printf("Error while starting server %s\n", err)
	}
}
