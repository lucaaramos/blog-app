package main

import (
	"blog/internal/controllers"
	"blog/internal/repository"
	"blog/internal/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Inicializar el repositorio de publicaciones
	postRepo := repository.NewPostRepository()

	// Inicializar el controlador de blog
	blogController := controllers.NewBlogController(postRepo)

	// Configurar el enrutador
	router := mux.NewRouter()
	routes.SetupRoutes(router, blogController)

	// Iniciar el servidor HTTP
	serverAddr := ":8000"
	log.Printf("Starting server on %s\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, router))
}
