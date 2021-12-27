package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}

	return result
}

func main() {

	strCon := "root:" + os.Getenv("MYSQL_LOCAL_PASSWORD") + "@/"
	db, err := sql.Open("mysql", strCon)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY(id)
	)`)

}
