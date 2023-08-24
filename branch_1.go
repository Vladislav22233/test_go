package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

const (
	DBHost     = "localhost"
	DBPort     = "5432"
	DBUser     = "postgres"
	DBPassword = "password"
	DBName     = "segments"
)

type Segment struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
}

type User struct {
	ID        int      `json:"id"`
	Segments  []string `json:"segments"`
	CreatedAt string   `json:"created_at"`
}

var db *sql.DB

func main() {
	var err error

	db, err = sql.Open("postgres", getDBConnectionString())

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/segment", createSegment).Methods("POST")
	router.HandleFunc("/segment/{slug}", deleteSegment).Methods("DELETE")
	router.HandleFunc("/user/{id}/segments", addUserSegments).Methods("POST")
	router.HandleFunc("/user/{id}/segments", deleteUserSegments).Methods("DELETE")
	router.HandleFunc("/user/{id}/segments", getUserSegments).Methods("GET")

	log.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func createSegment(w http.ResponseWriter, r *http.Request) {
	var segment Segment
	err := json.NewDecoder(r.Body).Decode(&segment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert segment into database
	insertSegment(segment)

	json.NewEncoder(w).Encode(segment)
}

func deleteSegment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]

	// Delete segment from database
	deleteSegmentFromDB(slug)

	w.WriteHeader(http.StatusOK)
}

func addUserSegments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var segments []Segment
	err := json.NewDecoder(r.Body).Decode(&segments)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add segments to user in database
	addSegmentsToUser(id, segments)

	w.WriteHeader(http.StatusOK)
}

func deleteUserSegments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var segments []Segment
	err := json.NewDecoder(r.Body).Decode(&segments)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete segments from user in database
	deleteSegmentsFromUser(id, segments)

	w.WriteHeader(http.StatusOK)
}

func getUserSegments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// Get segments for user from database
	userSegments := getSegmentsForUser(id)

	json.NewEncoder(w).Encode(userSegments)
}

// Функция для вставки сегмента в базу данных
func insertSegment(segment Segment) {
	_, err := db.Exec("INSERT INTO segments (id, slug) VALUES ($1, $2)", segment.ID, segment.Slug)

	if err != nil {
		log.Fatal(err)
	}
}

// Функция для удаления сегмента из базы данных
func deleteSegmentFromDB(slug string) {
	_, err := db.Exec("DELETE FROM segments WHERE slug = $1", slug)

	if err != nil {
		log.Fatal(err)
	}
}

// Функция для добавления сегментов пользователю в базе данных
func addSegmentsToUser(userID string, segments []Segment) {
	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	for _, segment := range segments {
		_, err := tx.Exec("INSERT INTO user_segments (user_id, segment_id) VALUES ($1, $2)", userID, segment.ID)

		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

// Функция для удаления сегментов пользователя из базы данных
func deleteSegmentsFromUser(userID string, segments []Segment) {
	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	for _, segment := range segments {
		_, err := tx.Exec("DELETE FROM user_segments WHERE user_id = $1 AND segment_id = $2", userID, segment.ID)

		if err != nil {
		tx.Rollback()
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

// Функция для получения сегментов пользователя из базы данных
func getSegmentsForUser(userID string) []Segment {
	rows, err := db.Query("SELECT s.id, s.slug FROM segments s INNER JOIN user_segments us ON s.id = us.segment_id WHERE us.user_id = $1", userID)

	if err != nil {
		log.Fatal(err)
	}

	var userSegments []Segment
	for rows.Next() {
		var segment Segment
		err := rows.Scan(&segment.ID, &segment.Slug)

		if err != nil {
			log.Fatal(err)
		}
		userSegments = append(userSegments, segment)
	}

	return userSegments
}

func getDBConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPassword, DBName)
}
