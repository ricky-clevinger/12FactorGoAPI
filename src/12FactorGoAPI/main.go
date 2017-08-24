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
	router.HandleFunc("/home", controllers.HomePage)
	router.HandleFunc("/", controllers.FourOhFour)
	router.HandleFunc("/members", controllers.GetMembers)
	router.HandleFunc("/members/search", controllers.GetSearchedMembers)
	router.HandleFunc("/members/add", controllers.AddMember)
	router.HandleFunc("/members/delete", controllers.DeleteMember)
	router.HandleFunc("/members/update", controllers.UpdateMember)
	router.HandleFunc("/members/id/{id}", controllers.GetMembersById)
	router.HandleFunc("/books", controllers.GetBooks)
	router.HandleFunc("/books/add/{title}/{authF}/{authL}", controllers.AddBook)
	router.HandleFunc("/books/{id}", controllers.GetBooksById)
	router.HandleFunc("/books/delete", controllers.DeleteBook)
	router.HandleFunc("/books/update", controllers.UpdateBook)
	router.HandleFunc("/books/search", controllers.GetSearchedBooks)
	router.HandleFunc("/books/checkedin", controllers.GetCheckedInBooks)
	router.HandleFunc("/books/checkedout", controllers.GetCheckedOutBooks)


	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8081",router))

}


func main() {
	handleRequests()

}
