package routes

import (
	"log"
	"net/http"
	"quiz-3/controllers"
	"quiz-3/middlewares"

	"github.com/julienschmidt/httprouter"
)

func SetupRouter() {
	router := httprouter.New()

	// Bangun Datar
	router.GET("/bangun-datar/segitiga-sama-sisi", controllers.SegitigaSamaSisi)
	router.GET("/bangun-datar/persegi", controllers.Persegi)
	router.GET("/bangun-datar/persegi-panjang", controllers.PersegiPanjang)
	router.GET("/bangun-datar/lingkaran", controllers.Lingkaran)
	router.GET("/bangun-datar/jajar-genjang", controllers.JajarGenjang)

	// Categories
	router.GET("/categories", controllers.IndexCategories)
	router.GET("/categories/:id", controllers.ShowCategory)
	router.POST("/categories", middlewares.Auth(controllers.StoreCategory))
	router.PUT("/categories/:id", middlewares.Auth(controllers.UpdateCategory))
	router.DELETE("/categories/:id", middlewares.Auth(controllers.DestroyCategory))

	// Books
	router.GET("/books", controllers.IndexBooks)
	router.GET("/books/:id", controllers.ShowBook)
	router.POST("/books", middlewares.Auth(controllers.StoreBook))
	router.PUT("/books/:id", middlewares.Auth(controllers.UpdateBook))
	router.DELETE("/books/:id", middlewares.Auth(controllers.DestroyBook))

	log.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
