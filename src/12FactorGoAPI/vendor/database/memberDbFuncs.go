package database

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"helper"

	"strconv"
	//"net/http"
	//"encoding/json"
	"fmt"
)

var connectionString = os.Getenv("LIBRARY")
var db *sql.DB

type Member struct {
	Member_id    int    `json:"Member_id"`
	Member_fname string `json:"Member_fname"`
	Member_lname string `json:"Member_lname"`
}

func ReturnAllMembers () []Member {

	var members []Member

	db, err := sql.Open("mysql", connectionString)
	helper.CheckErr(err)
	defer db.Close()

	memberRows, err1 := db.Query("SELECT member_id, member_fname, member_lname FROM member")
	helper.CheckErr(err1)

	for memberRows.Next() {

		m := Member{}
		err = memberRows.Scan(&m.Member_id, &m.Member_fname, &m.Member_lname)
		helper.CheckErr(err)
		members = append(members, m)
	}

	return members
}

//Get Member by ID
func GetMemberById(id string) []Member {
	var member []Member
	var members []Member

	intId, err := strconv.Atoi(id)
	helper.CheckErr(err)

	/*request, err := http.NewRequest("GET", url, nil)
	helper.CheckErr(err)

	client := &http.Client{}

	response, err := client.Do(request)
	helper.CheckErr(err)
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&members)
	helper.CheckErr(err)*/

	memberRows, err1 := db.Query("SELECT member_id, member_fname, member_lname FROM member WHERE member_id = ?", intId)
	helper.CheckErr(err1)

	for memberRows.Next() {

		m := Member{}
		err = memberRows.Scan(&m.Member_id, &m.Member_fname, &m.Member_lname)
		helper.CheckErr(err)
		members = append(members, m)
	}

	for i := 0; i < len(members); i++ {
		membersId := members[i].Member_id

		if membersId == intId {
			member = append(member, members[i])
		}
	}

	return member
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


