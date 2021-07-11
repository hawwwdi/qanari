package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawwwdi/qanari/db"
)

type App struct {
	route *gin.Engine
	db    *db.DB
}

func NewApp(db *db.DB) *App {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return &App{
		route: r,
		db:    db,
	}
}

func (a *App) Start() {
	a.route.Run()
}