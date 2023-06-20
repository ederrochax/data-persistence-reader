package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "mydb"
)

func NewConnector(driveName string) (*sql.DB, error) {
	db, err := sql.Open(driveName, fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatal(err)
	}
	migrate(db)
	return db, nil
}

func migrate(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS customer (
			id                        SERIAL PRIMARY KEY,
			cpf                       TEXT,
			private                   BOOLEAN,
			incomplete                BOOLEAN,
			date_last_purchase        DATE NULL,
			ticket_average            DECIMAL,
			ticket_last_purchase      DECIMAL,
			cnpj_most_frequent_store  TEXT,
			cnpj_last_purchase_store  TEXT)`)
	if err != nil {
		log.Fatal(err)
	}
}
