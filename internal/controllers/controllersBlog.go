package controllers

import (
	"blog/internal/models"
	"blog/internal/repository"
	"blog/internal/utils"
	"encoding/json"
	"fmt"
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

func (bc *BlogController) UpdateBlogHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatePost models.Post
	if err := json.NewDecoder(r.Body).Decode(&updatePost); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Bad Request")
		return
	}

	if err := bc.repo.UpdateBlog(r.Context(), id, &updatePost); err != nil {
		http.Error(w, "Error updating post", http.StatusInternalServerError)
		log.Println("Error updating post")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatePost)
	fmt.Println("Post updated succusfuly")
}

func (uc *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Buscar el usuario en la base de datos
	user, err := uc.repo.FindByUsername(r.Context(), creds.UserName)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Validar la contraseña
	if err := user.CheckPassword(creds.Password); err != nil {
		log.Println("Invalid credentials", err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generar el token JWT
	token, err := utils.GenerateJWT(user.UserName)
	if err != nil {
		log.Println("Error generating token", err)
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Devolver el token al usuario
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func (uc *UserController) ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	// Extraer la información del usuario del contexto (que el middleware pone)
	user := r.Context().Value("user").(map[string]interface{})
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println("Error getting data of user", err)
		http.Error(w, "Error al obtener datos del usuario", http.StatusUnauthorized)
		return
	}

	// Devolver los datos del usuario en formato JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Error al codificar datos del usuario: %v\n", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}
}

// func (bc *BlogController) DeleteBlogHandler(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id := params["id"]

// 	if err := bc.repo.DeleteBlog(r.Context(), id); err != nil {
// 		http.Error(w, "Error deleting post", http.StatusInternalServerError)
// 		log.Println("Error deleting post")
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode("ok")
// }
