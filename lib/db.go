package lib

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

type Table struct {
	Name    string
	Columns []Column
}

type Column struct {
	Name     string
	Datatype string
	Key      string
}

func (c Column) IsPK() bool {
	if c.Key == "PRI" {
		return true
	}
	return false
}

func InitDB(host string, port int, username, password, databaseName string) {
	var connStr string
	if password == "" {
		connStr = fmt.Sprintf("%s@tcp(%s:%d)/%s?", username, host, port, databaseName)
	}
	connStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?", username, password, host, port, databaseName)

	fmt.Println(connStr)
	db, _ = sql.Open("mysql", connStr)

	db.SetConnMaxLifetime(1)

	if err := db.Ping(); err != nil {
		log.Fatal("Open DB failed, ", err)
	}
	log.Println("DB connnect success")
}

func GetColumnAndType(tablename string) *Table {
	query := `select COLUMN_NAME,DATA_TYPE,COLUMN_KEY from information_schema.COLUMNS where table_name = ?`

	rows, err := db.Query(query, tablename)
	if err != nil {
		log.Fatal("Query DB failed, err is ", err)
	}
	defer rows.Close()

	var columns []Column

	for rows.Next() {
		var (
			name     string
			datatype string
			key      string
		)
		err := rows.Scan(&name, &datatype, &key)
		if err != nil {
			log.Fatal(err)
		}
		columns = append(columns, Column{name, datatype, key})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &Table{
		Name:    tablename,
		Columns: columns,
	}

}
