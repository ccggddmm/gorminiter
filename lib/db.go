package lib

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"fmt"
	"log"
)

var (
	db *sql.DB
)

type Table struct {
	Name string
	Columns []Column
}

type Column struct {
	Name string
	Coltype string
}

func InitDB(host string, port int, username, password, databaseName string)  {
	var connStr string
	if password == "" {
		connStr = fmt.Sprintf("%s@tcp(%s:%d)/%s?", username, host, port, databaseName)
	}
	connStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?", username, password, host, port, databaseName)

	db, _ = sql.Open("mysql", connStr)

	db.SetConnMaxLifetime(1)

	if err := db.Ping(); err != nil{
		log.Fatal("Open DB failed, err is ", err)
	}
	log.Println("DB connnect success")
}

func GetColumnAndType(tablename string) *Table{
	query := `select COLUMN_NAME,DATA_TYPE from information_schema.COLUMNS where table_name = ?`

	rows, err := db.Query(query, tablename)
	if err != nil {
		log.Fatal("Query DB failed, err is ", err)
	}
	defer rows.Close()

	var columns []Column

	for rows.Next() {
		var (
			name   	string
			coltype 	string
		)
		err := rows.Scan(&name, &coltype)
		if err != nil {
			log.Fatal(err)
		}
		columns = append(columns, Column{name, coltype})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &Table{
		Name: tablename,
		Columns: columns,
	}

}
