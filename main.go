package main

import (
	"github.com/gin-gonic/gin"
	m "gin/restAPI/middleware"
	r "gin/restAPI/routes"
)

func main() {
	router := gin.Default()

	router.Use(m.Header())
	{
		r.UserRoutes(router)
		r.ProductRoutes(router)
	}
	
	router.Run()
}