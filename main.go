package main

import (
	"github.com/abhaysinghs772/go-crud/db"
	"github.com/abhaysinghs772/go-crud/router"
)

func main() {
	db.InitMysqlDB()
	router.Initrouter().Run()
}
