package database

import (
	"context"
	"os"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

func connectToDatabase() (*pgx.Conn, error){
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/filestorage")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	return conn, nil
}

func addUserToDatabase(conn pgx.Conn, user User) error{
	userInDB, err1 := UserInDatabase(conn, user)
	if (err1 != nil) {
		return err1
	}
	if userInDB == 1 {
		log.Println("This user is already in Database!")
		return nil
	}
	query := "INSERT INTO users (login, password) VALUES (@login, @password)"
	args := pgx.NamedArgs{
        "login": user.Login,
        "password": user.Password,
    }
	_, err := conn.Exec(context.Background(), query, args)
    if err != nil {
        log.Println("Error Inserting")
		return err
    }
	return nil
}

func UserInDatabase(conn pgx.Conn, user User) (int, error) {
	query := "SELECT * FROM users"
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
        log.Printf("Error Querying")
        return 0, err
    }
	defer rows.Close()

	for rows.Next() {
		var curUser User
		var a int
		err := rows.Scan(&a, &curUser.Login, &curUser.Password)
        if err != nil {
            log.Printf("Error Fetching")
            return 1, err
        }
		if (curUser.Login == user.Login) {
			return 1, nil
		}
	}
	return 0, nil
}