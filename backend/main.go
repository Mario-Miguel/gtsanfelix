package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"gtsanfelix/backend/handlers"
	"gtsanfelix/backend/middleware"
	"gtsanfelix/backend/store"
)

const serviceAccountPath = "serviceAccount.json"

func main() {
	ctx := context.Background()

	var db store.DB
	if projectID, credFile := resolveFirebaseConfig(); projectID != "" {
		fs, err := store.NewFirestore(ctx, projectID, credFile)
		if err != nil {
			log.Fatalf("Firestore: %v", err)
		}
		defer fs.Close()
		fs.SeedIfEmpty()
		db = fs
		log.Printf("Usando Firestore (proyecto: %s)", projectID)
	}

	mux := http.NewServeMux()

	// ── Auth (public) ─────────────────────────────────────────────────────
	mux.HandleFunc("POST /api/auth/login", handlers.Login(db))

	// ── Plays ─────────────────────────────────────────────────────────────
	mux.HandleFunc("GET /api/plays", handlers.GetPlays(db))
	mux.HandleFunc("GET /api/plays/{id}", handlers.GetPlay(db))
	mux.HandleFunc("POST /api/plays", middleware.RequireAuth(handlers.CreatePlay(db)))
	mux.HandleFunc("PUT /api/plays/{id}", middleware.RequireAuth(handlers.UpdatePlay(db)))
	mux.HandleFunc("DELETE /api/plays/{id}", middleware.RequireAuth(handlers.DeletePlay(db)))

	// ── Members ───────────────────────────────────────────────────────────
	mux.HandleFunc("GET /api/members", handlers.GetMembers(db))
	mux.HandleFunc("GET /api/members/{id}", handlers.GetMember(db))
	mux.HandleFunc("POST /api/members", middleware.RequireAuth(handlers.CreateMember(db)))
	mux.HandleFunc("PUT /api/members/{id}", middleware.RequireAuth(handlers.UpdateMember(db)))
	mux.HandleFunc("DELETE /api/members/{id}", middleware.RequireAuth(handlers.DeleteMember(db)))

	// ── Performances ──────────────────────────────────────────────────────
	mux.HandleFunc("GET /api/performances", handlers.GetPerformances(db))
	mux.HandleFunc("GET /api/performances/{id}", handlers.GetPerformance(db))
	mux.HandleFunc("POST /api/performances", middleware.RequireAuth(handlers.CreatePerformance(db)))
	mux.HandleFunc("PUT /api/performances/{id}", middleware.RequireAuth(handlers.UpdatePerformance(db)))
	mux.HandleFunc("DELETE /api/performances/{id}", middleware.RequireAuth(handlers.DeletePerformance(db)))

	handler := middleware.CORS(mux)

	log.Println("Backend corriendo en http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}

// resolveFirebaseConfig returns the project ID and credentials file path.
//
// Priority:
//  1. GOOGLE_APPLICATION_CREDENTIALS env var (production / Kubernetes Secret mount)
//  2. serviceAccount.json next to the binary (local dev convenience)
//  3. FIREBASE_PROJECT_ID only — no cred file, relies on Workload Identity (GKE)
func resolveFirebaseConfig() (projectID, credFile string) {
	// 1. Production: GOOGLE_APPLICATION_CREDENTIALS points to a mounted Secret file.
	if cred := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"); cred != "" {
		return projectIDFromFile(cred, os.Getenv("FIREBASE_PROJECT_ID")), cred
	}
	// 2. Local dev: serviceAccount.json in the working directory.
	if data, err := os.ReadFile(serviceAccountPath); err == nil {
		var sa struct {
			ProjectID string `json:"project_id"`
		}
		if err := json.Unmarshal(data, &sa); err == nil && sa.ProjectID != "" {
			return sa.ProjectID, serviceAccountPath
		}
	}
	// 3. GKE Workload Identity: no credentials file needed.
	return os.Getenv("FIREBASE_PROJECT_ID"), ""
}

// projectIDFromFile reads the project_id field from a service account JSON file.
// Falls back to the provided default if the file cannot be read or parsed.
func projectIDFromFile(path, fallback string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return fallback
	}
	var sa struct {
		ProjectID string `json:"project_id"`
	}
	if err := json.Unmarshal(data, &sa); err != nil || sa.ProjectID == "" {
		return fallback
	}
	return sa.ProjectID
}
