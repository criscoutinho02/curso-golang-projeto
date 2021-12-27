package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:"+os.Getenv("MYSQL_LOCAL_PASSWORD")+"@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values (?,?)")
	stmt.Exec(2000, "Cristina")
	stmt.Exec(2001, "Lola")
	_, err = stmt.Exec(1, "Lola") //chave Duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()

}
