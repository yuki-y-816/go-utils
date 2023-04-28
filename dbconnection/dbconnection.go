package dbconnection

import (
    "database/sql"
    "log"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

func Connect(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("database connection error:", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal("ping error to database:", err)
    }

    fmt.Println("database connection succeeded!")

    return db
}
