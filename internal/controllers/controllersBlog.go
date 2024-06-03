package controllers

import (
	"blog/internal/models"
	"blog/internal/repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type BlogController struct {
	repo *repository.PostRepository
}

func NewBlogController(repo *repository.PostRepository) *BlogController {
	return &BlogController{repo}
}

func (bc *BlogController) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := bc.repo.CreatePost(r.Context(), &post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (bc *BlogController) GetPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	post, err := bc.repo.GetPostByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(post)
}

func (bc *BlogController) GetAllBlogsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := bc.repo.GetAllBlogs(r.Context())
	if err != nil {
		http.Error(w, "Error getting blogs", http.StatusInternalServerError)
		log.Print("Error getting blog")
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
