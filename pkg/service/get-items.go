package service

import (
	"encoding/hex"
	"encoding/json"
	"go_tutorials/pkg/config"
	"go_tutorials/pkg/models"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Service) GetItems(w http.ResponseWriter, r *http.Request) {

	var (
		err            error
		items          []models.Item
		args           []any
		responseFormat = make(responseFormat)
	)

	//get parameters from URL
	user_id := mux.Vars(r)["user_id"]
	status := r.URL.Query().Get("status")

	// mapping pending/completed
	// pending : status = false
	// completed : status = true
	statusBool := false
	if status == "pending" {
		statusBool = false
	} else if status == "completed" {
		statusBool = true
	}

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
	`

	args = append(args, user_id)

	if status != "" {
		query += `AND status = ?`
		args = append(args, statusBool)
	}

	query += ` ORDER BY id DESC ALLOW FILTERING;`

	session := s.db.GetConn()

	var pageState []byte

	// if URL contains page_state query param
	if r.URL.Query().Get("page_state") != "" {

		//get page_state from URL
		page_state_from_url := r.URL.Query().Get("page_state")

		//decode the hexstring from string to []byte
		pageState, err = hex.DecodeString(page_state_from_url)
		if err != nil {
			pageState = nil
		}
	}

	iter := session.Query(query, args...).PageSize(config.PageSize).PageState(pageState).Iter()

	// get next pageState
	nextPageState := iter.PageState()
	//Encode nextPageState to hexstring for response
	nextPageStateString := hex.EncodeToString(nextPageState)

	scanner := iter.Scanner()
	for scanner.Next() {
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
			return
		}

		items = append(items, item)
	}

	if err := scanner.Err(); err != nil {
		responseFormat["success"] = false
		responseFormat["error"] = err.Error()
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
	responseFormat["success"] = true
	responseFormat["data"] = items
	responseFormat["next_page_state"] = nextPageStateString //add nextPageState to response
	json.NewEncoder(w).Encode(responseFormat)
}
