package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	fmt.Println("!!")
	c.JSON(200, "ok!")
}
