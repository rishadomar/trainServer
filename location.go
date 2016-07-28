package main

import (
    "log"
    "fmt"
	"strconv"
)

const LocationTypeStation = "Station";
const LocationTypePoint = "Point";

type Location struct {
    id int
    name string
    latitude float64
    longitude float64
    locationType string
    timeCreated string
}

func printLocation(location *Location) {
    fmt.Println(getLocationAsString(location))
}

func getLocationAsString(location *Location) (string) {
    if location == nil {
        return "(none)"
    }
    return fmt.Sprintf("%5d %-20.20s %-30.30s %f %f %s",
        location.id,
        location.locationType,
        location.name,
        location.latitude,
        location.longitude,
        location.timeCreated)
}

func getLocationsForRoute(routeId int) ([]Location) {
    locations := readLocationsFromDatabase(routeId)
    return locations
}

func getStationsSurrounding(location Location) (Location, Location) {
    if (location.locationType == LocationTypeStation) {
        return location,location
    }
    locations := getClosestStationsTo(location.latitude, location.longitude, 2)
    return locations[0], locations[1]
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

func getClosestLocationTo(latitude float64, longitude float64) (Location, float64) {
    var id int
    var distance float64
    query := "SELECT location.ID, min(111.045 * haversine(latitude, longitude, ?, ?)) AS distance FROM location group by id order by distance asc limit 1"
    err := db.QueryRow(query, latitude , longitude).Scan(
        &id,
        &distance)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    return readLocationFromDatabase(id), distance
}

func getClosestStationsTo(latitude float64, longitude float64, count int) ([]Location) {
    query := "SELECT location.ID, min(111.045 * haversine(latitude, longitude, ?, ?)) AS distance FROM location WHERE locationType = 'Station' group by id order by distance asc limit " + strconv.Itoa(count)
	fmt.Println(query)
    rows, err := db.Query(query, latitude, longitude)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer rows.Close()

    var locations []Location;
    for rows.Next() {
        var id int
        var distance float64
        err := rows.Scan(&id, &distance)
        if err != nil {
            log.Fatal(err)
        }
        location := readLocationFromDatabase(id)
        locations = append(locations, location)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
    return locations
}

