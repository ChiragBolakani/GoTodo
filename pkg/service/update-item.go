package service

import (
	"encoding/json"
	"go_tutorials/pkg/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (s *Service) UpdateItem(w http.ResponseWriter, r *http.Request) {
	var responseFormat = make(responseFormat)
	// get parameters from URL
	vars := mux.Vars(r)

	session := s.db.GetConn()

	var item models.Item

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		responseFormat["success"] = false
		responseFormat["error"] = err.Error()
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responseFormat)
		return
	}

	item, err = s.GetItemInstance(vars["id"], vars["user_id"])
	if err != nil {
		responseFormat["success"] = false
		responseFormat["error"] = err.Error()
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responseFormat)
		return
	}

	updated_at := time.Now().UTC() //updated_at timestamp

	query := `
		UPDATE items 
		SET
			title=?, 
			description=?, 
			status=?, 
			updated_at=?
		WHERE user_id=? AND id=?	
	`

	err = session.Query(query,
		item.Title,
		item.Description,
		item.Status,
		updated_at,
		item.User_id,
		vars["id"]).Exec()

	if err != nil {
		responseFormat["success"] = false
		responseFormat["error"] = err.Error()
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responseFormat)
		return
	}

	w.WriteHeader(201)
	responseFormat["success"] = true
	responseFormat["data"] = item
	json.NewEncoder(w).Encode(responseFormat)
}
