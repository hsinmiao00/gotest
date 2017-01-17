package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	router.POST("/user/", func(c *gin.Context) {
		//message := c.PostForm("message")
		x, _ := ioutil.ReadAll(c.Request.Body)
		defer c.Request.Body.Close()
		fmt.Println(string(x))
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": string(x),
		})
	})
	router.Run(":8080")
}
