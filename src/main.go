package main

//Author: C Neuhardt
//Last Updated: 8/3/2017

import (
	"os"
	"fmt"
	"database/sql"
	//"html/template"
	"net/http"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//Gets the connection string
var connectionString = os.Getenv("LIBRARY")

type Member struct {
	Member_id int `json:"Member_id"`
	Member_fname string `json:"Member_fname"`
	Member_lname string `json:"Member_lname"`
}

type Book struct{
	Book_id int `json:"Book_id"`
	Book_title string `json:"Book_title"`
	Book_authF string `json:"Book_authF"`
	Book_authL string `json:"Book_authL"`
	Library_id int `json:"Library_id"`
	Book_check int `json:"Book_check"`
	Mid int `json:"Mid"`
	Book_out_date sql.NullString `json:"Book_out_date"`
	Member_fname sql.NullString `json:"Member_fname"`
	Member_lname sql.NullString `json:"Member_lname"`
}

//Checks for errors
func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

//Get Members
func getMembers(w http.ResponseWriter, r *http.Request) {
	var members []Member

	db, err := sql.Open("mysql", connectionString)
	checkErr(err)
	defer db.Close()

	memberRows, err := db.Query("SELECT member_id, member_fname, member_lname FROM member")
	checkErr(err)
	
	for memberRows.Next() {
		m := Member{}
		err = memberRows.Scan(&m.Member_id, &m.Member_fname, &m.Member_lname)
		checkErr(err)
		members = append(members, m)
	}

	json.NewEncoder(w).Encode(members)
}

//Get Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book //Holds Slice of Book Type

	//DB Connection
	db, err := sql.Open("mysql", connectionString)
	checkErr(err)
	defer db.Close() //Close after func GetBook ends

	//Grab entire rows of data within a query
	bookRows, err := db.Query("SELECT b.Book_Id, b.Book_Title, b.Book_AuthFName, b.Book_AuthLName, b.Library_Id, b.Book_Check, b.Mid, date_format(b.Book_Out_Date, '%Y-%m-%d'), m.Member_fname, m.Member_lname FROM books b LEFT JOIN member m ON b.Mid = m.Member_id")
	//Check for Errors in DB Query
	checkErr(err)
	//For Every Book Row that's not null/nil place
	for bookRows.Next() {
		b := Book{} //book type
		err = bookRows.Scan(&b.Book_id, &b.Book_title, &b.Book_authF, &b.Book_authL, &b.Library_id, &b.Book_check, &b.Mid, &b.Book_out_date, &b.Member_fname, &b.Member_lname)
		checkErr(err)
		if b.Book_out_date.Valid{
			books = append(books, b)
		} else {
			b.Book_out_date.String = ""
			books = append(books, b)
		}

	}

	json.NewEncoder(w).Encode(books)
}

func homePage (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/members", getMembers)
	http.HandleFunc("/books", getBooks)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}