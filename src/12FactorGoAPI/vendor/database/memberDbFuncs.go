package database

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"helper"

)

var connectionString = os.Getenv("LIBRARY")

type Member struct {
	Member_id    int    `json:"Member_id"`
	Member_fname string `json:"Member_fname"`
	Member_lname string `json:"Member_lname"`
}

func ReturnMembers () []Member {

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

func ReturnMembersById (id string) string {



	return id
}
