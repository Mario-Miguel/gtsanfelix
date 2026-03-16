package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gtsanfelix/backend/models"
	"gtsanfelix/backend/store"
)

func GetPerformances(s store.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonOK(w, s.GetPerformances(), http.StatusOK)
	}
}

func GetPerformance(s store.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			jsonError(w, "id inválido", http.StatusBadRequest)
			return
		}
		perf, ok := s.GetPerformance(id)
		if !ok {
			jsonError(w, "actuación no encontrada", http.StatusNotFound)
			return
		}
		jsonOK(w, perf, http.StatusOK)
	}
}

func CreatePerformance(s store.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.Performance
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			jsonError(w, "cuerpo de solicitud inválido", http.StatusBadRequest)
			return
		}
		if p.PlayTitle == "" || p.Date == "" || p.Venue == "" {
			jsonError(w, "título, fecha y sede son obligatorios", http.StatusBadRequest)
			return
		}
		jsonOK(w, s.CreatePerformance(p), http.StatusCreated)
	}
}

func UpdatePerformance(s store.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			jsonError(w, "id inválido", http.StatusBadRequest)
			return
		}
		var p models.Performance
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			jsonError(w, "cuerpo de solicitud inválido", http.StatusBadRequest)
			return
		}
		updated, ok := s.UpdatePerformance(id, p)
		if !ok {
			jsonError(w, "actuación no encontrada", http.StatusNotFound)
			return
		}
		jsonOK(w, updated, http.StatusOK)
	}
}

func DeletePerformance(s store.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			jsonError(w, "id inválido", http.StatusBadRequest)
			return
		}
		if !s.DeletePerformance(id) {
			jsonError(w, "actuación no encontrada", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
