package main

import (
	"example/testapi1/services"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	handom := rand.NewSource(time.Now().UnixNano())
	name := services.Sktnamegen(handom)

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": name,
		})
	})
	server.Run(":8080")
}
