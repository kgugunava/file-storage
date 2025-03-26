package app

import (
	"file-storage/internal/database"
	"file-storage/internal/server"
)

func RunApp() {
	database.ConnectToDatabase()
	server.RunServer()
}