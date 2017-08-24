package main

//Author: C Neuhardt
//Last Updated: 8/3/2017
//
import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"database"
	_ "github.com/gorilla/mux"
	"mux"
)


//Get Members
func getMembers(writer http.ResponseWriter, request *http.Request) {

	members := database.ReturnMembers()

	json.NewEncoder(writer).Encode(members)
}

//Get MembersById
func getMembersById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "Category: %v\n", vars["category"])
	
	members := database.ReturnMembersById(vars["id"])

	json.NewEncoder(w).Encode(members)
}

//Get Books
func getBooks(w http.ResponseWriter, r *http.Request) {

	books := database.ReturnBooks()

	json.NewEncoder(w).Encode(books)
}

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome")
}

func handleRequests() {

	//r := mux.NewRouter()
        //r.Methods("GET", "POST").HandleFunc("/members/{id}", getMembersById)

	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/members", getMembers)
	r.HandleFunc("/members/{id}", getMembersById)
	r.HandleFunc("/books", getBooks)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
