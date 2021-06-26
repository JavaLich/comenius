package main

import (
    "fmt"
    "log"
    "context"
    "io"
	"os"
    "net/http"
    "html/template"
    "encoding/json"

    "github.com/gorilla/mux"
    "firebase.google.com/go"
    "cloud.google.com/go/firestore"
    "google.golang.org/api/option"
)  

type Learner struct { 
    FullName string
	Login     string
}

type Donator struct {
	FullName string
	Login     string
}

type LoginRequest struct {
    User string
    Pass string
}

type CertificateRequest struct {
    User string
}

type DonateRequest struct {
    User string
    amount float32
}
type Contribution struct {
    amount float32
    certificateID string
    date string
}
func learner(w http.ResponseWriter, r *http.Request) {
	user := Learner{FullName: "Akash Melachuri", Login: "akash"}
	t, err := template.ParseFiles("static/learner.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	t.Execute(w, user)
}

func donator(w http.ResponseWriter, r *http.Request) {
	user := Donator{FullName: "Akash Melachuri", Login: "akash"}
	t, err := template.ParseFiles("static/contributor.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	t.Execute(w, user)
}

func loginPost(w http.ResponseWriter, r *http.Request) {
    body := r.Body
    buffer, err := io.ReadAll(body)

    if err != nil {
        fmt.Println("Request read error:")
        fmt.Println(err)
    }

    var request LoginRequest
    json.Unmarshal(buffer, &request)

    // Figure out user data from postgres

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"authenticate": true}`))
}

func loginGet(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/login.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}

func certificate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"submitted": true}`))
}

func donate(w http.ResponseWriter, r *http.Request) {
    body := r.Body
    buffer, err := io.ReadAll(body)

    if err != nil {
        fmt.Println("Request read error:")
        fmt.Println(err)
    }

    var request DonateRequest
    json.Unmarshal(buffer, &request)

    // Figure out user data from postgres

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"submitted": true}`))
}

func test(client *firestore.Client) {
    result, err := client.Collection("sampleData").Doc("inspiration").Set(context.Background(), map[string]int {"test": 5})
    
    if err != nil {
        log.Fatal(err)
    }

    doc := client.Doc("contribution/0rq0QKmbb8IBrOCGBnwd")
    docsnap, _ := doc.Get(context.Background())

    dataMap := docsnap.Data()
    fmt.Println(dataMap)
    
    fmt.Println(result)
}

func main() {
    opt := option.WithCredentialsFile("./serviceAccountKey.json")
    app, err := firebase.NewApp(context.Background(), nil, opt)

    if err != nil {
        log.Fatalf("error initializing app: %v", err)
    }

    client, err := app.Firestore(context.Background())
    if err != nil {
		log.Fatalf("app.Firestore: %v", err)
    }

    test(client)

	port := os.Getenv("PORT")
	port = "8080" // uncomment for local testing
	r := mux.NewRouter()
	r.HandleFunc("/login", loginPost).Methods(http.MethodPost)
	r.HandleFunc("/login", loginGet).Methods(http.MethodGet)
	r.HandleFunc("/certificate", certificate).Methods(http.MethodPost)
	r.HandleFunc("/donate", donate).Methods(http.MethodPost)
	r.PathPrefix("/learners").HandlerFunc(learner).Methods(http.MethodGet)
    r.PathPrefix("/donators").HandlerFunc(donator).Methods(http.MethodGet)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Print("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
