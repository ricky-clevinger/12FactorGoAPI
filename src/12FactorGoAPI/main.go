package main

//Author: C Neuhardt
//Last Updated: 8/3/2017
//
import (
	"log"
	"net/http"
	_ "github.com/gorilla/mux"
	"mux"
	"controllers"
)

func handleRequests() {

	//r := mux.NewRouter()
	//r.Methods("GET", "POST").HandleFunc("/members/{id}", getMembersById)

	router := mux.NewRouter()
	router.HandleFunc("/", controllers.HomePage)
	router.HandleFunc("/members", controllers.GetMembers)
	router.HandleFunc("/members/search", controllers.GetSearchedMembers)
	router.HandleFunc("/members/add", controllers.AddMember)
	router.HandleFunc("/members/delete", controllers.DeleteMember)
	router.HandleFunc("/members/{id}", controllers.GetMembersById)
	router.HandleFunc("/books", controllers.GetBooks)
	router.HandleFunc("/books/add", controllers.AddBooks)
	router.HandleFunc("/books/search", controllers.GetSearchedBooks)
	router.HandleFunc("/books/checkedin", controllers.GetCheckedInBooks)
	router.HandleFunc("/books/checkedout", controllers.GetCheckedOutBooks)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080",nil))

}


func main() {
	handleRequests()

}
