package database

import (
	"database/sql"
	"fmt"
	_ "fmt"

	"github.com/JakeDodd/mtgdataload/models"
)

var SetNotFound = &SetNotFoundError{"Set not found"}

type SetNotFoundError struct {
	Message string
}

func (e *SetNotFoundError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

func GetSetBySetId(setId string, db *sql.DB) (models.MtgSet, error) {
	var set models.MtgSet = models.MtgSet{}

	row := db.QueryRow("SELECT * FROM mtg_set WHERE set_id = $1", setId)

	err := row.Scan(&set.SetId, &set.SetCode, &set.SetName, &set.SetType, &set.SetUri, &set.SetSearchUri, &set.ScryfallSetUri)

	if err != nil {
		if err == sql.ErrNoRows {
			return set, SetNotFound
		}
		return set, fmt.Errorf("GetSetBysetId: %s: %v", setId, err)
	}
	return set, nil
}

func SaveSet(set models.MtgSet, db *sql.DB) error {
	row, err := db.Query("INSERT into mtg_set (set_id, set_code, set_name, set_type, set_uri, set_search_uri, scryfall_set_uri)"+
		"VALUES ($1, $2, $3, $4, $5, $6, $7)", set.SetId, set.SetCode, set.SetName, set.SetType, set.SetUri, set.SetSearchUri, set.ScryfallSetUri)

	if err != nil {
		return err
	}
	row.Close()

	return nil
}
