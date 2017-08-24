package controllers

import (
	"net/http"
	"database"
	"encoding/json"
)

//Get Books
func GetBooks(writer http.ResponseWriter, request *http.Request) {

	books := database.ReturnAllBooks()

	json.NewEncoder(writer).Encode(books)
}


func AddBooks(writer http.ResponseWriter, request *http.Request){

	var title string = request.FormValue("title")
	var authF string = request.FormValue("fname")
	var authL string = request.FormValue("lname")

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


}

func GetCheckedOutBooks(writer http.ResponseWriter, request *http.Request){


}
