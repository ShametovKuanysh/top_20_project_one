package main

import (
	"github.com/ShametovKuanysh/top_20_project_one/api"
	"github.com/gin-gonic/gin"
)

func main() {
	api.InitDB()

	router := gin.Default()

	router.GET("/articles", api.GetArticles)
	router.GET("/article/:id", api.GetArticle)
	router.POST("/article", api.CreateArticle)
	router.PUT("/article/:id", api.UpdateArticle)
	router.DELETE("/article/:id", api.DeleteArticle)

	router.Run(":8080")
}
