package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawwwdi/qanari/db"
)

type App struct {
	route *gin.Engine
}

func NewApp(db *db.DB) *App {
	r := gin.Default()
	initRouter(r, db)
	return &App{
		route: r,
	}
}

func (a *App) Start(addr ...string) {
	a.route.Run(addr...)
}

func initRouter(r *gin.Engine, db *db.DB) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	user := r.Group("/user")
	{
		user.POST("/signup", getSingUpHandler(db))
		user.POST("/login", getLoginHandler(db))
		user.POST("/follow", getFollowHandler(db))
		user.POST("/unfollow", getUnfollowHandler(db))
		user.POST("/block", getBlockHandler(db))
		user.POST("/unblock", getUnblockHandler(db))
	}
	msg := user.Group("/message")
	{
		msg.POST("/send", getSendMessageHandler(db))
		msg.GET("/receive", getMessagesFromAUserHandler(db))
		msg.GET("/list", getMessageSendersHandler(db))
	}
	ava := r.Group("/ava")
	{
		ava.POST("/postAVA", getPostAvaHandler(db))
		ava.POST("/postComment", getPostCommentHandler(db))
		ava.GET("/timeline", getTimeLineHandler(db))
		ava.GET("/comments", getAvasCommentHandler(db))
		ava.GET("/me", getPersonalAvasHandler(db))
		ava.GET("/user", getUserAvasHandler(db))
		ava.GET("/tag", getAvasByTagHandler(db))
		ava.GET("/popular", getAvasByLikesHandler(db))
	}
	like := ava.Group("/like")
	{
		like.POST("/doLike", getLikeHandler(db))
		like.GET("/list", getLikersHandler(db))
		like.GET("/count", getLikesCountHandler(db))
	}
}
