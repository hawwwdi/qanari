package db

import "github.com/hawwwdi/qanari/model"

func (db *DB) PostAva(a *model.Ava) error {
	query := `call postAVA_procedure(?)`
	_, err := db.SqlDB.Exec(query, a.Content)
	return err
}

func (db *DB) PostComment(a *model.Ava) error {
	query := `call postComment_procedure(?, ?)`
	_, err := db.SqlDB.Exec(query, a.Content, a.ReplyTo)
	return err
}

func (db *DB) GetTimeLine() ([]*model.Ava, error) {
	query := `call getTimeLine_procedure()`
	rows, err := db.SqlDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	timeLine := make([]*model.Ava, 0)
	for rows.Next() {
		var curr model.Ava
		if err := rows.Scan(&curr.ID, &curr.Publisher, &curr.Content, &curr.ReplyTo, &curr.PublishTime); err != nil {
			return nil, err
		}
		timeLine = append(timeLine, &curr)
	}
	return timeLine, nil
}

func (db *DB) GetComments(a *model.Ava) ([]*model.Ava, error) {
	query := `call checkComments_procedure(?)`
	rows, err := db.SqlDB.Query(query, a.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	timeLine := make([]*model.Ava, 0)
	for rows.Next() {
		var curr model.Ava
		if err := rows.Scan(&curr.ID, &curr.Publisher, &curr.Content, &curr.ReplyTo, &curr.PublishTime); err != nil {
			return nil, err
		}
		timeLine = append(timeLine, &curr)
	}
	return timeLine, nil
}

func (db *DB) GetPersonalAvas() ([]*model.Ava, error) {
	query := `call getPersonalAvas()`
	rows, err := db.SqlDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	timeLine := make([]*model.Ava, 0)
	for rows.Next() {
		var curr model.Ava
		if err := rows.Scan(&curr.ID, &curr.Publisher, &curr.Content, &curr.ReplyTo, &curr.PublishTime); err != nil {
			return nil, err
		}
		timeLine = append(timeLine, &curr)
	}
	return timeLine, nil
}

func (db *DB) GetUserAvas(u *model.User) ([]*model.Ava, error) {
	query := `call getUserAvas_procedure(?)`
	rows, err := db.SqlDB.Query(query, u.Username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	timeLine := make([]*model.Ava, 0)
	for rows.Next() {
		var curr model.Ava
		if err := rows.Scan(&curr.ID, &curr.Publisher, &curr.Content, &curr.ReplyTo, &curr.PublishTime); err != nil {
			return nil, err
		}
		timeLine = append(timeLine, &curr)
	}
	return timeLine, nil
}

func (db *DB) GetAvasByTags(tag string) ([]*model.Ava, error) {
	query := `call getPostByTag_procedure(?)`
	rows, err := db.SqlDB.Query(query, tag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	timeLine := make([]*model.Ava, 0)
	for rows.Next() {
		var curr model.Ava
		if err := rows.Scan(&curr.ID, &curr.Publisher, &curr.Content, &curr.ReplyTo, &curr.PublishTime); err != nil {
			return nil, err
		}
		timeLine = append(timeLine, &curr)
	}
	return timeLine, nil
}

func (db *DB) GetAvasByLikes() ([]*model.Ava, error) {
	query := `call getAVAsByLikes()`
	rows, err := db.SqlDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	timeLine := make([]*model.Ava, 0)
	for rows.Next() {
		var curr model.Ava
		if err := rows.Scan(&curr.ID, &curr.Content, &curr.Likes); err != nil {
			return nil, err
		}
		timeLine = append(timeLine, &curr)
	}
	return timeLine, nil

}