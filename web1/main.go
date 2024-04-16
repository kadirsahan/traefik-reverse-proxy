package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	router := app.Group("/")
	router.GET("/web1", Hello)
	app.Run(":8090")
	// app.Run(":8091")
}
func Hello(c *gin.Context) {
	c.Writer.Write([]byte("<h1> hello web1 </h1>"))
}
