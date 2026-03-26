package main

import (
	"log"

	"github.com/25Kamalesh/go_todo_api/internal/config"
	"github.com/25Kamalesh/go_todo_api/internal/database"
	"github.com/25Kamalesh/go_todo_api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg , err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}
	
pool , err := database.ConnectPostgres(cfg.DATABASE_URI)
if err != nil {
	log.Fatal("Failed to connect to database: ", err)
}
defer pool.Close()

	todoHandler := handlers.NewTodoHandler(pool)

	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status":  "OK",
			"Message": "SUCCESS!!",
			"database": "Connected Successfully",
		})
	})
	router.POST("/todos" , todoHandler.CreateTodoHandler)
	router.Run(":" + cfg.PORT)
}
