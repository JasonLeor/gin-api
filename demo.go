package main

import (
	"gindemo/apis"
	"gindemo/db"
	"github.com/gin-gonic/gin"
)

func main() {
	defer db.SqlDB.Close()
	router := gin.Default()

	router.GET("/index", apis.IndexHandler)
	router.POST("/person", apis.NewPersonHandler)

	router.Run(":8080")
}
