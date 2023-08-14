package main

//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/lib/pq"
//	"log"
//)
//
//type Person struct {
//	name string
//	age  int
//}
//
//func main() {
//	connStr := "user=user password=user dbname=database sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//	log.Print("Подключение к базе данных успешно")
//	defer func(db *sql.DB) {
//		err := db.Close()
//		if err != nil {
//
//		}
//	}(db)
//	rows, err := db.Query("select * from Person")
//	if err != nil {
//		panic(err)
//	}
//	defer func(rows *sql.Rows) {
//		err := rows.Close()
//		if err != nil {
//
//		}
//	}(rows)
//	var products []Person
//
//	for rows.Next() {
//		p := Person{}
//		err := rows.Scan(&p.name, &p.age)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		products = append(products, p)
//	}
//	for _, p := range products {
//		fmt.Println(p.name, p.age)
//	}
//}
