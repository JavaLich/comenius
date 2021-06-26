package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<p>Hi there, I love <b>%s</b>!</p>", r.URL.Path)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Users %s", r.URL.Path[len("/users/"):])
}
 
func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/users/", userHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
} 
