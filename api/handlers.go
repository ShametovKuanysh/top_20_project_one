package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	if err := DB.AutoMigrate(&Article{}); err != nil {
		log.Fatal("Error migrating the database", err)
	}
}

func CreateArticle(c *gin.Context) {
	var article Article

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
			"message": "Invalid request payload",
		})
		return
	}

	DB.Create(&article)
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Article created successfully",
		"data":    article,
	})
}

func GetArticles(c *gin.Context) {
	var articles []Article
	log.Println("Getting articles")

	DB.Find(&articles)
	ResponseJSON(c, http.StatusOK, "Articles fetched successfully", articles)

}

func GetArticle(c *gin.Context) {
	var article Article
	id := c.Param("id")

	if err := DB.First(&article, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"error":   err,
			"message": "Article not found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Article fetched successfully",
		"data":    article,
	})
}

func UpdateArticle(c *gin.Context) {
	var article Article
	id := c.Param("id")

	if err := DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"error":   err,
			"message": "Article not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
			"message": "Invalid request payload",
		})
		return
	}

	DB.Save(&article)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Article updated successfully",
		"data":    article,
	})
}

func DeleteArticle(c *gin.Context) {
	var article Article
	id := c.Param("id")

	if err := DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"error":   err,
			"message": "Article not found",
		})
		return
	}

	if err := DB.Delete(&article, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"error":   err.Error,
			"message": "Error deleting the article",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Article deleted successfully",
	})
}
