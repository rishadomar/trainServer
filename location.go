package main

import (
    "log"
    "fmt"
)

const LocationTypeStation = "Station";
const LocationTypePoint = "Point";

type Location struct {
    id int
    name string
    latitude float32
    longitude float32
    locationType string
    timeCreated string
}

func printLocation(location *Location) {
    fmt.Println(fmt.Sprintf("%5d %-20.20s %-30.30s %f %f %s",
        location.id,
        location.locationType,
        location.name,
        location.latitude,
        location.longitude,
        location.timeCreated))
}

func getLocationsForRoute(routeId int) ([]Location) {
    locations := readLocationsFromDatabase(routeId)
    return locations
}

func readLocationsFromDatabase(routeId int) ([]Location) {
    // Prepare statement for reading data
    rows, err := db.Query("SELECT * FROM route_location_link WHERE routeId = ?", routeId)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer rows.Close()

    var locations []Location;
    for rows.Next() {
        var id int
        var routeId int
        var locationId int
        var timeCreated string
        err := rows.Scan(&id, &routeId, &locationId, &timeCreated)
        if err != nil {
            log.Fatal(err)
        }

        location := readLocationFromDatabase(locationId)
        locations = append(locations, location)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
    return locations
}

func readLocationFromDatabase(locationId int) (Location) {
    // Prepare statement for reading data
    location := Location{}
    err := db.QueryRow("SELECT * FROM location WHERE ID = ?", locationId).Scan(
        &location.id,
        &location.name,
        &location.latitude,
        &location.longitude,
        &location.locationType,
        &location.timeCreated)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    return location
}
