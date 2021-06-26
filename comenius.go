package main

import (
    "fmt"
    "log"
    "net/http"
    "html/template"
)

type User struct {
    full_name string
    login string
}

func handler(w http.ResponseWriter, r *http.Request) {
    user := &User {full_name: "Akash Melachuri", login: "akash"}
    t, err := template.ParseFiles("static" + r.URL.Path)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
    }
    t.Execute(w, user)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    user := &User {full_name: "Akash Melachuri", login: "akash"}
    t, err := template.ParseFiles("./static/learner.html")
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
    }
    t.Execute(w, user)
}
 
func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)
    http.HandleFunc("/learner.html", userHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
} 
