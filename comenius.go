package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "html/template"
)

type User struct {
    Full_name string 
    Login string
}

func handler(w http.ResponseWriter, r *http.Request) {
    user := User {Full_name: "Akash Melachuri", Login: "akash"}
    t, err := template.ParseFiles("static/learner.html")
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
    }
    t.Execute(w, user)
}

func post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "post called"}`))
}
 
func main() {
    r := mux.NewRouter()
    r.HandleFunc("/users/", handler)
    r.HandleFunc("/users/", post).Methods(http.MethodPost)
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
    log.Fatal(http.ListenAndServe(":8080", r))
}
