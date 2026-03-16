package models

// Play represents a theater play in the repertoire.
type Play struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Genre    string `json:"genre"`
	Duration string `json:"duration"`
	ImageURL string `json:"imageUrl,omitempty"`
	Active   bool   `json:"active"`
}

// Member represents a member of the theater group.
type Member struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

// Performance represents a scheduled show on the calendar.
type Performance struct {
	ID        int    `json:"id"`
	PlayID    int    `json:"playId"`
	PlayTitle string `json:"playTitle"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Venue     string `json:"venue"`
}

// User is an admin account (password is never serialised to JSON).
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Name     string `json:"name"`
	Role     string `json:"role"`
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
