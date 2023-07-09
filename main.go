package main

import (
	"example/testapi1/services"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Use(gin.Logger())

	server.GET("/test", func(ctx *gin.Context) {
		handom := rand.NewSource(time.Now().UnixNano())
		name := services.Sktnamegen(handom)
		//time := time.Now().UnixNano()
		ctx.JSON(200, gin.H{
			"message": name,
		})
	})

	server.GET("/test10", func(ctx *gin.Context) {
		// how to put an array of strings into a json
		names := services.Sktnamegen10()
		namemap := make(map[string]string)

		//time := time.Now().UnixNano()
		for i := 0; i < 10; i++ {
			jkey := fmt.Sprintf("Name no 0%d", i+1)
			if i == 9 {
				jkey = fmt.Sprintf("Name no %d", i+1)
			}
			sname := namemap[jkey]
			if sname == "" {
				sname = names[i]
			}
			namemap[jkey] = sname
		}
		ctx.JSON(200, namemap)
	})
	server.Run(":8080")
}
