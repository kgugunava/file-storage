package app

import (
	"github.com/kgugunava/file-storage/internal/database"
	"github.com/kgugunava/file-storage/internal/server"
)

func RunApp(db database.Database) {
	db.ConnectToDatabase()
	server.RunServer(db)
}