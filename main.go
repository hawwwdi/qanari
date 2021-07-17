package main

import (
	"log"

	"net/http"
	_ "net/http/pprof"

	"github.com/hawwwdi/qanari/db"
	"github.com/hawwwdi/qanari/handlers"
)

func main() {
	db_user := "root"
	db_password := "12345678"
	db_name := "qanari"
	db := db.NewDB(db_user, db_password, db_name)
	defer db.Close()
	app := handlers.NewApp(db)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	app.Start()
}
