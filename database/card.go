package database

import (
	"database/sql"
	"fmt"
	_ "fmt"
	"strconv"

	"github.com/JakeDodd/mtgdataload/models"
)

var CardNotFound = &CardNotFoundError{"Card not found"}

type CardNotFoundError struct {
	Message string
}

func (e *CardNotFoundError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

func GetCardByOracleIdAndName(oracleId string, name string, db *sql.DB) (models.Cards, error) {
	var card models.Cards = models.Cards{}
	row := db.QueryRow("SELECT * FROM cards WHERE oracle_id = $1 and card_name = $2", oracleId, name)

	err := row.Scan(&card.OracleId, &card.Object, &card.CardName, &card.Layout, &card.ManaCost, &card.Cmc, &card.TypeLine, &card.Power, &card.Toughness, &card.Reserved, &card.StandardF, &card.FutureF, &card.HistoricF, &card.TimelessF, &card.GladiatorF, &card.PioneerF, &card.ExplorerF, &card.ModernF, &card.LegacyF, &card.PauperF, &card.VintageF, &card.PennyF, &card.CommanderF, &card.OathbreakerF, &card.StandardbrawlF, &card.BrawlF, &card.AlchemyF, &card.PaupercommanderF, &card.DuelF, &card.PremodernF, &card.PredhF, &card.Defense, &card.Loyalty, &card.EdhrecRank, &card.HandModifier, &card.LifeModifier)

	if err != nil {
		if err == sql.ErrNoRows {
			return card, CardNotFound
		}
		return card, fmt.Errorf("GetCardByoracleId: %s: %v", oracleId, err)
	}
	rows, err := db.Query("SELECT color FROM card_color WHERE oracle_id = $1 and card_name = $2", oracleId, name)
	var colors []string

	if err == nil {
		for rows.Next() {
			var color string
			err = rows.Scan(&color)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return card, fmt.Errorf("GetCardByoracleId: %s: %v", oracleId, err)
			}
			colors = append(colors, color)
		}
	}
	card.Colors = colors
	if rows != nil {
		rows.Close()
	}
	rows, err = db.Query("SELECT color FROM card_color_identity WHERE oracle_id = $1 and card_name = $2", oracleId, name)
	var color_identities []string

	if err == nil {
		for rows.Next() {
			var color string
			err = rows.Scan(&color)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return card, fmt.Errorf("GetCardByoracleId: %s: %v", oracleId, err)
			}
			color_identities = append(color_identities, color)
		}
	}
	card.ColorIdentity = color_identities
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT color FROM card_produced_mana WHERE oracle_id = $1 and card_name = $2", oracleId, name)
	var produced_mana []string

	if err == nil {
		for rows.Next() {
			var color string
			err = rows.Scan(&color)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return card, fmt.Errorf("GetCardByoracleId: %s: %v", oracleId, err)
			}
			produced_mana = append(produced_mana, color)
		}
	}
	card.ProducedMana = produced_mana
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT color FROM card_color_indicator WHERE oracle_id = $1 and card_name = $2", oracleId, name)
	var color_indicator []string

	if err == nil {
		for rows.Next() {
			var color string
			err = rows.Scan(&color)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return card, fmt.Errorf("GetCardByoracleId: %s: %v", oracleId, err)
			}
			color_indicator = append(color_indicator, color)
		}
	}
	card.ColorIndicator = color_indicator
	if rows != nil {
		rows.Close()
	}
	/*
		rows, err = db.Query("SELECT attraction_light FROM card_attraction_lights WHERE card_name = $1", name)
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
		if rows != nil {
			rows.Close()
		}
	*/
	rows, err = db.Query("SELECT keyword FROM card_keyword WHERE oracle_id = $1 and card_name = $2", oracleId, name)
	var keywords []string

	if err == nil {
		for rows.Next() {
			var keyword string
			err = rows.Scan(&keyword)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return card, fmt.Errorf("GetCardByoracleId: %s: %v", oracleId, err)
			}
			keywords = append(keywords, keyword)
		}
	}
	card.Keywords = keywords
	if rows != nil {
		rows.Close()
	}

	/*
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
				if cfs_rows != nil {
					cfs_rows.Close()
				}

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
				if cfs_rows != nil {
					cfs_rows.Close()
				}

				card_faces = append(card_faces, cf)
			}
		}
		card.CardFaces = card_faces
		if rows != nil {
			rows.Close()
		}
	*/

	return card, nil

}

