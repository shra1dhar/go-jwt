package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/eaigner/jet"
	"github.com/lib/pq"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	pgUrl, err := pq.ParseURL("postgres://coqgtcjx:5wyJJO...@raja.db.elephantsql.com:5432/coqgtcjx")

	db, err := jet.Open("postgres", pgUrl)
	fmt.Println(db)


	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	router := mux.NewRouter()

	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", signin).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleware(protectedEndpoint)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
type User struct {
	ID	int `json:"id"`
	Email string `json:"email"`
	Password string `json: "password"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}



func signup(w http.ResponseWriter, r *http.Request) {
	var user User
	var error Error

	if user.Email == "" {
		error.Message = "Email is missing"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	json.NewDecoder(r.Body).Decode(&user)
	spew.Dump(user)
}

func signin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signin invoked")
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protected endpoint invoked")
}

func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("TokenVerifyMiddleware invoked")
	return nil
}

