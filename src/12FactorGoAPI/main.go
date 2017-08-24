package main

//Author: C Neuhardt
//Last Updated: 8/3/2017
//
import (
	"log"
	"net/http"
	"controllers"
)

func handleRequests() {
	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/members", controllers.GetMembers)
	http.HandleFunc("/members/{member_id}", controllers.GetMembersById)
	http.HandleFunc("/members/search/{search_string}", controllers.GetSearchedMembers)
	http.HandleFunc("/books", controllers.GetBooks)
	http.HandleFunc("/books/add/", controllers.AddBooks)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
