package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawwwdi/qanari/db"
	"github.com/hawwwdi/qanari/model"
)

func getSingUpHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form model.SignupCredential
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Signup(form); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
}

func getLoginHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form model.LoginCredential
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Login(form); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successfully",
		})
		return
	}
}

func getFollowHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Follow(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("user %s followed successfully", user.Username),
		})
		return
	}
}

func getUnfollowHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.UnFollow(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("user %s Unfollowed successfully", user.Username),
		})
		return
	}
}

func getBlockHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Block(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("user %s blocked successfully", user.Username),
		})
		return
	}
}

func getUnblockHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.UnBlock(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("user %s unblocked successfully", user.Username),
		})
		return
	}
}
