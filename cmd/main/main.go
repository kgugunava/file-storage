package main

import (
	"file-storage/internal/app"
	"file-storage/internal/database"
)

func main() {
	Db := new(database.Database)
	app.RunApp(*Db)
}