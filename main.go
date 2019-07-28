package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/ccggddmm/gorminiter/lib"
)

var (
	h        bool
	f        bool
	host     string
	port     int
	user     string
	password string
	dbname   string
	table    string
)

//InitDB(host string, port int, username, password, databaseName string)
func main() {
	flag.BoolVar(&h, "h", false, "help")
	flag.BoolVar(&f, "f", false, "generate init file")
	flag.StringVar(&host, "host", "127.0.0.1", "database host, default locolhost")
	flag.IntVar(&port, "port", 3306, "database port number, default 3306")
	flag.StringVar(&user, "user", "root", "database user name, default root")
	flag.StringVar(&password, "password", "", "database password, default no password")
	flag.StringVar(&dbname, "dbname", "", "database name")
	flag.StringVar(&table, "table", "", "table name")

	flag.Parse()

	if h {
		usage()
	}
	//todo file gen or print on console
	if f {
		geniniter()
	} else {
		report()
	}

}

func geniniter() {

}

func report() {
	fmt.Println(host, port, user, password, dbname, table)
	lib.InitDB(host, port, user, password, dbname)
	gormStr := lib.BuildStruct(lib.GetColumnAndType(table))
	fmt.Println("\n", gormStr)
}

func usage() {
	s := `
Usage: ./gorminiter [-h] [options]

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

func check() {
	if net.ParseIP(host) == nil {
		fmt.Fprintln(os.Stderr, "invalid ip")
		usage()
	}
	if port < 0 {
		fmt.Fprintln(os.Stderr, "invalid port number")
		usage()
	}
	if dbname == "" {
		fmt.Fprintln(os.Stderr, "please input db name")
		usage()
	}
	if table == "" {
		fmt.Fprintln(os.Stderr, "please input table name")
		usage()
	}
}
