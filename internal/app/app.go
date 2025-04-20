package app

import (
	"file-storage/internal/database"
	"file-storage/internal/server"
)

func RunApp(db database.Database) {
	db.ConnectToDatabase()
	server.RunServer(db)
}