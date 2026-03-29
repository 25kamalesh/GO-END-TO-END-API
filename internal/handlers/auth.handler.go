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
	Pool   *pgxpool.Pool
	secret []byte
}

func NewAuthHandler(pool *pgxpool.Pool, secret []byte) *AuthHandler {
	return &AuthHandler{Pool: pool, secret: secret}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := auth.GenerateHashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	c.JSON(http.StatusCreated, res)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repository.GetUserByEmail(c.Request.Context(), h.Pool, req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "USER NOT FOUND!!"})
		return
	}
	if !auth.CompareHashPassword(user.Password, req.Password) {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}
	token, err := auth.GenerateToken(user.ID, h.secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed To generate Token"})
		return
	}

	c.SetCookie("token",
		token,
		3600*24, // 1 day
		"/",     // path
		"",      // domain
		false,   // secure (true in production HTTPS)
		true,    // httpOnly
	)
	c.JSON(http.StatusOK , gin.H{"message" : "Login Successfull" })
}
