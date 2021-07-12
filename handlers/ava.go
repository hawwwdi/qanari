package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawwwdi/qanari/db"
	"github.com/hawwwdi/qanari/model"
)

func getPostAvaHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ava model.Ava
		if err := c.ShouldBindJSON(&ava); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.PostAva(ava); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "post ava successfully",
		})
		return
	}
}

func getPostCommentHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var comment model.Ava
		if err := c.ShouldBindJSON(&comment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.PostAva(comment); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "post comment successfully",
		})
		return
	}
}
