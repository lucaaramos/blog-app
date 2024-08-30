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
	userRepo := repository.NewUserRepository()

	// Inicializar el controlador de blog
	blogController := controllers.NewBlogController(postRepo)
	userController := controllers.NewUserController(userRepo)

	// Configurar el enrutador
	router := mux.NewRouter()
	routes.SetupRoutes(router, blogController)
	routes.SetUpRoutes(router, userController)

	// Iniciar el servidor HTTP
	serverAddr := ":8000"
	log.Printf("Starting server on %s\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, router))
}
