package main

import (
	"github.com/gin-gonic/gin"
	"restframework/withginversioning/version_one"
	"restframework/withginversioning/version_two"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/detail", version_one.ViewDetail)
	}

	v2 := router.Group("/v2", version_two.ViewDetail)
	{
		v2.GET("/detail")
	}
	router.Run(":8082")

}
