package controllers

import (
	"blog-backend/internal/models"
	"blog-backend/internal/repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type BlogController struct {
	postRepo *repository.PostRepository
}

func NewBlogController(postRepo *repository.PostRepository) *BlogController {
	return &BlogController{postRepo}
}

func (bc *BlogController) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Error decoding post data", http.StatusBadRequest)
		return
	}

	err = bc.postRepo.CreatePost(r.Context(), &post)
	if err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		log.Println("Error creating post:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (bc *BlogController) GetPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID := params["id"]

	post, err := bc.postRepo.GetPostByID(r.Context(), postID)
	if err != nil {
		http.Error(w, "Error getting post", http.StatusInternalServerError)
		log.Println("Error getting post:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
