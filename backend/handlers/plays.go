package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gtsanfelix/backend/models"
	"gtsanfelix/backend/store"
)

func GetPlays(s store.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonOK(w, s.GetPlays(), http.StatusOK)
	}
}

func GetPlay(s store.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			jsonError(w, "id inválido", http.StatusBadRequest)
			return
		}
		play, ok := s.GetPlay(id)
		if !ok {
			jsonError(w, "obra no encontrada", http.StatusNotFound)
			return
		}
		jsonOK(w, play, http.StatusOK)
	}
}

func CreatePlay(s store.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var play models.Play
		if err := json.NewDecoder(r.Body).Decode(&play); err != nil {
			jsonError(w, "cuerpo de solicitud inválido", http.StatusBadRequest)
			return
		}
		if play.Title == "" || play.Author == "" {
			jsonError(w, "título y autor son obligatorios", http.StatusBadRequest)
			return
		}
		jsonOK(w, s.CreatePlay(play), http.StatusCreated)
	}
}

func UpdatePlay(s store.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			jsonError(w, "id inválido", http.StatusBadRequest)
			return
		}
		var play models.Play
		if err := json.NewDecoder(r.Body).Decode(&play); err != nil {
			jsonError(w, "cuerpo de solicitud inválido", http.StatusBadRequest)
			return
		}
		updated, ok := s.UpdatePlay(id, play)
		if !ok {
			jsonError(w, "obra no encontrada", http.StatusNotFound)
			return
		}
		jsonOK(w, updated, http.StatusOK)
	}
}

func DeletePlay(s store.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			jsonError(w, "id inválido", http.StatusBadRequest)
			return
		}
		if !s.DeletePlay(id) {
			jsonError(w, "obra no encontrada", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
