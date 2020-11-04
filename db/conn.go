package db

import (
	_ "database/sql"
	"log"
	"os"

	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

func Init() {
	gotenv.Load()
	pgURL, err := pq.ParseURL(os.Getenv("PG_URL"))

	if err != nil {
		log.Fatal(err)
	}

	log.Println(pgURL)

}
