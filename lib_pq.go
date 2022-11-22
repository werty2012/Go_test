package main

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 10101
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func getPostgresData() string {
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Established a successful connection!")
	rows, err := db.Query("select * from test")
	defer rows.Close()

	result := ""
	for rows.Next() {
		var (
			age  int64
			name string
		)
		if err := rows.Scan(&name, &age); err != nil {
			log.Fatal(err)
		}
		result += fmt.Sprintf("id %d name is %s\n", age, name)
		//log.Printf("id %d name is %s\n", age, name)
	}

	return result
}
