package models

import "github.com/go-pg/pg/v10"

type Set struct {
	tableName struct{} `pg:"vw_sets"`
	ID        int64    `json:"id"`
	Sport     string   `json:"sport"`
	Brand     string   `json:"brand"`
	Year      int64    `json:"year"`
	Set       string   `json:"set"`
}

func GetSets(db *pg.DB, year string) ([]*Set, error) {
	sets := make([]*Set, 0)
	var err error

	if year == "" {
		err = db.Model(&sets).Select()
	} else {
		err = db.Model(&sets).Where("year = ?", year).Select()
	}

	return sets, err
}
