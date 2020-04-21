package lib

import "database/sql"

//Db varite
var Db *sql.DB
var err error

// Connect connect to database
func Connect() {
	Db, err = sql.Open("postgres", "postgres://elearning:elearning@192.168.1.150/postgres?sslmode=disable")
	checkErr(err)
}

//Disconnect discut db connection
func Disconnect() {
	Db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
