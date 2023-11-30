package routes

import (
	"github.com/gin-gonic/gin"
	c "gin/restAPI/controllers"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/getUser/:id", c.GetUsers)
	router.GET("/listUser", c.ListUsers)
	router.PUT("/editUser", c.EditUsers)
	router.POST("/addUser", c.AddUsers)
}