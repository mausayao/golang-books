package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Init() *sql.DB {
	gotenv.Load()
	pgURL, err := pq.ParseURL(os.Getenv("PG_URL"))

	logFatal(err)

	db, err := sql.Open("postgres", pgURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
