package database

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"helper"
	"fmt"
)

var connectionString = os.Getenv("LIBRARY")
var db *sql.DB

type Member struct {
	Member_id    int    `json:"Member_id"`
	Member_fname string `json:"Member_fname"`
	Member_lname string `json:"Member_lname"`
	Email string `json:"Email"`
	Password string `json:"Password"`
	Role string `json:"Role"`
}

func ReturnAllMembers () []Member {

	var members []Member

	db, err := sql.Open("mysql", connectionString)
	helper.CheckErr(err)
	defer db.Close()

	memberRows, err1 := db.Query("SELECT member_id, member_fname, member_lname, Email, Password, Role FROM member")
	helper.CheckErr(err1)

	for memberRows.Next() {

		m := Member{}
		err = memberRows.Scan(&m.Member_id, &m.Member_fname, &m.Member_lname, &m.Email, &m.Password, &m.Role)
		helper.CheckErr(err)
		members = append(members, m)
	}

	return members
}


func ReturnMembersById (id string) []Member {


	var members []Member

	db, err := sql.Open("mysql", connectionString)
	helper.CheckErr(err)
	defer db.Close()

	memberRows, err1 := db.Query("SELECT member_id, member_fname, member_lname FROM member WHERE member_id=" + id)
	helper.CheckErr(err1)

	for memberRows.Next() {

		m := Member{}
		err = memberRows.Scan(&m.Member_id, &m.Member_fname, &m.Member_lname)
		helper.CheckErr(err)
		members = append(members, m)
	}

	return members
}

func MemberExist (mail, pass string) []Member {


	var members []Member

	db, err := sql.Open("mysql", connectionString)
	helper.CheckErr(err)
	defer db.Close()

	memberRows, err1 := db.Query("SELECT Email, Password, Role FROM member WHERE Email like ? AND Password like ?", mail, pass)
	helper.CheckErr(err1)

	for memberRows.Next() {

		m := Member{}
		err = memberRows.Scan(&m.Email, &m.Password, &m.Role)
		helper.CheckErr(err)
		members = append(members, m)
	}

	return members
}


//Get Members using search
func GetSearchedMember(s string) []Member {

	var members []Member //Hold Slice of Member Type
	search :=  fmt.Sprintf("%s%s%s", "%", s, "%")

	//DB Connection
	db, err := sql.Open("mysql", connectionString)
	helper.CheckErr(err)
	defer db.Close() //Close after func GetSearchedMember(s string) ends

	//Prepare entire rows of data within a query
	memberRows, err := db.Query("SELECT member_id, member_fname, member_lname FROM member WHERE member_fname like ? OR member_lname like ?", search, search)

	//Check for Errors in DB the Query
	helper.CheckErr(err)

	//For Every Member Row that's not null/nil place
	for memberRows.Next() {
		m := Member{}
		err = memberRows.Scan(&m.Member_id, &m.Member_fname, &m.Member_lname)
		helper.CheckErr(err)
		members = append(members, m)
	}

	return members
}



