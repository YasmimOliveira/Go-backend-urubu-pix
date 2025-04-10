package utils

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "./urubu_do_pix.db")
    if err != nil {
        log.Fatal(err)
    }

    createTable := `
    CREATE TABLE IF NOT EXISTS users (
        id TEXT PRIMARY KEY,
        initial_amount REAL NOT NULL,
        deposit_date TEXT NOT NULL,
        balance REAL NOT NULL
    );`
    _, err = DB.Exec(createTable)
    if err != nil {
        log.Fatal(err)
    }
}