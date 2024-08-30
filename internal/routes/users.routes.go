package routes

import (
	"blog/internal/controllers"

	"github.com/gorilla/mux"
)

func SetUpRoutes(router *mux.Router, userController *controllers.UserController) {
	router.HandleFunc("/users", userController.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users", userController.GetAllUsersHandler).Methods("GET")
	// router.HandleFunc("/update-user/{id}", userController.UpdateUserHandler).Methods("PUT")
	// router.HandleFunc("/delete-user/{id}", userController.DeleteUserHandler).Methods("PUT")
	// router.HandleFunc("/login", userController.LoginHandler).Methods("POST")
	// router.HandleFunc("/logout", userController.LogoutHandler).Methods("POST")
	// router.HandleFunc("/refresh-token", userController.RefreshTokenHandler).Methods("POST")
}
