package handlers

import (
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
		}
		if err := db.Signup(form); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}

func getLoginHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form model.LoginCredential
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		if err := db.Login(form); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successfuly",
		})
	}
}
