package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var DB *sql.DB

func InitPostgres() {
    var err error
    DB, err = sql.Open("postgres", "user=postgres password=vakhaboff dbname=auth_service sslmode=disable")
    if err != nil {
        log.Fatalf("Failed to connect to PostgreSQL: %v", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatalf("Failed to ping PostgreSQL: %v", err)
    }

    log.Println("Connected to PostgreSQL")
}
