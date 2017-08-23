package main

//Author: C Neuhardt
//Last Updated: 8/3/2017

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"database"
)


//Get Members
func getMembers(writer http.ResponseWriter, request *http.Request) {

	members := database.ReturnMembers()

	json.NewEncoder(writer).Encode(members)
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
	http.HandleFunc("/", homePage)
	http.HandleFunc("/members", getMembers)
	http.HandleFunc("/books", getBooks)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
