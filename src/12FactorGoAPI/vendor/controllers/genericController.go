package controllers

import (
	"net/http"
	"fmt"
)

func HomePage(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(writer, "Welcome")
}

func FourOhFour(writer http.ResponseWriter, request *http.Request){

	fmt.Fprintf(writer, "Bad Request")
}
