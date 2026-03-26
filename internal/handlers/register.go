package handlers

import (
	"net/http"

	"github.com/25Kamalesh/go_todo_api/internal/auth"
	"github.com/25Kamalesh/go_todo_api/internal/models"
	"github.com/25Kamalesh/go_todo_api/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthHandler struct {
	Pool *pgxpool.Pool
}

func NewAuthHandler(pool *pgxpool.Pool) *AuthHandler {
	return &AuthHandler{Pool: pool}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	err := c.ShouldBindJSON(&req); if err != nil {
		c.JSON(http.StatusBadRequest , gin.H{"error": err.Error()})
		return
	}
	hashedPassword , err := auth.GenerateHashPassword(req.Password) ; if err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{"error": err.Error()})
		return
	}

	user, err := repository.CreateUser(
		c.Request.Context(),
		h.Pool,
		req.Name,
		req.Email,
		hashedPassword,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := models.UserResponse{
		ID: user.ID,
		Name: user.Name ,
		Email : user.Email ,

	}
	c.JSON(http.StatusCreated , res)
}