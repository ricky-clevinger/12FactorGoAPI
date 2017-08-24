package controllers

import (
	"net/http"
	"fmt"
	"database"
)

func HomePage(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(writer, "Welcome")
}

func FourOhFour(writer http.ResponseWriter, request *http.Request){

	fmt.Fprintf(writer, "Bad Request")
}

func WriteTransactionLogCheckIn(writer http.ResponseWriter, request *http.Request){

	database.WriteTransactionLogCheckIn()
}

func WriteTransactionLogCheckOut(writer http.ResponseWriter, request *http.Request){

	database.WriteTransactionLogCheckOut()
}
