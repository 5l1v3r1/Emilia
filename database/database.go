package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	// password = ""
	dbname = "postgres"
)

var db *sql.DB

func init() {
	fmt.Println("Database init")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, os.Getenv("pqDev"), dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sqlStatement := `CREATE TABLE IF NOT EXISTS users (
		userid text, 
		username text,
		xp int,
		level int,
		PRIMARY KEY (userid)  
	  )`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

}

// CreateUser ...
func CreateUser(userid, username string) {
	sqlStatement := `
	INSERT INTO users 
	(userid, username, xp, level) 
	VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(sqlStatement, userid, username, 0, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created user %v \n", userid)
}

// UpdateUserXP ...
func UpdateUserXP(userid string) {
	sqlStatement := `
	UPDATE users
	SET xp = xp + 1
	WHERE userid = $1;`
	_, err := db.Exec(sqlStatement, userid)
	if err != nil {
		panic(err)
	}
	fmt.Println("Update user")
}

// CheckUser ...
func CheckUser(userid, username string) {
	var id string
	result := db.QueryRow("SELECT userid from users where userid = $1", userid).Scan(&id)
	if result == sql.ErrNoRows {
		fmt.Println("Found no user")
		CreateUser(userid, username)
	} else {
		fmt.Println("Found user")
		UpdateUserXP(userid)
	}
}

//ToDo
func Level() {

}

//ToDo
func ReturnXP() int {
	return -1
}
