package main

import (
	"database/sql"
	"fmt"
	"log"
	conf "tugas-15/config"
	"tugas-15/routes"
)

var err error

func main() {
	conf.DB, err = sql.Open("mysql", conf.DbUrl())

    if err != nil {
        log.Fatal(err)
    }

    err := conf.DB.Ping()
    if err != nil {
        panic(err.Error()) 
    }

	defer conf.DB.Close()

    fmt.Println("Success")

    routes.SetupRouter()
}