package routes

import (
	"blog/internal/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, blogController *controllers.BlogController) {
	router.HandleFunc("/posts", blogController.CreatePostHandler).Methods("POST")
	router.HandleFunc("/blog/{id}", blogController.GetPostByIDHandler).Methods("GET")
	router.HandleFunc("/blogs", blogController.GetAllBlogsHandler).Methods("GET")
	router.HandleFunc("/update-post/{id}", blogController.UpdateBlogHandler).Methods("PATCH")
	// router.HandleFunc("/delete-post/{id}", blogController.DeleteBlogHandler).Methods("PUT")
}
