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
		if err := db.PostComment(comment); err != nil {
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

func getTimeLineHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		avas, err := db.GetTimeLine()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"timeLine": avas,
		})
		return
	}
}

func getAvasCommentHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ava model.Ava
		if err := c.ShouldBindJSON(&ava); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		comments, err := db.GetComments(ava)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"comments": comments,
		})
		return
	}
}

func getPersonalAvasHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		avas, err := db.GetPersonalAvas()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"avas": avas,
		})
		return
	}
}

func getUserAvasHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		avas, err := db.GetUserAvas(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"avas": avas,
		})
		return
	}
}

func getAvasByTagHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tag model.Tag
		if err := c.ShouldBindJSON(&tag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		avas, err := db.GetAvasByTags(tag)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"avas": avas,
		})
		return
	}
}

func getAvasByLikesHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		avas, err := db.GetAvasByLikes()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"avas": avas,
		})
		return
	}
}