func SaveCard(card models.Cards, db *sql.DB) error {

	row, err := db.Query("INSERT into cards (oracle_id ,\"object\", card_name, layout, mana_cost, cmc, type_line, power,"+
		" toughness, reserved, standard_f, future_f, historic_f, timeless_f, gladiator_f, pioneer_f, explorer_f, modern_f, legacy_f, pauper_f, vintage_f, penny_f, commander_f, oathbreaker_f, standardbrawl_f, brawl_f, alchemy_f, paupercommander_f, duel_f, premodern_f, predh_f, defense, loyalty, edhrec_rank, "+
		"hand_modifier, life_modifier) "+
		"VALUES ($1,$2,$3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36)",
		card.OracleId, card.Object, card.CardName, card.Layout, card.ManaCost, strconv.FormatFloat(card.Cmc, 'f', -1, 64), card.TypeLine, card.Power,
		card.Toughness, card.Reserved,
		card.StandardF, card.FutureF, card.HistoricF, card.TimelessF, card.GladiatorF, card.PioneerF,
		card.ExplorerF, card.ModernF, card.LegacyF,
		card.PauperF, card.VintageF, card.PennyF, card.CommanderF, card.OathbreakerF,
		card.StandardbrawlF, card.BrawlF, card.AlchemyF, card.PaupercommanderF, card.DuelF,
		card.PremodernF, card.PredhF, card.Defense, card.Loyalty,
		card.EdhrecRank, card.HandModifier, card.LifeModifier)

	if err != nil {
		return err
	}
	row.Close()

	for i := 0; i < len(card.Colors); i++ {
		row, err = db.Query("INSERT INTO card_color (card_name, oracle_id, color) VALUES ($1, $2, $3)", card.CardName, card.OracleId, card.Colors[i])
		if err != nil {
			return err
		}
		row.Close()
	}
	for i := 0; i < len(card.ColorIdentity); i++ {
		row, err = db.Query("INSERT INTO card_color_identity (card_name, oracle_id, color) VALUES ($1, $2, $3)", card.CardName, card.OracleId, card.ColorIdentity[i])
		if err != nil {
			return err
		}
		row.Close()
	}
	for i := 0; i < len(card.ProducedMana); i++ {
		row, err = db.Query("INSERT INTO card_produced_mana (card_name, oracle_id, color) VALUES ($1, $2, $3)", card.CardName, card.OracleId, card.ProducedMana[i])
		if err != nil {
			return err
		}
		row.Close()
	}
	for i := 0; i < len(card.ColorIndicator); i++ {
		row, err = db.Query("INSERT INTO card_color_indicator (card_name, oracle_id, color) VALUES ($1, $2, $3)", card.CardName, card.OracleId, card.ColorIndicator[i])
		if err != nil {
			return err
		}
		row.Close()
	}
	/*
		for i := 0; i < len(card.AttractionLights); i++ {
			row, err = db.Query("INSERT INTO card_attraction_light (card_name, attraction_light) VALUES ($1, $2)", card.CardName, card.AttractionLights[i])
			if err != nil {
				return err
			}
			row.Close()
		}
	*/
	for i := 0; i < len(card.Keywords); i++ {
		row, err = db.Query("INSERT INTO card_keyword (card_name, oracle_id, keyword) VALUES ($1, $2, $3)", card.CardName, card.OracleId, card.Keywords[i])
		if err != nil {
			return err
		}
		row.Close()
	}

	/*
		for i := 0; i < len(card.CardFaces); i++ {
			cf := card.CardFaces[i]
			r := db.QueryRow("SELECT * FROM card_faces WHERE card_name = $1", cf.Name)
			if err := r.Scan(); err == sql.ErrNoRows {
				row, err = db.Query("INSERT INTO card_faces (card_name, artist, artist_id, cmc, defense, flavor_text, illustration_id, png_uri, boarder_crop_uri, art_crop_uri, large_uri, normal_uri, small_uri, layout, loyalty, mana_cost, object_type, oracle_id, oracle_text, power, printed_name, printed_text, printed_type_line, toughness, type_line, watermark) VALUES ($1, $2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26)", cf.Name, cf.Artist, cf.ArtistId, cf.Cmc, cf.Defense, cf.FlavorText, cf.IllustrationId, cf.PngUri, cf.BoarderCropUri, cf.ArtCropUri, cf.LargeUri, cf.NormalUri, cf.SmallUri, cf.Layout, cf.Loyalty, cf.ManaCost, cf.Object, cf.OracleId, cf.OracleText, cf.Power, cf.PrintedName, cf.PrintedText, cf.PrintedTypeLine, cf.Toughness, cf.TypeLine, cf.Watermark)
				if err != nil {
					return err
				}
				row.Close()
				row, err = db.Query("INSERT INTO card_card_faces (card_card_name, card_faces_card_name) VALUES ($1, $2)", card.CardName, cf.Name)
				if err != nil {
					return err
				}
				row.Close()
				for i := 0; i < len(cf.Colors); i++ {
					row, err = db.Query("INSERT INTO card_faces_color (card_name, color) VALUES ($1, $2)", cf.Name, cf.Colors[i])
					if err != nil {
						return err
					}
					row.Close()
				}
				for i := 0; i < len(cf.ColorIndicator); i++ {
					row, err = db.Query("INSERT INTO card_faces_color_indicator (card_name, color) VALUES ($1, $2)", cf.Name, cf.ColorIndicator[i])
					if err != nil {
						return err
					}
					row.Close()
				}
			}
		}
	*/
	return nil
}
