package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/ccggddmm/gorminiter/lib"
)

type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Db   string `json:"db"`
	Table    string `json:"table"`
}

var (
	h      bool
	f      bool
	c      bool
	config Config
)

var templateConf = []byte(
	`{
	"host": "127.0.0.1",
	"port": 3306,
	"user": "root",
	"password": "",
	"db": "mydb",
	"table": "mytable"
}`)

func main() {
	flag.BoolVar(&h, "h", false, "help")
	flag.BoolVar(&f, "f", false, "generate init file")
	flag.BoolVar(&c, "c", false, "using config file")

	flag.StringVar(&(config.Host), "host", "127.0.0.1", "database host, default locolhost")
	flag.IntVar(&(config.Port), "port", 3306, "database port number, default 3306")
	flag.StringVar(&(config.User), "user", "root", "database user name, default root")
	flag.StringVar(&(config.Password), "password", "", "database password, default no password")
	flag.StringVar(&(config.Db), "db", "", "database name")
	flag.StringVar(&(config.Table), "table", "", "table name")

	flag.Parse()

	if h {
		usage()
	}

	if c { //get setting parameters from config file
		content, err := ioutil.ReadFile("setting.json")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Seting.json file not exit, system helped you build template file!")
			fmt.Println(err)
			if ioutil.WriteFile("setting.json", templateConf, 0777) != nil {
				fmt.Fprintln(os.Stderr, "Build template file failed, please try again!")
			}
			os.Exit(-1)
		}
		err = json.Unmarshal(content, &config)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Format of seting.json error, system helped you build template file!")
			if ioutil.WriteFile("setting.json", templateConf, 0777) != nil {
				fmt.Fprintln(os.Stderr, "Build template file failed, please try again!")
			}
			os.Exit(-1)
		}
	}
	//todo file gen or print on console
	check()
	lib.InitDB(config.Host, config.Port, config.User, config.Password, config.Db)
	if f {
		geniniter()
	} else {
		report()
	}

}

func geniniter() {

}

func report() {
	gormStr := lib.BuildStruct(lib.GetTable(config.Table))
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
	-table 	  table name`

	fmt.Fprintln(os.Stderr, os.Args[0], s)
	flag.PrintDefaults()
	os.Exit(-1)
}

func check() {
	if net.ParseIP(config.Host) == nil {
		fmt.Fprintln(os.Stderr, "invalid ip")
		usage()
	}
	if config.Port < 0 {
		fmt.Fprintln(os.Stderr, "invalid port number")
		usage()
	}
	if config.Db == "" {
		fmt.Fprintln(os.Stderr, "please input db name")
		usage()
	}
	if config.Table == "" {
		fmt.Fprintln(os.Stderr, "please input table name")
		usage()
	}
}
