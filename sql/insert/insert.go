package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:"+os.Getenv("MYSQL_LOCAL_PASSWORD")+"@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")
	stmt.Exec("João")
	stmt.Exec("Maria")
	stmt.Exec("Ayka")

	res, _ := stmt.Exec("Pedro")
	id, _ := res.LastInsertId()
	fmt.Println(id)
	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)

}
