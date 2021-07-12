package db

import (
	"fmt"

	"github.com/hawwwdi/qanari/model"
)

func (db *DB) Like(a *model.Ava) error {
	query := `call likeAva_procedure(?)`
	res, err := db.SqlDB.Exec(query, a.ID)
	rows, err := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("Ava not found")
	}
	return err
}

func (db *DB) GetLikers(a *model.Ava) ([]*model.User, error) {
	query := `call getLikers(?)`
	rows, err := db.SqlDB.Query(query, a.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	likers := make([]*model.User, 0)
	for rows.Next() {
		var curr model.User
		if err := rows.Scan(&curr.ID, &curr.Username); err != nil {
			return nil, err
		}
		likers = append(likers, &curr)
	}
	return likers, nil
}

func (db *DB) GetLikesCount(a *model.Ava) (int, error) {
	query := `call getLikesCount(?)`
	rows, err := db.SqlDB.Query(query, a.ID)
	if err != nil {
		return -1, err
	}
	var count int
	if err := rows.Scan(&count); err != nil {
		return -1, err
	}
	return count, nil
}
