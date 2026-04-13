package store

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"gtsanfelix/backend/models"
)

type DB interface {
	GetPlays() []models.Play
	GetPlay(id int) (models.Play, bool)
	CreatePlay(p models.Play) models.Play
	UpdatePlay(id int, p models.Play) (models.Play, bool)
	DeletePlay(id int) bool

	GetMembers() []models.Member
	GetMember(id int) (models.Member, bool)
	CreateMember(m models.Member) models.Member
	UpdateMember(id int, m models.Member) (models.Member, bool)
	DeleteMember(id int) bool

	GetPerformances() []models.Performance
	GetPerformance(id int) (models.Performance, bool)
	CreatePerformance(p models.Performance) models.Performance
	UpdatePerformance(id int, p models.Performance) (models.Performance, bool)
	DeletePerformance(id int) bool

	FindUserByEmail(email string) (models.User, bool)
}

// FirestoreStore implements DB using Cloud Firestore.
type FirestoreStore struct {
	client *firestore.Client
	ctx    context.Context
}

// NewFirestore connects to Firestore. credentialsFile is the path to a service
// account JSON file; pass an empty string to use Application Default Credentials.
func NewFirestore(ctx context.Context, projectID string, credentialsFile string) (*FirestoreStore, error) {
	var opts []option.ClientOption
	if credentialsFile != "" {
		opts = append(opts, option.WithCredentialsFile(credentialsFile))
	}
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: projectID}, opts...)
	if err != nil {
		return nil, fmt.Errorf("firebase.NewApp: %w", err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("app.Firestore: %w", err)
	}
	return &FirestoreStore{client: client, ctx: ctx}, nil
}

// Close releases the Firestore client.
func (fs *FirestoreStore) Close() {
	fs.client.Close()
}

// nextID atomically increments a named counter and returns the previous value.
func (fs *FirestoreStore) nextID(name string) int {
	ref := fs.client.Collection("counters").Doc(name)
	var id int
	err := fs.client.RunTransaction(fs.ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(ref)
		if err != nil {
			// First time: initialise counter at 1, return 1
			id = 1
			return tx.Set(ref, map[string]interface{}{"next": int64(2)})
		}
		raw, err := doc.DataAt("next")
		if err != nil {
			return err
		}
		n, ok := raw.(int64)
		if !ok {
			return fmt.Errorf("counter %q: unexpected type %T", name, raw)
		}
		id = int(n)
		return tx.Set(ref, map[string]interface{}{"next": n + 1})
	})
	if err != nil {
		log.Printf("nextID(%s): %v", name, err)
	}
	return id
}

// ── Plays ──────────────────────────────────────────────────────────────────

func (fs *FirestoreStore) GetPlays() []models.Play {
	iter := fs.client.Collection("plays").Documents(fs.ctx)
	defer iter.Stop()
	var plays []models.Play
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("GetPlays: %v", err)
			return nil
		}
		var p models.Play
		if err := doc.DataTo(&p); err != nil {
			log.Printf("GetPlays DataTo: %v", err)
			continue
		}
		plays = append(plays, p)
	}
	return plays
}

func (fs *FirestoreStore) GetPlay(id int) (models.Play, bool) {
	doc, err := fs.client.Collection("plays").Doc(strconv.Itoa(id)).Get(fs.ctx)
	if err != nil {
		return models.Play{}, false
	}
	var p models.Play
	if err := doc.DataTo(&p); err != nil {
		return models.Play{}, false
	}
	return p, true
}

func (fs *FirestoreStore) CreatePlay(p models.Play) models.Play {
	p.ID = fs.nextID("plays")
	if _, err := fs.client.Collection("plays").Doc(strconv.Itoa(p.ID)).Set(fs.ctx, p); err != nil {
		log.Printf("CreatePlay: %v", err)
	}
	return p
}

func (fs *FirestoreStore) UpdatePlay(id int, p models.Play) (models.Play, bool) {
	p.ID = id
	ref := fs.client.Collection("plays").Doc(strconv.Itoa(id))
	doc, err := ref.Get(fs.ctx)
	if err != nil || !doc.Exists() {
		return models.Play{}, false
	}
	if _, err := ref.Set(fs.ctx, p); err != nil {
		log.Printf("UpdatePlay: %v", err)
		return models.Play{}, false
	}
	return p, true
}

func (fs *FirestoreStore) DeletePlay(id int) bool {
	ref := fs.client.Collection("plays").Doc(strconv.Itoa(id))
	doc, err := ref.Get(fs.ctx)
	if err != nil || !doc.Exists() {
		return false
	}
	if _, err := ref.Delete(fs.ctx); err != nil {
		log.Printf("DeletePlay: %v", err)
		return false
	}
	return true
}

// ── Members ────────────────────────────────────────────────────────────────

func (fs *FirestoreStore) GetMembers() []models.Member {
	iter := fs.client.Collection("members").Documents(fs.ctx)
	defer iter.Stop()
	var members []models.Member
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("GetMembers: %v", err)
			return nil
		}
		var m models.Member
		if err := doc.DataTo(&m); err != nil {
			log.Printf("GetMembers DataTo: %v", err)
			continue
		}
		members = append(members, m)
	}
	return members
}

