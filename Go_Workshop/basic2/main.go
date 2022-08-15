package main

import (
	"net/http"

	"github.com/basic2/handle"
	"github.com/basic2/model"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func main() {
	color.Blue("Hello World")
	r := gin.Default()
	m := handle.NewMember()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/member", func(c *gin.Context) {
		c.JSON(200, m.AllData())
	})

	r.POST("/member", func(c *gin.Context) {
		var v model.Member
		err := c.ShouldBindJSON(&v)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Error Format",
			})
			return
		}

		m.AddData(v)
		c.JSON(200, v)
	})

	r.Run(":8011") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
