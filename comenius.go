package main

import (
    "fmt"
    "log"
    "net/http"
    "html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("static/" + r.URL.Path[1:])
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
    }
    t.Execute(w, nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Users %s", r.URL.Path[len("/users/"):])
}
 
func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)
    http.HandleFunc("/users/", userHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
} 
