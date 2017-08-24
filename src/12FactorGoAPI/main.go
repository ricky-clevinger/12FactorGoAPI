package main

//Author: C Neuhardt
//Last Updated: 8/3/2017
//
import (
	"net/http"
	"database"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
	"mux"
	"fmt"
	"encoding/json"
	"log"
)


//Get Members
func getMembers(writer http.ResponseWriter, request *http.Request) {

	members := database.ReturnAllMembers()

	json.NewEncoder(writer).Encode(members)
}

//Get MembersById
func getMembersById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	
	members := database.ReturnMembersById(vars["id"])

	json.NewEncoder(w).Encode(members)
}

//Get Books
func getBooks(w http.ResponseWriter, r *http.Request) {

	books := database.ReturnAllBooks()

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
	log.Fatal(http.ListenAndServe(":8081",nil))

}


func main() {
	handleRequests()
	
}
