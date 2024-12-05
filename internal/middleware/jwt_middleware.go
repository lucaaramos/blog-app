package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"blog/internal/utils" // Asegúrate de que esta importación apunte a donde está tu función de validación
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Printf("Authorization header missing")
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validar el token y extraer las claims
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			log.Println("Invalid Token", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Agregar las claims al contexto
		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
