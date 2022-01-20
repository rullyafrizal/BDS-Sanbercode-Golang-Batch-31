package main

import (
	"database/sql"
	"fmt"
	"log"
	conf "tugas-14/config"
	"tugas-14/routes"
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