package main

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"fmt"
	"log"
)

var (
	db *sql.DB
)

type TableInfo struct {
	column_name string
	column_type string
}

func InitDB(host string, port int, username, password, databaseName string)  {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?", username, password, host, port, databaseName)

	db, _ = sql.Open("mysql", connStr)

	db.SetConnMaxLifetime(1)

	if err := db.Ping(); err != nil{
		log.Fatal("Open DB failed, err is ", err)
	}
	log.Println("DB connnect success")
}

func GetColumnAndType(tablename string) *[]TableInfo{
	query := `select COLUMN_NAME,DATA_TYPE from information_schema.COLUMNS where table_name = ?`

	rows, err := db.Query(query, tablename)
	if err != nil {
		log.Fatal("Query DB failed, err is ", err)
	}

	defer rows.Close()

	var result []TableInfo
	for rows.Next() {
		var (
			column_name   	string
			column_type 	string
		)
		err := rows.Scan(&column_name, &column_type)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, TableInfo{column_name, column_type})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	if result != nil {
		return &result
	}
	return nil
}
