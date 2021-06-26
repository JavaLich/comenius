package main

import (
    "fmt"
    "log"
    "io"
	"os"
    "net/http"
    "html/template"
    "encoding/json"

    "github.com/gorilla/mux"
)

type User struct {
	Full_name string
	Login     string
}

type JSONRequest struct {
    User string
    Pass string
}

func handler(w http.ResponseWriter, r *http.Request) {
	user := User{Full_name: "Akash Melachuri", Login: "akash"}
	t, err := template.ParseFiles("static/learner.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	t.Execute(w, user)
}

func post(w http.ResponseWriter, r *http.Request) {
    body := r.Body
    buffer, err := io.ReadAll(body)

    if err != nil {
        fmt.Println("Request read error:")
        fmt.Println(err)
    }

    var request JSONRequest
    json.Unmarshal(buffer, &request)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"authenticate": true}`))
}

func main() {
	port := os.Getenv("PORT")
	// port = "8080" // uncomment for local testing
	r := mux.NewRouter()
	r.HandleFunc("/users/", handler).Methods(http.MethodGet)
	r.HandleFunc("/login", post).Methods(http.MethodPost)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Print("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
