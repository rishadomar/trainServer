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

func getClosestLocationTo(latitude float64, longitude float64) (Location, int) {
    var id int
    var distance int
    query := "SELECT location.ID, 111.045 * haversine(latitude, longitude, ?, ?) AS distance FROM location";
    err := db.QueryRow(query, latitude , longitude).Scan(
        &id,
        &distance)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    return readLocationFromDatabase(id), distance
}

/**
	static public function getClosestPointsToXY($latitude, $longitude, udo_Customer $customer = null)
	{
		global $controller;
		$query = 'SELECT udo_location.ID, 111.045 * haversine(gpsLatitude, gpsLongitude, ' . $latitude . ', ' . $longitude . ') AS distance FROM udo_location';

		if ($customer != null) {
			$query .= ', udo_customerlocations';
		}

		$query .= ' WHERE gpsLatitude is not null AND gpsLatitude != 0 AND gpsLongitude is not null AND gpsLatitude != 0 AND gpsLongitude != 0 AND active = 1 AND _type = "udo_Point"';

		if ($customer != null) {
			$query .= ' AND udo_customerlocations.location_id = udo_location.ID AND udo_customerlocations.customer_id = ' . $customer->getId();
		}

		$query .= ' HAVING distance is null OR distance < 10 ORDER BY distance';
		$rows = $controller->db->getAll($query);
		if ($rows === false) {
			$msg = 'Failed getting the points close by.';
			Log::addErrorMessage(	$msg .
				' Query: ' . $query .
				' Reason: ' . $controller->db->errorMsg());
			throw new Exception($msg);
		}

		return $rows;
	}
***/
