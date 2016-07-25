package main

import (
    "fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
    "log"
)

type Route struct {
    id int
    name string
    code string
    start Location
    points [] Location
    end Location
    timeCreated string
}

func handleRequest(subRequest string) {
   if (subRequest == "list") {
       fmt.Println("list all routes now")
       getAllRoutes()
   } else {
       fmt.Println("Unknown request in route " + subRequest)
   }
}

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

func getAllRoutes() {
    routes := readRoutesFromDatabase()

    for i := range(routes) {
        route := routes[i]
        route.points = getLocationsForRoute(route.id)
        fmt.Println("Route: ", route)
    }
}

func readRoutesFromDatabase() ([]Route) {

    // Prepare statement for reading data
    rows, err := db.Query("SELECT * FROM route")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer rows.Close()

    var routes []Route;
    for rows.Next() {
        route := Route{}
        err := rows.Scan(&route.id, &route.name, &route.code, &route.timeCreated)
        if err != nil {
            log.Fatal(err)
        }
        routes = append(routes, route)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
    return routes
}

