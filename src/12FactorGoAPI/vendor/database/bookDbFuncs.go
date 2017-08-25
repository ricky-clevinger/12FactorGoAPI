package database

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"helper"
	"fmt"
	"os"
)

type Book struct {
	Book_id       int            `json:"Book_id"`
	Book_title    string         `json:"Book_title"`
	Book_authF    string         `json:"Book_authF"`
	Book_authL    string         `json:"Book_authL"`
	Library_id    int            `json:"Library_id"`
	Book_check    string         `json:"Book_check"`
	Mid           int            `json:"Mid"`
	Book_out_date sql.NullString `json:"Book_out_date"`
	Member_fname  sql.NullString `json:"Member_fname"`
	Member_lname  sql.NullString `json:"Member_lname"`
}

//currently working

func ReturnAllBooks() []Book {

	var books []Book //Holds Slice of Book Type

	//DB Connection
	//connectionString is defined in memberDbFuncs
	db, err := sql.Open("mysql", connectionString)
	helper.CheckErr(err)
	defer db.Close() //Close after func ReturnAllBooks ends

	//Grab entire rows of data within a query
	bookRows, err := db.Query("SELECT b.Book_Id, b.Book_Title, b.Book_AuthFName, b.Book_AuthLName, b.Library_Id, CASE WHEN b.Book_check = 1 THEN 'In' ELSE 'Out' END AS B_check, b.Mid, date_format(b.Book_Out_Date, '%Y-%m-%d'), m.Member_fname, m.Member_lname FROM books b LEFT JOIN member m ON b.Mid = m.Member_id")
	//Check for Errors in DB Query
	helper.CheckErr(err)
	//For Every Book Row that's not null/nil place
	for bookRows.Next() {
		b := Book{} //book type
		err = bookRows.Scan(&b.Book_id, &b.Book_title, &b.Book_authF, &b.Book_authL, &b.Library_id, &b.Book_check, &b.Mid, &b.Book_out_date, &b.Member_fname, &b.Member_lname)
		helper.CheckErr(err)
		if b.Book_out_date.Valid {
			books = append(books, b)
		} else {
			b.Book_out_date.String = ""
			books = append(books, b)
		}

	}

	return books
}

//currently working

func AddBook(title string, authF string, authL string) {
	Book_title := title
	Book_authF := authF
	Book_authL := authL

	db, err := sql.Open("mysql", os.Getenv("LIBRARY"))
	helper.CheckErr(err)
	defer db.Close()

	//Insert new book instance into table
	stmt, err := db.Prepare("INSERT INTO books (Book_Title, Book_AuthFName, Book_AuthLName, Library_Id, Book_Check, Mid) VALUES (?, ?, ?, 1, 1, 0)")
	helper.CheckErr(err)

	stmt.Exec(Book_title, Book_authF, Book_authL)
}

//function is working, but
//	wild cards are not working, troubleshoot

func GetSearchedBook(s string) []Book {

	var books []Book //Hold Slice of Book Type
	search := fmt.Sprintf("%s%s%s", "%", s, "%")

	//DB Connection
	db, err := sql.Open("mysql", connectionString)
	helper.CheckErr(err)
	defer db.Close() //Close after func GetBook ends

	//Prepare entire rows of data within a query
	bookRows, err := db.Query("SELECT b.Book_Id, b.Book_Title, b.Book_AuthFName, b.Book_AuthLName, b.Library_Id, CASE WHEN b.Book_check = 1 THEN 'In' ELSE 'Out' END AS B_check, b.Mid, date_format(b.Book_Out_Date, '%Y-%m-%d'), m.Member_fname, m.Member_lname FROM books b LEFT JOIN member m ON b.Mid = m.Member_id WHERE book_title like ?", search)

	//Check for Errors in DB the Query
	helper.CheckErr(err)

	//For Every Book Row that's not null/nil place
	for bookRows.Next() {
		b := Book{} //book type
		err = bookRows.Scan(&b.Book_id, &b.Book_title, &b.Book_authF, &b.Book_authL, &b.Library_id, &b.Book_check, &b.Mid, &b.Book_out_date, &b.Member_fname, &b.Member_lname)
		helper.CheckErr(err)
		if b.Book_out_date.Valid {
			books = append(books, b)
		} else {
			b.Book_out_date.String = ""
			books = append(books, b)
		}

	}

	return books
}



func GetCheckedOutBooks() []Book {
	var books []Book

	//DB Connection
	db, err := sql.Open("mysql", connectionString)
	helper.CheckErr(err)
	defer db.Close() //Close after func GetBook ends

	bookRows, err := db.Query("SELECT * FROM books WHERE Book_Check=2")
	helper.CheckErr(err)

	for bookRows.Next() {

		b := Book{}
		err = bookRows.Scan(&b.Book_id, &b.Book_title, &b.Book_authF, &b.Book_authL)
		helper.CheckErr(err)
		books = append(books, b)
	}

	return books

}



func GetCheckedInBooks() []Book {
	var books []Book

	//DB Connection
	db, err := sql.Open("mysql", connectionString)
	helper.CheckErr(err)
	defer db.Close() //Close after func GetBook ends

	bookRows, err := db.Query("SELECT * FROM books WHERE Book_Check=1")

	for bookRows.Next() {

		b := Book{}
		err = bookRows.Scan(&b.Book_id, &b.Book_title, &b.Book_authF, &b.Book_authL)
		helper.CheckErr(err)
		books = append(books, b)
	}

	return books
}



//currently working
func GetBookById(id string) []Book {

	var books []Book


	db, err := sql.Open("mysql", connectionString)
	helper.CheckErr(err)
	defer db.Close()

	bookRows, err1 := db.Query("SELECT book_id, book_title, book_authfname, book_authlname FROM books WHERE book_id = ?", id)
	helper.CheckErr(err1)

	for bookRows.Next() {

		b := Book{}
		err = bookRows.Scan(&b.Book_id, &b.Book_title, &b.Book_authF, &b.Book_authL)
		helper.CheckErr(err)
		books = append(books, b)
	}
/*
	for i := 0; i < len(books); i++ {
		bookId := books[i].Book_id

		if bookId == intId {
			book = append(book, books[i])
		}
	}*/

	return books
}

//currently working

func DeleteBookById(id string){


	db, err := sql.Open("mysql", os.Getenv("LIBRARY"))
	helper.CheckErr(err)
	defer db.Close()

	stmt, err := db.Prepare("Delete from books where book_id = ?")
	helper.CheckErr(err)

	stmt.Exec(id)
}

func UpdateBook(id, title, authF, authL string){

	db, err := sql.Open("mysql", os.Getenv("LIBRARY"))
	helper.CheckErr(err)
	defer db.Close()

	stmt, err := db.Prepare("Update books Set book_title = ?, book_authfname = ?, book_authlname = ? where book_id = ?")
	helper.CheckErr(err)

	stmt.Exec(title, authF, authL, id)
}