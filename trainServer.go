package main

import (
	"fmt"
	"net/http"
    "html"
)

const LocationTypeStation = "Station";
const LocationTypePoint = "Point";

type Location struct {
    latitude float32
    longitude float32
    name string
    locationType string
}

func handleGetNext(rw http.ResponseWriter, request *http.Request) {
    fmt.Fprint(rw, "GetNext, %q", html.EscapeString(request.URL.Path))
}

func handlePostPosition(rw http.ResponseWriter, request *http.Request) {
    fmt.Fprint(rw, "PostPosition , %q", html.EscapeString(request.URL.Path))
}

func main() {
    http.HandleFunc("/next", handleGetNext)
    http.HandleFunc("/position", handlePostPosition)
    http.ListenAndServe(":8082", nil)
}
