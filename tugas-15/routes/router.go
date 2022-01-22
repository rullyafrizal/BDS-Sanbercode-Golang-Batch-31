package routes

import (
	"fmt"
	"log"
	"net/http"
	"tugas-15/controllers"

	"github.com/julienschmidt/httprouter"
)

func SetupRouter() {
	router := httprouter.New()

	router.GET("/students", controllers.GetStudents)
	router.POST("/students", controllers.StoreStudent)
	router.GET("/students/:id", controllers.GetStudent)
	router.PUT("/students/:id", controllers.UpdateStudent)
	router.DELETE("/students/:id", controllers.DestroyStudent)

	router.GET("/subjects", controllers.GetSubjects)
	router.POST("/subjects", controllers.StoreSubject)
	router.GET("/subjects/:id", controllers.GetSubject)
	router.PUT("/subjects/:id", controllers.UpdateSubject)
	router.DELETE("/subjects/:id", controllers.DestroySubject)

	router.GET("/grades", controllers.GetGrades)
	router.POST("/grades", controllers.StoreGrade)
	router.GET("/grades/:id", controllers.GetGrade)
	router.PUT("/grades/:id", controllers.UpdateGrade)
	router.DELETE("/grades/:id", controllers.DestroyGrade)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
