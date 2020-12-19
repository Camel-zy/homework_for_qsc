package main

import (
	"flag"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"git.zjuqsc.com/rop/rop-back-neo/database"
)

func cmdlineDbArgParse() string {
	var user string
	var passwd string
	var host string
	var port string
	var dbname string

	flag.StringVar(&user, "db_user", "", "username for database")
	flag.StringVar(&passwd, "db_passwd", "", "passwd for database")
	flag.StringVar(&host, "db_host", "", "host(ip) for database")
	flag.StringVar(&port, "db_port", "", "port for database")
	flag.StringVar(&dbname, "db_name", "", "db name for database")

	flag.Parse()

	dbconfig := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		user,
		passwd,
		host,
		port,
		dbname)
	return dbconfig
}

func makeDB(dbconfig string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbconfig), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}

func main() {
	dbconfig := cmdlineDbArgParse()
	fmt.Printf("db config: [%s]\n", dbconfig)
	database.DB = makeDB(dbconfig)
}
