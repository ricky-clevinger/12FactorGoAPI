package controllers

import (
	"net/http"
	"database"
	"encoding/json"
	"mux"
	"helper"
)

//Get Books
func GetBooks(writer http.ResponseWriter, request *http.Request) {

	books := database.ReturnAllBooks()

	json.NewEncoder(writer).Encode(books)
}

//currently working
func AddBook(writer http.ResponseWriter, request *http.Request){

	//can grab the form if submitted via post if necessary

	/*var title string = helper.HtmlClean(request.FormValue("title"))
	var authF string = helper.HtmlClean(request.FormValue("fname"))
	var authL string = helper.HtmlClean(request.FormValue("lname"))*/

	varArray := mux.Vars(request)

	title := varArray["title"]
	authF := varArray["authF"]
	authL := varArray["authL"]

	if(len(title) == 0 || len(authF) == 0 || len(authL) == 0){

		writer.WriteHeader(http.StatusNotAcceptable)
		writer.Write([]byte("All fields are required"))
	}else{

		database.AddBook(title, authF, authL)

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Book was Added\n"))
	}


}

//currently working

func GetSearchedBooks (writer http.ResponseWriter, request *http.Request){

	varArray := mux.Vars(request)

	searchString := varArray["searchString"]
	searchString = helper.HtmlClean(searchString)

	returnedBooks := database.GetSearchedBook(searchString)

	json.NewEncoder(writer).Encode(returnedBooks)
}

func GetCheckedInBooks(writer http.ResponseWriter, request *http.Request){

	//TODO
}

func GetCheckedOutBooks(writer http.ResponseWriter, request *http.Request){

	//TODO
}

//currently working, but
//	shows delete successful when an invalid id is passed

func DeleteBook(writer http.ResponseWriter, request *http.Request){

	varArray := mux.Vars(request)
	id := varArray["bookId"]

	if len(id) == 0 {

		writer.WriteHeader(http.StatusNotAcceptable)
		writer.Write([]byte("ID is missing, please contact an Administrator"))

	}else{

		database.DeleteBookById(id)

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Book was Deleted\n"))
	}

}

//currently working

func UpdateBook(writer http.ResponseWriter, request *http.Request){

	varArray := mux.Vars(request)

	id := helper.HtmlClean(varArray["bookId"])
	title := helper.HtmlClean(varArray["bookTitle"])
	authF := helper.HtmlClean(varArray["bookAuthF"])
	authL := helper.HtmlClean(varArray["bookAuthL"])

	if(len(id) == 0 || len(title) == 0 || len(authF) == 0 || len(authL) == 0){

		writer.WriteHeader(http.StatusNotAcceptable)
		writer.Write([]byte("All fields are required"))

	}else{

		database.UpdateBook(id, title, authF, authL)

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Book was Updated\n"))
	}


}

//currently working
func GetBooksById(writer http.ResponseWriter, request *http.Request){

	//var id string = helper.HtmlClean(request.FormValue("bookId"))

	vars := mux.Vars(request)

	id := vars["bookid"]


	if(len(id) == 0){

		writer.WriteHeader(http.StatusNotAcceptable)
		writer.Write([]byte("Missing Id - Notify administrator"))
	}else{

		books := database.GetBookById(id)
		json.NewEncoder(writer).Encode(books)
	}
}