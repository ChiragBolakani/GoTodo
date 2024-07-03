package service

import (
	"encoding/json"
	"go_tutorials/pkg/models"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Service) DeleteItem(w http.ResponseWriter, r *http.Request) {

	var responseFormat = make(responseFormat)
	//get parameters from URL
	vars := mux.Vars(r)

	var (
		err  error
		item models.Item
	)

	item, err = s.GetItemInstance(vars["id"], vars["user_id"])
	if err != nil {
		responseFormat["success"] = false
		responseFormat["error"] = err.Error()
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responseFormat)
		return
	}

	session := s.db.GetConn()

	query := `
		DELETE FROM items 
		WHERE 
			id = ? AND user_id=?`

	err = session.Query(query, vars["id"], vars["user_id"]).Exec()

	if err != nil {
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
