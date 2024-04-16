package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	router := app.Group("/")
	router.GET("/", Hello)
	app.Run(":8090")
	// app.Run(":8092")
}
func Hello(c *gin.Context) {
	c.Writer.Write([]byte("<h1> hello web2 </h1>"))
}
