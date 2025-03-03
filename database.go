package main

import (
	"context"
	"os"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func connectToDatabase() {
	// dbUrl := "postgres://postgres:postgres@localhost:5432/filestorage"

	conn, err := pgx.Connect(context.Background(), os.Getenv("postgres://postgres:postgres@localhost:5432/filestorage"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

}