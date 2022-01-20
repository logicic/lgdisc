package controller

import "github.com/gin-gonic/gin"

func MakeRoute() *gin.Engine {
	router := gin.Default()
	router.POST("/discory/register", register)
	return router
}
