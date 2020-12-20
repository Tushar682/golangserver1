package main
import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	_ "strings"
	"time"
)

type User struct{
	Id           string `json:"id" bson:"_id,omitempty"`
	Name         string `json:"name" bson:"name"`
	DoB          string `json:"DoB" bson:"dob"`
	Phone_Number string `json:"Phone_Number" bson:"phone"`
	Email        string `json:"email" bson:"email"`
	time_stamp   int64  `json:"time_stamp" bson:"timestamp"`
}

type Contact struct{
	Useridone     string `json:"useridone"`
	Useridtwo     string `json:"useridtwo"`
	timestamp     int64  `json:"timestamp"`
}

var Users []User
var Contacts []Contact
var client *mongo.Client

func handleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(r.Method)
	switch r.Method {
	case "POST":
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}

		var user User
		_ = json.NewDecoder(r.Body).Decode(&user)
		fmt.Println(user)
		Users = append(Users,user)

		json.NewEncoder(w).Encode(Users)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called "}`))
	case "GET":
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}


func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}

		now := time.Now()
		time_stamp := now.Unix()

		var contact Contact
		_ = json.NewDecoder(r.Body).Decode(&contact)
		contact.timestamp = time_stamp

		fmt.Println(contact)
		Contacts = append(Contacts,contact)

		json.NewEncoder(w).Encode(Contacts)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called "}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func handleAllRequests() {
	http.HandleFunc("/users", handleUser)
	http.HandleFunc("/contacts", handleConnection)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Starting Server...")
	Users = []User{
		User{Id: "identity1", Name: "abcd", DoB: "1/11/1111", Phone_Number: "12345", Email: "abc@gmail.com"},
	}

	Contacts = []Contact{
		{Useridone: "Index1", Useridtwo: "Index2", timestamp: 1554445},
	}

	handleAllRequests()
}
