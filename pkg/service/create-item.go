package service

import (
	"encoding/json"
	"go_tutorials/pkg/models"
	"net/http"
	"time"
)

type responseFormat map[string]any

func (s *Service) CreateItem(w http.ResponseWriter, r *http.Request) {
	var responseFormat = make(responseFormat)

	session := s.db.GetConn()

	var item models.NewItem

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		responseFormat["success"] = false
		responseFormat["error"] = err.Error()
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responseFormat)
		return
	}

	created_at := time.Now().UTC() //created_at timestamp

	query := `
			INSERT INTO items(
				id, 
				user_id, 
				title, 
				description, 
				status, 
				created_at, 
				updated_at
				) 
			VALUES (now(), ?, ?, ?, ?, ?, ?);
		`

	err = session.Query(
		query,
		item.User_id,
		item.Title,
		item.Description,
		item.Status,
		created_at,
		created_at).Exec()

	if err != nil {
		responseFormat["success"] = false
		responseFormat["error"] = err.Error()
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responseFormat)
		return
	}

	w.WriteHeader(200)
	responseFormat["success"] = true
	json.NewEncoder(w).Encode(responseFormat)
}
