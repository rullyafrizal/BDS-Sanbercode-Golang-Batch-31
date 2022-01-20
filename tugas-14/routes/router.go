package routes

import (
	"fmt"
	"log"
	"net/http"
	"tugas-14/controllers"

	"github.com/julienschmidt/httprouter"
)

func SetupRouter() {
	router := httprouter.New()

	router.GET("/", controllers.GetStudents)
	router.POST("/", controllers.StoreStudent)

    fmt.Println("Server Running at Port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}