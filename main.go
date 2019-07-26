package main

import (
	"fmt"
	"os"

	"flag"
	"github.com/ccggddmm/gorminiter/lib"
)


var (
	//cmd  	 string
	host     string
	port     int
	user     string
	password string
	dbname   string
	table 	 string
	//dir		 string
)


//InitDB(host string, port int, username, password, databaseName string)
func main() {
	flag.StringVar(&host, "host", "127.0.0.1", "database host, default locolhost")
	flag.IntVar(&port, "port", 3306, "database port number, default 3306")
	flag.StringVar(&user, "user", "root", "database user name, default root")
	flag.StringVar(&password, "password", "", "database password, default no password")
	flag.StringVar(&dbname, "dbname", "", "database name")
	flag.StringVar(&table, "table", "", "table name")
	//flag.StringVar(&dir, "dir", "./", "generate file path, default current path")

	flag.Parse()

	//todo format check
	//todo file gen or print on console
/*	if host == "" {
		fmt.Fprintln(os.Stderr, "invalid host")
		usage()
	}
	if port < 0 {
		fmt.Fprintln(os.Stderr, "invalid port number")
		usage()
	}*/

	execute()
}


func usage() {
	s := `
Usage: ./gorminiter [options]

Parameters:
	-host	  database host, default locolhost
	-port	  database port number, default 3306
	-username database user name, default root
	-password database password, default no password
	-bdname	  database name
	-table 	  table name
    -dir 	  file to put
	`
	fmt.Fprintln(os.Stderr, os.Args[0], s)
	flag.PrintDefaults()
	os.Exit(-1)
}


func execute() {
	lib.InitDB(host, port, user, password, dbname)
	gormStr := lib.BuildStruct(lib.GetColumnAndType(table))
	fmt.Println(gormStr)
}