func (fs *FirestoreStore) GetMember(id int) (models.Member, bool) {
	doc, err := fs.client.Collection("members").Doc(strconv.Itoa(id)).Get(fs.ctx)
	if err != nil {
		return models.Member{}, false
	}
	var m models.Member
	if err := doc.DataTo(&m); err != nil {
		return models.Member{}, false
	}
	return m, true
}

func (fs *FirestoreStore) CreateMember(m models.Member) models.Member {
	m.ID = fs.nextID("members")
	if _, err := fs.client.Collection("members").Doc(strconv.Itoa(m.ID)).Set(fs.ctx, m); err != nil {
		log.Printf("CreateMember: %v", err)
	}
	return m
}

func (fs *FirestoreStore) UpdateMember(id int, m models.Member) (models.Member, bool) {
	m.ID = id
	ref := fs.client.Collection("members").Doc(strconv.Itoa(id))
	doc, err := ref.Get(fs.ctx)
	if err != nil || !doc.Exists() {
		return models.Member{}, false
	}
	if _, err := ref.Set(fs.ctx, m); err != nil {
		log.Printf("UpdateMember: %v", err)
		return models.Member{}, false
	}
	return m, true
}

func (fs *FirestoreStore) DeleteMember(id int) bool {
	ref := fs.client.Collection("members").Doc(strconv.Itoa(id))
	doc, err := ref.Get(fs.ctx)
	if err != nil || !doc.Exists() {
		return false
	}
	if _, err := ref.Delete(fs.ctx); err != nil {
		log.Printf("DeleteMember: %v", err)
		return false
	}
	return true
}

// ── Performances ───────────────────────────────────────────────────────────

func (fs *FirestoreStore) GetPerformances() []models.Performance {
	iter := fs.client.Collection("performances").Documents(fs.ctx)
	defer iter.Stop()
	var perfs []models.Performance
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("GetPerformances: %v", err)
			return nil
		}
		var p models.Performance
		if err := doc.DataTo(&p); err != nil {
			log.Printf("GetPerformances DataTo: %v", err)
			continue
		}
		perfs = append(perfs, p)
	}
	return perfs
}

func (fs *FirestoreStore) GetPerformance(id int) (models.Performance, bool) {
	doc, err := fs.client.Collection("performances").Doc(strconv.Itoa(id)).Get(fs.ctx)
	if err != nil {
		return models.Performance{}, false
	}
	var p models.Performance
	if err := doc.DataTo(&p); err != nil {
		return models.Performance{}, false
	}
	return p, true
}

func (fs *FirestoreStore) CreatePerformance(p models.Performance) models.Performance {
	p.ID = fs.nextID("performances")
	if _, err := fs.client.Collection("performances").Doc(strconv.Itoa(p.ID)).Set(fs.ctx, p); err != nil {
		log.Printf("CreatePerformance: %v", err)
	}
	return p
}

func (fs *FirestoreStore) UpdatePerformance(id int, p models.Performance) (models.Performance, bool) {
	p.ID = id
	ref := fs.client.Collection("performances").Doc(strconv.Itoa(id))
	doc, err := ref.Get(fs.ctx)
	if err != nil || !doc.Exists() {
		return models.Performance{}, false
	}
	if _, err := ref.Set(fs.ctx, p); err != nil {
		log.Printf("UpdatePerformance: %v", err)
		return models.Performance{}, false
	}
	return p, true
}

func (fs *FirestoreStore) DeletePerformance(id int) bool {
	ref := fs.client.Collection("performances").Doc(strconv.Itoa(id))
	doc, err := ref.Get(fs.ctx)
	if err != nil || !doc.Exists() {
		return false
	}
	if _, err := ref.Delete(fs.ctx); err != nil {
		log.Printf("DeletePerformance: %v", err)
		return false
	}
	return true
}

// ── Auth ───────────────────────────────────────────────────────────────────

func (fs *FirestoreStore) FindUserByEmail(email string) (models.User, bool) {
	iter := fs.client.Collection("users").Where("email", "==", email).Limit(1).Documents(fs.ctx)
	defer iter.Stop()
	doc, err := iter.Next()
	if err != nil {
		return models.User{}, false
	}
	var u models.User
	if err := doc.DataTo(&u); err != nil {
		return models.User{}, false
	}
	return u, true
}

// SeedIfEmpty seeds Firestore with initial data if the collections are empty.
// Call this once after creating a FirestoreStore.
func (fs *FirestoreStore) SeedIfEmpty() {
	// Check if plays already exist
	iter := fs.client.Collection("users").Limit(1).Documents(fs.ctx)
	doc, err := iter.Next()
	iter.Stop()
	if err == nil && doc != nil {
		// Data already exists, skip seeding
		return
	}

	log.Println("Seeding Firestore with initial data...")

	// Seed admin user
	adminHash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	adminUser := models.User{ID: 1, Email: "gteatrusanfelix@hotmail.com", Password: string(adminHash), Name: "Admin", Role: "admin"}
	if _, err := fs.client.Collection("users").Doc("1").Set(fs.ctx, adminUser); err != nil {
		log.Printf("seed users: %v", err)
	}

	log.Println("Seeding complete.")
}
