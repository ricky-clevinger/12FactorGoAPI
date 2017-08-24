package controllers

import (

	"net/http"
	"database"
	"encoding/json"
)


func GetMembers(writer http.ResponseWriter, request *http.Request) {

	members := database.ReturnAllMembers()

	json.NewEncoder(writer).Encode(members)
}

/**
	Searches for a member by their member id number, using the method
	of the same name in the memberDbFuncs file
 */
func GetMembersById(writer http.ResponseWriter, request *http.Request){

	id := request.PostForm.Get("memId")

	/*vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	members := database.ReturnMembersById(vars["id"])

	json.NewEncoder(w).Encode(members)*/


	if(len(id) == 0){

		writer.WriteHeader(http.StatusNotAcceptable)
		writer.Write([]byte("Missing Id - Notify administrator"))
	}else{

		members := database.GetMemberById(id)
		json.NewEncoder(writer).Encode(members)
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Success"))
	}




}

/**
	Searches for members by both first and last name using the method of
	the same name in the memberDbFuncs file

 */
func GetSearchedMembers(writer http.ResponseWriter, request *http.Request){

	searchString := request.FormValue("s-bar")

	members := database.GetSearchedMember(searchString)

	json.NewEncoder(writer).Encode(members)
}

func AddMember(writer http.ResponseWriter, request *http.Request){

	//TODO
}

func DeleteMember(writer http.ResponseWriter, request *http.Request){

	//TODO
}

func UpdateMember(writer http.ResponseWriter, request *http.Request){

	//TODO
}