package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//handom := rand.NewSource(time.Now().UnixNano())
	//name := services.Sktnamegen(handom)
	time := time.Now().UnixNano()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": time,
		})
	})
	server.Run(":8080")
}
