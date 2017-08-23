package database

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"helper"

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

func ReturnBooks() []Book {

	var books []Book //Holds Slice of Book Type

	//DB Connection
	//connectionString is defined in memberDbFuncs
	db, err := sql.Open("mysql", connectionString)
	helper.CheckErr(err)
	defer db.Close() //Close after func GetBook ends

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