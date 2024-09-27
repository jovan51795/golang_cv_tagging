package models

import (
	"database/sql"
	"errors"
	"fmt"

	"77gsi_mynt.com/cv_tagging/db"
	"77gsi_mynt.com/cv_tagging/util"
)

type Keyword struct {
	Id      int64
	Keyword string
	User_id int64
}

func (k *Keyword) Save() error {
	query := `INSERT INTO keywords(keyword, user_id) VALUES($1, $2)`

	isExist, err := GetKeywordByKey(k.Keyword)

	if err != nil {
		return err
	}

	if isExist != nil {
		fmt.Println("already exist")
		return errors.New("already exist")
	}

	fmt.Println("hi", isExist)
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(k.Keyword, k.User_id)

	if err != nil {
		return err
	}

	return err
}

func GetAllKeywords() ([]Keyword, error) {
	query := `SELECT * FROM keywords`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	var keywords []Keyword
	for rows.Next() {
		var keyword Keyword
		err = rows.Scan(&keyword.Id, &keyword.Keyword, &keyword.User_id)
		if err != nil {
			return nil, err
		}

		keywords = append(keywords, keyword)
	}

	return keywords, nil

}

func GetKeywordByKey(key string) (*Keyword, error) {
	query := `SELECT * FROM keywords WHERE keyword = $1`

	row := db.DB.QueryRow(query, key)

	var keyword Keyword
	err := row.Scan(&keyword.Id, &keyword.Keyword, &keyword.User_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &keyword, nil
}

func keywordsContains(k *Keyword) (bool, error) {
	keywords, _ := GetAllKeywords()
	return util.Contains(keywords, k), nil
}
