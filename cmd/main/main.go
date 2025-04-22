package main

import (
	"github.com/kgugunava/file-storage/internal/app"
	"github.com/kgugunava/file-storage/internal/database"
)

func main() {
	Db := new(database.Database)
	app.RunApp(*Db)
}