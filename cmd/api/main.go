package main

import (
	"log"

	"github.com/25Kamalesh/go_todo_api/internal/config"
	"github.com/25Kamalesh/go_todo_api/internal/database"
	"github.com/25Kamalesh/go_todo_api/internal/handlers"
	"github.com/25Kamalesh/go_todo_api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	pool, err := database.ConnectPostgres(cfg.DATABASE_URI)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer pool.Close()

	todoHandler := handlers.NewTodoHandler(pool)
	authHandler := handlers.NewAuthHandler(pool, []byte(cfg.JWT_SECRET))

	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status":   "OK",
			"Message":  "SUCCESS!!",
			"database": "Connected Successfully",
		})
	})

	// Public routes (no authentication required)
	public := router.Group("/api/v1")
	public.POST("/register", authHandler.Register)
	public.POST("/login", authHandler.Login)

	// Protected routes (authentication required)
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware([]byte(cfg.JWT_SECRET)))
	protected.POST("/todos", todoHandler.CreateTodoHandler)

	router.Run(":" + cfg.PORT)
}
