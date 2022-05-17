package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "github.com/zqb-knight/applescript/utils"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("server start")

		lat, ok := c.GetQuery("latitude")
		fmt.Println("latitude: " + lat)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"data": utils.BuildResponse(-1, "lack of latitude"),
			})

		}
		long, ok := c.GetQuery("longitude")
		fmt.Println("longitude: " + long)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"data": utils.BuildResponse(-1, "lack of longitude"),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"data": utils.GetDetail(lat, long),
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
