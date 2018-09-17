package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

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
		id SERIAL,
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
	Level()
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
		CreateUser(userid, username)
	} else {
		UpdateUserXP(userid)
	}
}

//ToDo
func Level() {
	start := time.Now()
	var (
		userid string
		xp     int
	)
	rows, err := db.Query("SELECT userid, xp from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&userid, &xp)
		if err != nil {
			log.Fatal(err)
		}
		//example level
		if xp > 0 && xp < 10 {
			_, err = db.Exec("UPDATE users SET level = 1 WHERE userid = $1", userid)
			if err != nil {
				log.Fatal(err)
			}
		} else if xp > 10 && xp < 20 {
			_, err = db.Exec("UPDATE users SET level = 2 WHERE userid = $1", userid)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	elapsed := time.Since(start)
	log.Printf("Level update took %s", elapsed)
	// update every 60 seconds
	time.AfterFunc(time.Second*60, Level)
}

func ReturnXP(userid string) string {
	var xp string
	result := db.QueryRow("SELECT xp from users where userid = $1", userid).Scan(&xp)
	if result == sql.ErrNoRows {
		fmt.Println("Found no user")
	} else {
		return xp
	}
	return ""
}
