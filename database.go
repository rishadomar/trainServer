package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func openDatabase() {
    var err error
    db, err = sql.Open("mysql", "train:kaluma@/train")
    if err != nil {
        panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
    }

    // Open doesn't open a connection. Validate DSN data:
    err = db.Ping()
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
}
