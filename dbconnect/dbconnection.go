package dbconnect

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func DBConnection() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname="+os.Getenv("DB_Name")+" sslmode=disable password="+os.Getenv("DB_PASSWORD")+" host="+os.Getenv("DB_HOST"))
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("DataBase Connected Sucessfully")
	}
}
