package handlers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"gtsanfelix/backend/auth"
	"gtsanfelix/backend/models"
	"gtsanfelix/backend/store"
)

// Login handles POST /api/auth/login.
func Login(s store.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			jsonError(w, "cuerpo de solicitud inválido", http.StatusBadRequest)
			return
		}

		user, ok := s.FindUserByEmail(req.Email)
		if !ok {
			jsonError(w, "credenciales incorrectas", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			jsonError(w, "credenciales incorrectas", http.StatusUnauthorized)
			return
		}

		token, err := auth.GenerateToken(user.ID, user.Email, user.Role)
		if err != nil {
			jsonError(w, "error al generar token", http.StatusInternalServerError)
			return
		}

		safeUser := models.User{ID: user.ID, Email: user.Email, Name: user.Name, Role: user.Role}
		jsonOK(w, models.LoginResponse{Token: token, User: safeUser}, http.StatusOK)
	}
}
