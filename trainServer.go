package main

import (
	"fmt"
    "os"
	"net/http"
    "html"
)

func handleGetNext(rw http.ResponseWriter, request *http.Request) {
    fmt.Fprint(rw, "GetNext, %q", html.EscapeString(request.URL.Path))
}

func handlePostPosition(rw http.ResponseWriter, request *http.Request) {
    fmt.Fprint(rw, "PostPosition , %q", html.EscapeString(request.URL.Path))
}

func main() {
    request := os.Args[1]

    if request == "server" {
        http.HandleFunc("/next", handleGetNext)
        http.HandleFunc("/position", handlePostPosition)
        http.ListenAndServe(":8082", nil)
    } else if request == "route" {
        subRequest := os.Args[2]
        handleRequest(subRequest)
    } else {
        fmt.Println("Expected some parameter ...")
    }
}
