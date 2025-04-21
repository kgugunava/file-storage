package database

import (
	"context"
	"file-storage/internal/models"
	"fmt"
	"log"
	"os"
	"github.com/jackc/pgx/v5"
	"github.com/doug-martin/goqu/v9"
)

type Database struct {
	Connection *pgx.Conn
}

func (db *Database) ConnectToDatabase() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/filestorage")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	db.Connection = conn
	return conn, nil
}

func (db Database) AddUserToDatabase(user models.User) error {
	userInDB, err1 := db.IsUserInDatabase(user)
	if err1 != nil {
		return err1
	}
	if userInDB == 1 {
		log.Println("This user is already in Database!")
		return nil
	}
	query, _, _:= goqu.Insert("users").Cols("login", "password").Vals(goqu.Vals{user.Login, user.Password}).ToSQL()
	_, err := db.Connection.Exec(context.Background(), query)
	if err != nil {
		log.Println("Error Inserting")
		return err
	}
	return nil
}

func (db Database) IsUserInDatabase(user models.User) (int, error) {
	query, _, _ := goqu.From("users").ToSQL()
	rows, err := db.Connection.Query(context.Background(), query)
	if err != nil {
		log.Printf("Error Querying")
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var curUser models.User
		var a int
		err := rows.Scan(&a, &curUser.Login, &curUser.Password)
		if err != nil {
			log.Printf("Error Fetching")
			return 1, err
		}
		if curUser.Login == user.Login {
			return 1, nil
		}
	}
	return 0, nil
}

func (db Database) CheckAuthRequest(user models.User) int {
	query, _, _ := goqu.From("users").Where(goqu.Ex{"login" : user.Login}).ToSQL()
	fmt.Print(query)
	row, err := db.Connection.Query(context.Background(), query)
	if err != nil {
		log.Printf("Error Querying")
		return 0
	}
	var checkUser models.User
	err1 := row.Scan(&checkUser.Id, &checkUser.Login, &checkUser.Password)
	fmt.Println(checkUser.Password)
	if (err1 != nil) {
		log.Printf("Error Fetching")
		return 1
	} 
	if (checkUser.Password == user.Password) {
		return 1
	}
	return 0
}
