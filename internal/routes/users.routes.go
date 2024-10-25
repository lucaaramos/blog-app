package routes

import (
	"blog/internal/controllers"
	"blog/internal/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUpRoutes(router *mux.Router, userController *controllers.UserController) {
	router.HandleFunc("/users", userController.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users", userController.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/login", userController.LoginHandler).Methods("POST")

	// Proteger rutas con JWT
	router.Handle("/protected", middleware.JWTAuthMiddleware(http.HandlerFunc(userController.ProtectedEndpoint))).Methods("GET")
}
