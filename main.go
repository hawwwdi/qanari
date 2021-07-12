package main

import (
	"os"

	"github.com/hawwwdi/qanari/db"
	"github.com/hawwwdi/qanari/handlers"
)

func main() {
	db_user := os.Getenv("db_user")
	db_password := os.Getenv("db_password")
	db_name := os.Getenv("db_name")
	db := db.NewDB(db_user, db_password, db_name)
	defer db.Close()
	app := handlers.NewApp(db)
	app.Start()
}
