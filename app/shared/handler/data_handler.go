package handler

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//User struct
type User struct {
	email string
	pass  string
}

//ReturnStruct of user
func ReturnStruct(email string, pass string) User {
	return User{email: email, pass: pass}
}

//ConnectDB mysql
func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "quannh:quanhen121@tcp(192.168.200.247:3306)/test") //network ip address
	if err != nil {
		panic(err.Error())
	}
	return db
}

//CheckUser valid
func CheckUser(UserInfo User) bool {
	//select multi rows
	results, err := ConnectDB().Query("SELECT * FROM user_info WHERE email = ? AND pass = ?", UserInfo.email, UserInfo.pass)
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()

	var i = 0
	for results.Next() {
		i++
	}

	if i == 0 {
		return false
	}
	return true

}

//CheckEmail exited
func CheckEmail(UserInfo User) bool {
	//select multi rows
	results, err := ConnectDB().Query("SELECT * FROM user_info WHERE email = ?", UserInfo.email)
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()

	var i = 0
	for results.Next() {
		i++
	}

	if i == 0 {
		return false
	}
	return true

}

//InsertUser to database
func InsertUser(UserInfo User) {
	insert, err := ConnectDB().Query("INSERT INTO user_info(email , pass) VALUES( ?, ?)", UserInfo.email, UserInfo.pass)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
