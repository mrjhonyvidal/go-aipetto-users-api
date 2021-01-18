package users_db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

const (
	MYSQL_USER				= "MYSQL_USER"
	MYSQL_PASSWORD			= "MYSQL_PASSWORD"
	MYSQL_HOST				= "MYSQL_HOST"
	MYSQL_SCHEMA			= "MYSQL_DATABASE"
)

var (
	Client *sql.DB
	username 				= os.Getenv(MYSQL_USER)
	password 				= os.Getenv(MYSQL_PASSWORD)
	host 					= os.Getenv(MYSQL_HOST)
	schema 					= os.Getenv(MYSQL_SCHEMA)
)

func init(){
	//user:password@tcp(host)/db-schema?charset=utf8
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	retryCount := 30
	for {
		if err = Client.Ping(); err != nil {
			if retryCount == 0 {
				log.Fatalf("No able to establish connection to db.")
			}
			log.Printf(fmt.Sprintf("Could not connect to database. Wait 2 seconds. %d retries left...", retryCount))
			retryCount--
			time.Sleep(2 * time.Second)
		}else{
			break;
		}
	}

	log.Println("Database successfully configured")
}