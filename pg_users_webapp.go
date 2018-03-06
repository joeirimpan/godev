package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	pg_user  = "USER"
	password = "PASSWORD"
	dbname   = "DB_NAME"
)

var db *sql.DB

func main() {
	initDB()
	defer db.Close()
	http.HandleFunc("/", userHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type user struct {
	Id         int    `json:"id"`
	Age        int    `json:"age"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
}

type users struct {
	Users []user `json:"users"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	usrs := users{}
	err := queryUsers(&usrs)
	fmt.Println(usrs)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(usrs)
	fmt.Println(string(out))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

func queryUsers(usrs *users) error {
	rows, err := db.Query(`
		SELECT
			id,
			age,
			first_name,
      last_name,
      email
		FROM users`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		usr := user{}
		err = rows.Scan(
			&usr.Id,
			&usr.Age,
			&usr.First_name,
			&usr.Last_name,
			&usr.Email,
		)
		if err != nil {
			return err
		}
		usrs.Users = append(usrs.Users, usr)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func initDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, pg_user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
