package httpServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
func Server2(){

	r := gin.Default()

	r.POST("ping", func(context *gin.Context) {
		json := make(map[string]interface{}) //注意该结构接受的内容
		context.BindJSON(&json)
		fmt.Println(json)
		context.JSON(200, gin.H{
			"message": "pong",
			"addr" :"server2",
		})
	})


	r.Run(":9001")
}

