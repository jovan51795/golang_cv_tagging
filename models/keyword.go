package models

import (
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

	isExist, err := keywordsContains(k)

	if err != nil {
		return err
	}

	if isExist {
		fmt.Println("already exist")
		return errors.New("already exist")
	}

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

func keywordsContains(k *Keyword) (bool, error) {
	keywords, _ := GetAllKeywords()
	return util.Contains(keywords, k), nil
}
