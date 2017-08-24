package controllers

import (
	"net/http"
	"database"
	"encoding/json"
//	"helper"
	"mux"
)

//Get Books
func GetBooks(writer http.ResponseWriter, request *http.Request) {

	books := database.ReturnAllBooks()

	json.NewEncoder(writer).Encode(books)
}


func AddBook(writer http.ResponseWriter, request *http.Request){

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
		writer.Write([]byte("Success\n"))
	}


}

func GetSearchedBooks (writer http.ResponseWriter, request *http.Request){

	searchString := ""

	database.GetSearchedBook(searchString)
}

func GetCheckedInBooks(writer http.ResponseWriter, request *http.Request){

	//TODO
}

func GetCheckedOutBooks(writer http.ResponseWriter, request *http.Request){

	//TODO
}

func DeleteBook(writer http.ResponseWriter, request *http.Request){

	//TODO
}

func UpdateBook(writer http.ResponseWriter, request *http.Request){

	//TODO
}

func GetBooksById(writer http.ResponseWriter, request *http.Request){

	//TODO
}