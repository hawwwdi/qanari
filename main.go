package main

import (
	"github.com/hawwwdi/qanari/db"
	"github.com/hawwwdi/qanari/handlers"
)

func main() {
	db := db.NewDB("tss", "tss", "qanari")
	app := handlers.NewApp(db)
	app.Start()
}
