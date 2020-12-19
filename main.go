package main

import (
	"flag"
	"fmt"
)

func cmdlineArgParse() string {
	var user string
	var passwd string
	var addr string
	var dbname string

	flag.StringVar(&user, "db_user", "", "username for database")
	flag.StringVar(&passwd, "db_passwd", "", "passwd for database")
	flag.StringVar(&addr, "db_addr", "", "spec(ip:port) for database")
	flag.StringVar(&dbname, "db_name", "", "db name for database")

	flag.Parse()

	dbconfig := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		user,
		passwd,
		addr,
		dbname)
	return dbconfig
}

func main() {
	dbconfig := cmdlineArgParse()
	fmt.Printf("db config uri: [%s]\n", dbconfig)
}
