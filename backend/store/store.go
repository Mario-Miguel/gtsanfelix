// Package store holds the in-memory mock data store.
// Replace with a real database driver when ready.
package store

import (
	"sync"

	"golang.org/x/crypto/bcrypt"

	"gtsanfelix/backend/models"
)

// Store is a thread-safe in-memory data store.
type Store struct {
	mu           sync.RWMutex
	plays        []models.Play
	members      []models.Member
	performances []models.Performance
	users        []models.User
	nextPlayID   int
	nextMemberID int
	nextPerfID   int
}

// New returns a Store seeded with mock data.
func New() *Store {
	adminHash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	return &Store{
		nextPlayID:   7,
		nextMemberID: 6,
		nextPerfID:   6,

		users: []models.User{
			{ID: 1, Email: "gtvaldesoto@gmail.com", Password: string(adminHash), Name: "Admin", Role: "admin"},
		},

		plays: []models.Play{
			{ID: 1, Title: "La Fonda", Author: "Pepín García", Genre: "Comedia", Duration: "80 min", Active: true, Summary: "Lorem ipsum"},
			{ID: 2, Title: "Al Loro", Author: "Pepín García", Genre: "Comedia", Duration: "80 min", Active: true, Summary: "Lorem ipsum"},
			{ID: 3, Title: "Cróniques de Valdesoto", Author: "Pepín García", Genre: "Comedia", Duration: "60 - 80 min", Active: true, Summary: "Lorem ipsum"},
			{ID: 3, Title: "Torbolín", Author: "Pepín García", Genre: "Comedia", Duration: "80 min", Active: false, Summary: "Lorem ipsum"},
		},

		members: []models.Member{
			{ID: 1, Name: "Pepín García", Role: "Actor & Director & Autor", Email: ""},
			{ID: 2, Name: "Nacho Martínez", Role: "Actor & Director", Email: ""},
			{ID: 3, Name: "Cristina Palacio", Role: "Actriz", Email: ""},
			{ID: 4, Name: "Alex Cueto", Role: "Actor", Email: ""},
			{ID: 5, Name: "Cristina García", Role: "Actriz", Email: ""},
			{ID: 6, Name: "Alba Martínez", Role: "Actriz", Email: ""},
			{ID: 7, Name: "Mario Miguel", Role: "Técnico de luces y sonido", Email: ""},
		},

		performances: []models.Performance{
			{ID: 1, PlayID: 1, PlayTitle: "La Casa de Bernarda Alba", Date: "2024-11-05", Time: "20:00", Venue: "Teatro Principal"},
			{ID: 2, PlayID: 1, PlayTitle: "Cyrano de Bergerac", Date: "2024-11-18", Time: "19:30", Venue: "Teatro La Latina"},
			{ID: 3, PlayID: 3, PlayTitle: "Bodas de Sangre", Date: "2024-12-02", Time: "20:30", Venue: "Teatro del Barrio"},
			{ID: 4, PlayID: 5, PlayTitle: "Midsummer Night", Date: "2024-12-10", Time: "20:00", Venue: "Teatro del Barrio"},
			{ID: 5, PlayID: 2, PlayTitle: "Esperando a Godot", Date: "2024-12-15", Time: "19:00", Venue: "Teatro Principal"},
		},
	}
}

// ── Plays ──────────────────────────────────────────────────────────────────

func (s *Store) GetPlays() []models.Play {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]models.Play, len(s.plays))
	copy(out, s.plays)
	return out
}

func (s *Store) GetPlay(id int) (models.Play, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, p := range s.plays {
		if p.ID == id {
			return p, true
		}
	}
	return models.Play{}, false
}

func (s *Store) CreatePlay(p models.Play) models.Play {
	s.mu.Lock()
	defer s.mu.Unlock()
	p.ID = s.nextPlayID
	s.nextPlayID++
	s.plays = append(s.plays, p)
	return p
}

func (s *Store) UpdatePlay(id int, p models.Play) (models.Play, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, existing := range s.plays {
		if existing.ID == id {
			p.ID = id
			s.plays[i] = p
			return p, true
		}
	}
	return models.Play{}, false
}

func (s *Store) DeletePlay(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, p := range s.plays {
		if p.ID == id {
			s.plays = append(s.plays[:i], s.plays[i+1:]...)
			return true
		}
	}
	return false
}

// ── Members ────────────────────────────────────────────────────────────────

func (s *Store) GetMembers() []models.Member {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]models.Member, len(s.members))
	copy(out, s.members)
	return out
}

func (s *Store) GetMember(id int) (models.Member, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, m := range s.members {
		if m.ID == id {
			return m, true
		}
	}
	return models.Member{}, false
}

func (s *Store) CreateMember(m models.Member) models.Member {
	s.mu.Lock()
	defer s.mu.Unlock()
	m.ID = s.nextMemberID
	s.nextMemberID++
	s.members = append(s.members, m)
	return m
}

func (s *Store) UpdateMember(id int, m models.Member) (models.Member, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, existing := range s.members {
		if existing.ID == id {
			m.ID = id
			s.members[i] = m
			return m, true
		}
	}
	return models.Member{}, false
}

func (s *Store) DeleteMember(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, m := range s.members {
		if m.ID == id {
			s.members = append(s.members[:i], s.members[i+1:]...)
			return true
		}
	}
	return false
}

// ── Performances ───────────────────────────────────────────────────────────

func (s *Store) GetPerformances() []models.Performance {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]models.Performance, len(s.performances))
	copy(out, s.performances)
	return out
}

func (s *Store) GetPerformance(id int) (models.Performance, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, p := range s.performances {
		if p.ID == id {
			return p, true
		}
	}
	return models.Performance{}, false
}

func (s *Store) CreatePerformance(p models.Performance) models.Performance {
	s.mu.Lock()
	defer s.mu.Unlock()
	p.ID = s.nextPerfID
	s.nextPerfID++
	s.performances = append(s.performances, p)
	return p
}

func (s *Store) UpdatePerformance(id int, p models.Performance) (models.Performance, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, existing := range s.performances {
		if existing.ID == id {
			p.ID = id
			s.performances[i] = p
			return p, true
		}
	}
	return models.Performance{}, false
}

func (s *Store) DeletePerformance(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, p := range s.performances {
		if p.ID == id {
			s.performances = append(s.performances[:i], s.performances[i+1:]...)
			return true
		}
	}
	return false
}

// ── Auth ───────────────────────────────────────────────────────────────────

func (s *Store) FindUserByEmail(email string) (models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, u := range s.users {
		if u.Email == email {
			return u, true
		}
	}
	return models.User{}, false
}
