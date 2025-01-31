package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
	Tags  string `json:"tags"`
}

type JsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ResponseJSON(c *gin.Context, status int, message string, data any) {
	response := JsonResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
	c.JSON(status, response)
}
