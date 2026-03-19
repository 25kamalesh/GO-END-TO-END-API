package handlers

import (
	"github.com/25Kamalesh/go_todo_api/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)


type TodoHandler struct {
	Pool *pgxpool.Pool
}

type CreateTodoInput struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

func (h *TodoHandler) CreateTodoHandler(c *gin.Context) {
	var input *CreateTodoInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400 , gin.H {"error" : err.Error()})
		return
	}
todo ,err := repository.CreateTodo(c.Request.Context(),h.Pool, input.Title, input.Completed)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, todo)
}