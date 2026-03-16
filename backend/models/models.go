package models

// Play represents a theater play in the repertoire.
type Play struct {
	ID       int    `json:"id" firestore:"id"`
	Title    string `json:"title" firestore:"title"`
	Author   string `json:"author" firestore:"author"`
	Genre    string `json:"genre" firestore:"genre"`
	Duration string `json:"duration" firestore:"duration"`
	ImageURL string `json:"imageUrl,omitempty" firestore:"imageUrl"`
	Active   bool   `json:"active" firestore:"active"`
	Summary  string `json:"summary" firestore:"summary"`
}

// Member represents a member of the theater group.
type Member struct {
	ID    int    `json:"id" firestore:"id"`
	Name  string `json:"name" firestore:"name"`
	Role  string `json:"role" firestore:"role"`
	Email string `json:"email" firestore:"email"`
}

// Performance represents a scheduled show on the calendar.
type Performance struct {
	ID        int    `json:"id" firestore:"id"`
	PlayID    int    `json:"playId" firestore:"playId"`
	PlayTitle string `json:"playTitle" firestore:"playTitle"`
	Date      string `json:"date" firestore:"date"`
	Time      string `json:"time" firestore:"time"`
	Venue     string `json:"venue" firestore:"venue"`
}

// User is an admin account (password is never serialised to JSON).
type User struct {
	ID       int    `json:"id" firestore:"id"`
	Email    string `json:"email" firestore:"email"`
	Password string `json:"-" firestore:"password"`
	Name     string `json:"name" firestore:"name"`
	Role     string `json:"role" firestore:"role"`
}

// LoginRequest is the body for POST /api/auth/login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse is returned on successful login.
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
