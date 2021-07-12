package main

import (
	"github.com/hawwwdi/qanari/db"
	"github.com/hawwwdi/qanari/handlers"
)

func main() {
	db := db.NewDB("root", "12345678", "qanari")
	app := handlers.NewApp(db)
	app.Start()
}
