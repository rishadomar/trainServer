package main

import (
	"fmt"
    "os"
	"net/http"
    "html"
    "database/sql"
    "strconv"
	//"encoding/json"
	"encoding/json"
)

var db *sql.DB = nil

func handleGetNext(rw http.ResponseWriter, request *http.Request) {
    fmt.Fprint(rw, "GetNext, %q", html.EscapeString(request.URL.Path))
}

func handlePostPosition(rw http.ResponseWriter, request *http.Request) {
    fmt.Fprint(rw, "PostPosition , %q", html.EscapeString(request.URL.Path))
}

func handleGetWhereami(rw http.ResponseWriter, request *http.Request) {
	//fmt.Printf("GetWhereami, %q", html.EscapeString(request.URL.Path))
	getParameters := request.URL.Query()
	latitude, _ := strconv.ParseFloat(getParameters.Get("latitude"), 64)
	longitude, _ := strconv.ParseFloat(getParameters.Get("longitude"), 64)
	fmt.Printf("Request whereami: %f %f\n", latitude, longitude)
	location, _ := getClosestLocationTo(latitude, longitude)
	station1, station2 := getStationsSurrounding(location)
	//fmt.Fprint(rw, "Stations are " + station1.name + " and " + station2.name + "\n")
	type WhereamiResponse struct {
		Station1 string
		Station2 string
	}
	jsonResponse, err := json.Marshal(WhereamiResponse {
		Station1: station1.name,
		Station2: station2.name,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	//fmt.Fprint(rw, string(jsonResponse))
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonResponse)
}

func whereAmI(latitude float64, longitude float64) {
    // 1. find closest location in Location
    // 2. find the route(s) that you're on
    // 3. find the stations around this location
    location, distance := getClosestLocationTo(latitude, longitude)
    fmt.Println("The closest location is: ", location.id, "and distance is: ", distance)
    if location.locationType == LocationTypeStation {
        fmt.Println("You are at station: ", location.name)
    } else {
        station1, station2 := getStationsSurrounding(location)
		fmt.Println("You are closest to these stations: ", station1.name, " and ", station2.name)
    }
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
		http.HandleFunc("/whereami", handleGetWhereami)
        http.ListenAndServe(":8083", nil)
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
