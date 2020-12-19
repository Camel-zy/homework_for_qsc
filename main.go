package main

import (
	"flag"
	"fmt"

	"git.zjuqsc.com/rop/rop-back-neo/database"
)

type CmdlineArgs struct {
	dbconfig string
	dbinit bool
}

func cmdlineArgParse() *CmdlineArgs {
	var user string
	var passwd string
	var host string
	var port string
	var dbname string
	var init bool

	flag.StringVar(&user, "db_user", "", "username for database")
	flag.StringVar(&passwd, "db_passwd", "", "passwd for database")
	flag.StringVar(&host, "db_host", "", "host(ip) for database")
	flag.StringVar(&port, "db_port", "", "port for database")
	flag.StringVar(&dbname, "db_name", "", "db name for database")
	flag.BoolVar(&init, "db_init", false, "init db tables")

	flag.Parse()

	dbconfig := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		user,
		passwd,
		host,
		port,
		dbname)
	return &CmdlineArgs{dbconfig, init}
}

func main() {
	args := cmdlineArgParse()
	dbconfig := args.dbconfig
	dbinit := args.dbinit
	fmt.Printf("db config: [%s]\n", dbconfig)
	database.MakeDB(dbconfig)
	if dbinit {
		fmt.Print("init db tables\n")
		database.Init()
	}
}
