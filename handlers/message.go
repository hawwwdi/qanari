package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawwwdi/qanari/db"
	"github.com/hawwwdi/qanari/model"
)

func getSendMessageHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var msg model.Message
		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.SendMessage(msg); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "message was sent successfully",
		})
		return
	}
}

func getMessagesFromAUserHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		messages, err := db.GetMessagesFromAUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"messages": messages,
		})
	}
}

func getMessageSendersHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := db.GetMessageSenders()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"senders": users,
		})
		return
	}
}
