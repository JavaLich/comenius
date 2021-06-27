package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type LearnerDetails struct {
	CertificateList            []Certificate
	MoneyRaisedWeek            int64
	TotalContributionsReceived int64
	ContributionHistory        []int64
}

type Certificate struct {
	CertificateURL string    `json:"certificateURL"`
	CourseImageURL string    `json:"courseImageURL"`
	Name           string    `json:"name"`
	Platform       string    `json:"platform"`
	Price          int64     `json:"price"`
	URL            string    `json:"url"`
	Date           time.Time `json:"date"`
	FullyFunded    bool      `json:"fullyFunded"`
	RaisedAmount   int64     `json:"raisedAmount"`
}

type Learner struct {
	FullName string
	Login    string
}

type Contributor struct {
	FullName string
	Login    string
}

type LoginRequest struct {
	User string `json:"username"`
	Pass string `json:"password"`
	Type string `json:"type"`
}

type CertificateRequest struct {
	User string
}

type DonateRequest struct {
	User   string
	amount float32
}
type Contribution struct {
	amount        float32
	certificateID string
	date          string
}

var opt = option.WithCredentialsFile("./serviceAccountKey.json")
var app, _ = firebase.NewApp(context.Background(), nil, opt)
var client, _ = app.Firestore(context.Background())

func learner(w http.ResponseWriter, r *http.Request) {
	user := Learner{FullName: r.URL.Path[len("/learner/"):], Login: r.URL.Path[1:]}
	t, err := template.ParseFiles("static/learner.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	t.Execute(w, user)
}

func contributor(w http.ResponseWriter, r *http.Request) {
    user := Contributor{FullName: r.URL.Path[len("/contributor/"):], Login: r.URL.Path[len("/contributor/"):]}
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

	username := request.User

	iter := client.Collection(request.Type).Documents(context.Background())

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			fmt.Fprintf(w, "Error %v", err)
		}

		if username == doc.Data()["username"].(string) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"authenticate": true}`))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"authenticate": false}`))
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"submitted": true}`))
}

// func queryLearner(user string) *Learner {
//
// }

func getLearnerDetails(w http.ResponseWriter, r *http.Request) {
	var Certs []Certificate

	username := r.URL.Query().Get("username")
	iter := client.Collection("learner").Documents(context.Background())

	var certList []interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Fprintf(w, "Error %v", err)
		}
		if username == doc.Data()["username"] {
			certList = doc.Data()["certificateList"].([]interface{})
			break
		}
	}

	for _, s := range certList {
		learner_doc := client.Doc(s.(string))
		docsnap, _ := learner_doc.Get(context.Background())
		dataMap := docsnap.Data()
		course_doc := client.Doc(dataMap["courseID"].(string))
		coursedocsnap, _ := course_doc.Get(context.Background())
		courseDataMap := coursedocsnap.Data()

		Cert := Certificate{
			CertificateURL: dataMap["certificateURL"].(string),
			CourseImageURL: courseDataMap["courseImageURL"].(string),
			Name:           courseDataMap["name"].(string),
			Platform:       courseDataMap["platform"].(string),
			Price:          courseDataMap["price"].(int64),
			URL:            courseDataMap["url"].(string),
			Date:           dataMap["date"].(time.Time),
			FullyFunded:    dataMap["fullyFunded"].(bool),
			RaisedAmount:   dataMap["raisedAmount"].(int64),
		}
		Certs = append(Certs, Cert)
	}

	learnerDetails := LearnerDetails{
		CertificateList:            Certs,
		MoneyRaisedWeek:            200,
		TotalContributionsReceived: 300,
		ContributionHistory:        []int64{200, 300, 150, 200, 400, 300, 100},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(learnerDetails)
}

func main() {
	port := os.Getenv("PORT")

    port = "8080" // uncomment for local testing

	r := mux.NewRouter()
	r.HandleFunc("/learner_details", getLearnerDetails).Methods(http.MethodGet)
	r.HandleFunc("/login", loginPost).Methods(http.MethodPost)
	r.HandleFunc("/login", loginGet).Methods(http.MethodGet)
	r.HandleFunc("/certificate", certificate).Methods(http.MethodPost)
	r.HandleFunc("/donate", donate).Methods(http.MethodPost)
	r.PathPrefix("/learner").HandlerFunc(learner).Methods(http.MethodGet)
	r.PathPrefix("/contributor").HandlerFunc(contributor).Methods(http.MethodGet)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Print("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
