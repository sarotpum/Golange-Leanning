package todo

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/workshop-jwt/auth"
	"strings"
)

type Todo struct {
	Title string `json:"text"`
	gorm.Model
}

func (Todo) TableName() string {
	return "todoslist"
}

type TodoHandler struct {
	db *gorm.DB
}

func NewTodoHandler(db *gorm.DB) *TodoHandler {
	return &TodoHandler{db: db}
}

func (t *TodoHandler) NewTask(c *gin.Context) {

	s := c.Request.Header.Get("Authorization")
	tokenString :=  strings.TrimPrefix(s, "Bearer ")

	if err := auth.Protect(tokenString); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil { //ShouldBindJSON => status other or BindJSON => status 400
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	r := t.db.Create(&todo)
	if err := r.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID": todo.Model.ID,
	})
}
