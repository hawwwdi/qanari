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
	r.Use(getCORSMiddlware())
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
		msg.POST("/receive", getMessagesFromAUserHandler(db))
		msg.GET("/list", getMessageSendersHandler(db))
	}
	ava := r.Group("/ava")
	{
		ava.POST("/postAVA", getPostAvaHandler(db))
		ava.POST("/postComment", getPostCommentHandler(db))
		ava.GET("/timeline", getTimeLineHandler(db))
		ava.POST("/comments", getAvasCommentHandler(db))
		ava.GET("/me", getPersonalAvasHandler(db))
		ava.POST("/user", getUserAvasHandler(db))
		ava.POST("/tag", getAvasByTagHandler(db))
		ava.GET("/popular", getAvasByLikesHandler(db))
	}
	like := ava.Group("/like")
	{
		like.POST("/doLike", getLikeHandler(db))
		like.POST("/list", getLikersHandler(db))
		like.GET("/count", getLikesCountHandler(db))
	}
}

func getCORSMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, content-type, method, x-pingother")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, struct{}{})
			return
		}
		// buf := new(bytes.Buffer)
		// buf.ReadFrom(c.Request.Body)
		// fmt.Println(buf.String())
		c.Next()
	}
}
