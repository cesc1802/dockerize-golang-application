package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateTodoRequest struct {
	Description string `form:"description"`
}
type Todo struct {
	ID          int        `gorm:"primaryKey"`
	Description string     `gorm:"column:description"`
	CreatedAt   *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func pingHandler(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong!!!",
		})
	})
}

func createTodoHandler(router *gin.Engine, db *gorm.DB) {
	router.POST("/todos", func(c *gin.Context) {
		var request CreateTodoRequest
		if err := c.ShouldBind(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "cannot read request body",
			})
			return
		}

		var todo Todo
		todo.Description = request.Description

		if err := db.Create(&todo).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "cannot create a new todo",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":          todo.ID,
			"description": todo.Description,
			"created_at":  todo.CreatedAt,
			"updated_at":  todo.UpdatedAt,
		})
		return
	})
}

func getTodoByIdHandler(router *gin.Engine, db *gorm.DB) {
	router.GET("/todos/:id", func(c *gin.Context) {
		var id = c.Param("id")
		var todo Todo
		if err := db.Table(todo.TableName()).Where("id = ?", id).First(&todo).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "cannot read request body",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":          todo.ID,
			"description": todo.Description,
			"created_at":  todo.CreatedAt,
			"updated_at":  todo.UpdatedAt,
		})
		return
	})
}
func TodoHandlers(router *gin.Engine, db *gorm.DB) {
	pingHandler(router)
	createTodoHandler(router, db)
	getTodoByIdHandler(router, db)

}
