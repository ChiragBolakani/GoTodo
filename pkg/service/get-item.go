package service

import (
	"encoding/json"
	"go_tutorials/pkg/models"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Service) GetItem(w http.ResponseWriter, r *http.Request) {
	var responseFormat = make(responseFormat)
	var err error

	// get parameters from URL
	user_id := mux.Vars(r)["user_id"]
	id := mux.Vars(r)["id"]

	query := `
		SELECT
			id,
			user_id,
			title,
			description,
			status,
			created_at,
			updated_at
		FROM items
		WHERE user_id = ? 
		AND id = ?
	`
	query += ` ORDER BY id DESC ALLOW FILTERING;`

	session := s.db.GetConn()

	scanner := session.Query(query, user_id, id).Iter().Scanner()
	scanner.Next()

	var (
		item models.Item
	)

	err = scanner.Scan(
		&item.Id,
		&item.User_id,
		&item.Title,
		&item.Description,
		&item.Status,
		&item.Created_at,
		&item.Updated_at,
	)

	if err != nil {
		responseFormat["success"] = false
		responseFormat["error"] = err.Error()
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responseFormat)
		return
	}

	if err := scanner.Err(); err != nil {
		responseFormat["success"] = false
		responseFormat["error"] = err.Error()
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responseFormat)
		return
	}

	w.WriteHeader(200)
	responseFormat["success"] = true
	responseFormat["data"] = item
	json.NewEncoder(w).Encode(responseFormat)
}
