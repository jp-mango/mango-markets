package main

import (
	"database/sql"
	"fmt"
	"mangomarkets/internal"

	_ "github.com/lib/pq"
)

func main() {
	_, DB_CONN, _ := internal.LoadEnv()

	db, err := sql.Open("postgres", DB_CONN)
	if err != nil {
		fmt.Println("Error opening db:", err)
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO test(name, age) VALUES($1, $2)`, "jp-mango", 26)
	if err != nil {
		fmt.Println("Error executing INSERT:", err)
		return
	}

	fmt.Println("Data inserted successfully")

}
