package main

import "fmt"

const LocationTypeStation = "Station";
const LocationTypePoint = "Point";

type Location struct {
    name string
    latitude float32
    longitude float32
    locationType string
}

func makeNewLocation() {
    fmt.Println("in location");
}
