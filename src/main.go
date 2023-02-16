package main

import (
	"github.com/ve-weiyi/go-examplse/src/core/db"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("init invoke")
}

const dsn = "root:mysql7914@(127.0.0.1:3306)/anker?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db.Connect(dsn)
	defer db.CloseDB()

}
