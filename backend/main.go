package main

import (
	"log"
	"net/http"

	"gtsanfelix/backend/handlers"
	"gtsanfelix/backend/middleware"
	"gtsanfelix/backend/store"
)

func main() {
	s := store.New()
	mux := http.NewServeMux()

	// ── Auth (public) ─────────────────────────────────────────────────────
	mux.HandleFunc("POST /api/auth/login", handlers.Login(s))

	// ── Plays ─────────────────────────────────────────────────────────────
	// GET is public so the public site can display the repertoire.
	// Write operations require a valid JWT.
	mux.HandleFunc("GET /api/plays", handlers.GetPlays(s))
	mux.HandleFunc("GET /api/plays/{id}", handlers.GetPlay(s))
	mux.HandleFunc("POST /api/plays", middleware.RequireAuth(handlers.CreatePlay(s)))
	mux.HandleFunc("PUT /api/plays/{id}", middleware.RequireAuth(handlers.UpdatePlay(s)))
	mux.HandleFunc("DELETE /api/plays/{id}", middleware.RequireAuth(handlers.DeletePlay(s)))

	// ── Members ───────────────────────────────────────────────────────────
	mux.HandleFunc("GET /api/members", handlers.GetMembers(s))
	mux.HandleFunc("GET /api/members/{id}", handlers.GetMember(s))
	mux.HandleFunc("POST /api/members", middleware.RequireAuth(handlers.CreateMember(s)))
	mux.HandleFunc("PUT /api/members/{id}", middleware.RequireAuth(handlers.UpdateMember(s)))
	mux.HandleFunc("DELETE /api/members/{id}", middleware.RequireAuth(handlers.DeleteMember(s)))

	// ── Performances ──────────────────────────────────────────────────────
	mux.HandleFunc("GET /api/performances", handlers.GetPerformances(s))
	mux.HandleFunc("GET /api/performances/{id}", handlers.GetPerformance(s))
	mux.HandleFunc("POST /api/performances", middleware.RequireAuth(handlers.CreatePerformance(s)))
	mux.HandleFunc("PUT /api/performances/{id}", middleware.RequireAuth(handlers.UpdatePerformance(s)))
	mux.HandleFunc("DELETE /api/performances/{id}", middleware.RequireAuth(handlers.DeletePerformance(s)))

	handler := middleware.CORS(mux)

	log.Println("Backend corriendo en http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}



// TODO:
/**
const firebaseConfig = {
  apiKey: "AIzaSyCp7eWiuk0-oyvVt2BguHtSjifCROijtMo",
  authDomain: "gtsanfelix.firebaseapp.com",
  projectId: "gtsanfelix",
  storageBucket: "gtsanfelix.firebasestorage.app",
  messagingSenderId: "29579318892",
  appId: "1:29579318892:web:ea2accbf4440bfd253aa49",
  measurementId: "G-QYX323TENV"
};
*/
