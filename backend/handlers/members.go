package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gtsanfelix/backend/models"
	"gtsanfelix/backend/store"
)

func GetMembers(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonOK(w, s.GetMembers(), http.StatusOK)
	}
}

func GetMember(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			jsonError(w, "id inválido", http.StatusBadRequest)
			return
		}
		member, ok := s.GetMember(id)
		if !ok {
			jsonError(w, "miembro no encontrado", http.StatusNotFound)
			return
		}
		jsonOK(w, member, http.StatusOK)
	}
}

func CreateMember(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m models.Member
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			jsonError(w, "cuerpo de solicitud inválido", http.StatusBadRequest)
			return
		}
		if m.Name == "" || m.Email == "" {
			jsonError(w, "nombre y email son obligatorios", http.StatusBadRequest)
			return
		}
		jsonOK(w, s.CreateMember(m), http.StatusCreated)
	}
}

func UpdateMember(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			jsonError(w, "id inválido", http.StatusBadRequest)
			return
		}
		var m models.Member
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			jsonError(w, "cuerpo de solicitud inválido", http.StatusBadRequest)
			return
		}
		updated, ok := s.UpdateMember(id, m)
		if !ok {
			jsonError(w, "miembro no encontrado", http.StatusNotFound)
			return
		}
		jsonOK(w, updated, http.StatusOK)
	}
}

func DeleteMember(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			jsonError(w, "id inválido", http.StatusBadRequest)
			return
		}
		if !s.DeleteMember(id) {
			jsonError(w, "miembro no encontrado", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
