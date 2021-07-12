package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawwwdi/qanari/db"
	"github.com/hawwwdi/qanari/model"
)

func getLikeHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ava model.Ava
		if err := c.ShouldBindJSON(&ava); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Like(ava); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "successful",
		})
	}
}

func getLikersHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ava model.Ava
		if err := c.ShouldBindJSON(&ava); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		likers, err := db.GetLikers(ava)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		bytes, err := json.Marshal(likers)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"likers": string(bytes),
		})
	}
}

func getLikesCountHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ava model.Ava
		if err := c.ShouldBindJSON(&ava); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		likes, err := db.GetLikesCount(ava)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"likes": likes,
		})
		return
	}
}
