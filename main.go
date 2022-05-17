package main

import "github.com/gin-gonic/gin"
import "github.com/zqb-knight/applescript/utils"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		lat, ok := c.GetQuery("latitude")
		if !ok {
			utils.BuildResponse(c, -1, "lack of latitude")
		}
		long, ok := c.GetQuery("longitude")
		if !ok {
			utils.BuildResponse(c, -1, "lack of longitude")
		}
		utils.GetDetail(c, lat, long)

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
