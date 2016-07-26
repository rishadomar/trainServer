package main

import (
	"fmt"
    "os"
	"net/http"
    "html"
    "database/sql"
    "strconv"
)

var db *sql.DB = nil

func handleGetNext(rw http.ResponseWriter, request *http.Request) {
    fmt.Fprint(rw, "GetNext, %q", html.EscapeString(request.URL.Path))
}

func handlePostPosition(rw http.ResponseWriter, request *http.Request) {
    fmt.Fprint(rw, "PostPosition , %q", html.EscapeString(request.URL.Path))
}

func whereAmI(latitude float64, longitude float64) {
    // 1. find closest location in Location
    // 2. find the route(s) that you're on
    // 3. find the stations around this location
    location, distance := getClosestLocationTo(latitude, longitude)
    fmt.Println("The closest location is: ", location.id, "and distance is: ", distance)
}

func main() {

    if db == nil {
        openDatabase()
        defer db.Close()
    }

    request := os.Args[1]

    if request == "server" {
        http.HandleFunc("/next", handleGetNext)
        http.HandleFunc("/position", handlePostPosition)
        http.ListenAndServe(":8082", nil)
    } else if request == "route" {
        subRequest := os.Args[2]
        handleRequest(subRequest)
    } else if request == "whereami" {
        latitude, _ := strconv.ParseFloat(os.Args[2], 64)
        longitude, _ := strconv.ParseFloat(os.Args[3], 64)
        //stationComingFrom, stationGoingTo := whereAmI(latitude, longitude)
        whereAmI(latitude, longitude)
        //fmt.Println("You are coming from: ", getLocationAsString(&stationComingFrom), " and you are going to: ", getLocationAsString(&stationGoingTo));
    } else {
        fmt.Println("Expected some parameter ...")
    }
}
