package service

import (
	"errors"
	"go_tutorials/pkg/models"
)

func (s *Service) GetItemInstance(id string, user_id string) (models.Item, error) {
	var (
		item models.Item
		err  error
	)
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

	iter := session.Query(query, user_id, id).Iter()
	if iter.NumRows() == 0 {
		return item, errors.New("no record found")
	}

	scanner := iter.Scanner()
	scanner.Next()

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
		return item, err
	}

	if err := scanner.Err(); err != nil {
		return item, err
	}

	return item, err
}
