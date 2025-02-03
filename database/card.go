package database

import (
	"database/sql"
	"fmt"
	_ "fmt"

	"github.com/JakeDodd/mtgdataload/models"
)

var CardNotFound = &CardNotFoundError{"Card not found"}

type CardNotFoundError struct {
	Message string
}

func (e *CardNotFoundError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

func GetCardByName(name string, db *sql.DB) (models.Cards, error) {
	var card models.Cards = models.Cards{}
	row := db.QueryRow("SELECT * FROM cards WHERE card_name = $1", name)

	err := row.Scan(&card.Object, &card.OracleId, &card.CardName, &card.ScryfallUri, &card.Layout, &card.ManaCost, &card.Cmc, &card.TypeLine, &card.OracleText, &card.Power, &card.Toughness, &card.Reserved, &card.RulingsUri, &card.StandardF, &card.FutureF, &card.HistoricF, &card.TimelessF, &card.GladiatorF, &card.PioneerF, &card.ExplorerF, &card.ModernF, &card.LegacyF, &card.PauperF, &card.VintageF, &card.PennyF, &card.CommanderF, &card.OathbreakerF, &card.StandardbrawlF, &card.BrawlF, &card.AlchemyF, &card.PaupercommanderF, &card.DuelF, &card.OldschoolF, &card.PremodernF, &card.PredhF, &card.Defense, &card.Loyalty, &card.EdhrecRank, &card.HandModifier, &card.LifeModifier, &card.ContentWarning)

	if err != nil {
		if err == sql.ErrNoRows {
			return card, CardNotFound
		}
		return card, fmt.Errorf("GetCardByName: %s: %v", name, err)
	}

	rows, err := db.Query("SELECT color FROM card_color WHERE card_name = $1", name)
	defer rows.Close()
	var colors []string

	if err == nil {
		for rows.Next() {
			var color string
			err = rows.Scan(&color)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return card, fmt.Errorf("GetCardByName: %s: %v", name, err)
			}
			colors = append(colors, color)
		}
	}
	card.Colors = colors

	rows, err = db.Query("SELECT color FROM card_color_identity WHERE card_name = $1", name)
	var color_identities []string

	if err == nil {
		for rows.Next() {
			var color string
			err = rows.Scan(&color)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return card, fmt.Errorf("GetCardByName: %s: %v", name, err)
			}
			color_identities = append(color_identities, color)
		}
	}
	card.ColorIdentity = color_identities

	rows, err = db.Query("SELECT color FROM card_produced_mana WHERE card_name = $1", name)
	var produced_mana []string

	if err == nil {
		for rows.Next() {
			var color string
			err = rows.Scan(&color)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return card, fmt.Errorf("GetCardByName: %s: %v", name, err)
			}
			produced_mana = append(produced_mana, color)
		}
	}
	card.ProducedMana = produced_mana

	rows, err = db.Query("SELECT color FROM card_color_indicator WHERE card_name = $1", name)
	var color_indicator []string

	if err == nil {
		for rows.Next() {
			var color string
			err = rows.Scan(&color)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return card, fmt.Errorf("GetCardByName: %s: %v", name, err)
			}
			color_indicator = append(color_indicator, color)
		}
	}
	card.ColorIndicator = color_indicator

	rows, err = db.Query("SELECT color FROM card_attraction_lights WHERE card_name = $1", name)
	var attraction_lights []int

	if err == nil {
		for rows.Next() {
			var light int
			err = rows.Scan(&light)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return card, fmt.Errorf("GetCardByName: %s: %v", name, err)
			}
			attraction_lights = append(attraction_lights, light)
		}
	}
	card.AttractionLights = attraction_lights

	rows, err = db.Query("SELECT color FROM card_keyword WHERE card_name = $1", name)
	var keywords []string

	if err == nil {
		for rows.Next() {
			var keyword string
			err = rows.Scan(&keyword)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return card, fmt.Errorf("GetCardByName: %s: %v", name, err)
			}
			keywords = append(keywords, keyword)
		}
	}
	card.Keywords = keywords

	rows, err = db.Query("SELECT card_faces_card_name FROM card_card_faces WHERE card_card_name = $1", name)

	var card_faces []models.CardFaces

	if err == nil {
		for rows.Next() {
			var cf_name string
			err = rows.Scan(&cf_name)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return card, fmt.Errorf("GetCardByName: %s: %v", name, err)
			}
			var cf models.CardFaces

			cfs_row := db.QueryRow("SELECT * FROM card_faces WHERE card_name = $1", cf_name)
			err = cfs_row.Scan(&cf.Name, &cf.Artist, &cf.ArtistId, &cf.Cmc, &cf.Defense, &cf.FlavorText, &cf.IllustrationId, &cf.PngUri, &cf.BoarderCropUri, &cf.ArtCropUri, &cf.LargeUri, &cf.NormalUri, &cf.SmallUri, &cf.Layout, &cf.Loyalty, &cf.ManaCost, &cf.Object, &cf.OracleId, &cf.OracleText, &cf.Power, &cf.PrintedName, &cf.PrintedText, &cf.PrintedTypeLine, &cf.Toughness, &cf.TypeLine, &cf.Watermark)

			if err != nil {
				if err == sql.ErrNoRows {
					break
					//TODO: Maybe we should be returning an error here because we found a card face name in the join table, but no corresponding card face record
				}
				return card, fmt.Errorf("GetCardByName: %s: %v", name, err)
			}

			cfs_rows, err := db.Query("SELECT color FROM card_faces_colors WHERE card_name = $1", cf_name)
			defer cfs_rows.Close()

			var cf_colors []string
			if err == nil {
				for cfs_rows.Next() {
					var cf_color string
					err = rows.Scan(&cf_color)
					if err != nil {
						if err == sql.ErrNoRows {
							break
						}
						return card, fmt.Errorf("GetCardByName: %s: %v", name, err)
					}
					cf_colors = append(cf_colors, cf_color)
				}
			}
			cf.Colors = cf_colors

			cfs_rows, err = db.Query("SELECT color FROM card_faces_color_identity WHERE card_name = $1", cf_name)

			var cf_color_indicator []string
			if err == nil {
				for cfs_rows.Next() {
					var cf_color string
					err = rows.Scan(&cf_color)
					if err != nil {
						if err == sql.ErrNoRows {
							break
						}
						return card, fmt.Errorf("GetCardByName: %s: %v", name, err)
					}
					cf_color_indicator = append(cf_color_indicator, cf_color)
				}
			}
			cf.ColorIndicator = color_indicator

			card_faces = append(card_faces, cf)
		}
	}
	card.CardFaces = card_faces

	return card, nil

}

func getArrayFromRelation[T string | int](table string, selectColumn string, filterColumn string, filter string, db *sql.DB, arr *[]T) error {
	rows, _ := db.Query("SELECT $1 FROM $2 WHERE $3 = $4", selectColumn, table, filterColumn, filter)
	var ret []T

	for rows.Next() {
		var val T
		err := rows.Scan(&val)
		if err != nil {
			if err == sql.ErrNoRows {
				*arr = ret
				break
			}
			return fmt.Errorf("getArrayFromRelation: %v", err)
		}
		ret = append(ret, val)
	}
	*arr = ret
	return nil
}
