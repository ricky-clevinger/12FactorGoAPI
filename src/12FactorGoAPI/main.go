package main

//Author: C Neuhardt
//Last Updated: 8/3/2017
//
import (
	"log"
	"net/http"
	"mux"
	"controllers"
)

func handleRequests() {

	router := mux.NewRouter()
	router.HandleFunc("/home", controllers.HomePage)
	router.HandleFunc("/home/", controllers.HomePage)
	router.HandleFunc("/", controllers.FourOhFour)
	router.HandleFunc("/members", controllers.GetMembers)
	router.HandleFunc("/members/", controllers.GetMembers)
	router.HandleFunc("/members/search/{searchString}", controllers.GetSearchedMembers)
	router.HandleFunc("/members/add", controllers.AddMember)
	router.HandleFunc("/members/delete", controllers.DeleteMember)
	router.HandleFunc("/members/update", controllers.UpdateMember)
	router.HandleFunc("/members/{memid}", controllers.GetMembersById)
	router.HandleFunc("/members/login/{mail}/{pass}", controllers.MemberExist)
	router.HandleFunc("/books", controllers.GetBooks)
	router.HandleFunc("/books/", controllers.GetBooks)
	router.HandleFunc("/books/add/{title}/{authF}/{authL}", controllers.AddBook)
	router.HandleFunc("/books/{bookid}", controllers.GetBooksById)
	router.HandleFunc("/books/delete/{bookId}", controllers.DeleteBook)
	router.HandleFunc("/books/update/{bookId}/{bookTitle}/{bookAuthF}/{bookAuthL}", controllers.UpdateBook)
	router.HandleFunc("/books/search/{searchString}", controllers.GetSearchedBooks)
	router.HandleFunc("/books/checkedin", controllers.GetCheckedInBooks)
	router.HandleFunc("/books/checkedout", controllers.GetCheckedOutBooks)



	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080",router))

}


func main() {
	handleRequests()

}
