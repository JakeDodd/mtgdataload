package database

import (
	"database/sql"
	"fmt"
	_ "fmt"

	"github.com/JakeDodd/mtgdataload/models"
)

var PrintNotFound = &PrintNotFoundError{"Print not found"}

type PrintNotFoundError struct {
	Message string
}

func (e *PrintNotFoundError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

func GetPrintByCardNameAndSetId(card_name string, set_id string, db *sql.DB) (models.Prints, error) {
	var print models.Prints = models.Prints{}

	row := db.QueryRow("SELECT * FROM prints WHERE card_name = $1 and set_id = $2", card_name, set_id)

	err := row.Scan(&print.MtgoId, &print.MtgoFoilId, &print.ArenaId, &print.ArenaId, &print.ScryfallUri, &print.RulingsUri, &print.TcgplayerId,
		&print.TcgplayerEtchedId, &print.ReleasedAt, &print.Oversized, &print.SetId, &print.OracleText, &print.CollectorNumber, &print.Digital, &print.OldschoolF,
		&print.Rarity, &print.CardBackId, &print.Artist, &print.IllustrationId, &print.BorderColor, &print.Frame, &print.FullArt, &print.Textless, &print.Booster,
		&print.StorySpotlight, &print.GathererUri, &print.TcgArticlesUri, &print.TcgDecksUri, &print.EdhrecUri, &print.TcgBuyUri, &print.CardmarketBuyUri, &print.CardhoarderBuyUri,
		&print.OracleId, &print.CardName, &print.PrintsSearchUri, &print.FlavorName, &print.SecurityStamp, &print.PreviewedAt, &print.PreviewUri,
		&print.PreviewSource, &print.ContentWarning)

	if err != nil {
		if err == sql.ErrNoRows {
			return print, PrintNotFound
		}
		return print, fmt.Errorf("GetPrintByCardName: %s: %v", card_name, err)
	}

	rows, err := db.Query("SELECT attraction_light FROM print_attraction_light WHERE card_name = $1 and set_id = $2", card_name, set_id)
	var attraction_lights []int

	if err == nil {
		for rows.Next() {
			var attraction_light int
			err = rows.Scan(&attraction_light)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName: %s: %v", card_name, err)
			}
			attraction_lights = append(attraction_lights, attraction_light)
		}
	}
	print.AttractionLights = attraction_lights
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT multiverse_id FROM print_multiverse_id WHERE card_name = $1 and set_id = $2", card_name, set_id)
	var multiverse_ids []int

	if err == nil {
		for rows.Next() {
			var multiverse_id int
			err = rows.Scan(&multiverse_id)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName: %s: %v", card_name, err)
			}
			multiverse_ids = append(multiverse_ids, multiverse_id)
		}
	}
	print.MultiverseIds = multiverse_ids
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT game FROM print_games WHERE card_name = $1 and set_id = $2", card_name, set_id)
	var games []string

	if err == nil {
		for rows.Next() {
			var game string
			err = rows.Scan(&game)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName: %s: %v", card_name, err)
			}
			games = append(games, game)
		}
	}
	print.Games = games
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT border_effect FROM print_border_effect WHERE card_name = $1 and set_id = $2", card_name, set_id)
	var border_effects []string

	if err == nil {
		for rows.Next() {
			var border_effect string
			err = rows.Scan(&border_effect)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardName: %s: %v", card_name, err)
			}
			border_effects = append(border_effects, border_effect)
		}
	}
	print.BorderEffects = border_effects
	if rows != nil {
		rows.Close()
	}

	rows, err = db.Query("SELECT related_id FROM print_related WHERE print_card_name = $1 and print_set_id = $2", card_name, set_id)

	var related_cards []models.Related

	if err == nil {
		for rows.Next() {
			var related_id string
			err = rows.Scan(&related_id)
			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardNameAndSetId: %s: %v", card_name, set_id, err)
			}
			var related models.Related

			related_row := db.QueryRow("SELECT * FROM related WHERE card_name = $1", related_id)
			err = related_row.Scan(&related.Object, &related.Id, &related.Component, &related.Name, &related.TypeLine, &related.Uri)

			if err != nil {
				if err == sql.ErrNoRows {
					break
				}
				return print, fmt.Errorf("GetPrintByCardNameAndSetId: %s: %v", card_name, set_id, err)

			}
			related_cards = append(related_cards, related)
		}
		print.Related = related_cards
		if rows != nil {
			rows.Close()
		}
	}

	return print, nil
}

func SavePrint(print models.Prints, db *sql.DB) error {
	row, err := db.Query("INSERT into prints (mtgo_id, mtgo_foil_id, arena_id, scryfall_uri, rulings_uri, tcgplayer_id, tcgplayer_etched_id, released_at, oversized, set_id, oracle_text,"+
		"collector_number, digital, rarity, oldschool_f, card_back_id, artist, illustration_id, border_color, frame, full_art, textless, booster, story_spotlight, gatherer_uri, tcg_articles_uri,"+
		"tcg_decks_uri, edhrec_uri, tcg_buy_uri, cardmarket_buy_uri, cardhoarder_buy_uri, oracle_id, card_name, prints_search_uri, flavor_name, security_stamp, previewed_at, previewed_source_uri,"+
		"previewsource, content_warning)"+
		"VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37,$38,$39,$40)",
		print.MtgoId, print.MtgoFoilId, print.ArenaId, print.ScryfallUri, print.RulingsUri, print.TcgplayerId, print.TcgplayerEtchedId, print.ReleasedAt, print.Oversized, print.SetId, print.OracleText,
		print.CollectorNumber, print.Digital, print.Rarity, print.OldschoolF, print.CardBackId, print.Artist, print.IllustrationId, print.BorderColor, print.Frame, print.FullArt, print.Textless, print.Booster,
		print.StorySpotlight, print.GathererUri, print.TcgArticlesUri, print.TcgDecksUri, print.EdhrecUri, print.TcgBuyUri, print.CardmarketBuyUri, print.CardhoarderBuyUri, print.OracleId, print.CardName,
		print.PrintsSearchUri, print.FlavorName, print.SecurityStamp, print.PreviewedAt, print.PreviewUri, print.PreviewSource, print.ContentWarning)

	if err != nil {
		return err
	}
	row.Close()

	for i := 0; i < len(print.AttractionLights); i++ {
		row, err = db.Query("INSERT INTO print_attraction_lights (card_name, set_id, attraction_light) VALUES ($1, $2, $3)", print.CardName, print.SetId, print.AttractionLights[i])
		if err != nil {
			return err
		}
		row.Close()
	}

	for i := 0; i < len(print.MultiverseIds); i++ {
		row, err = db.Query("INSERT INTO print_multiverse_id (card_name, set_id, multiverse_id) VALUES ($1, $2, $3)", print.CardName, print.SetId, print.MultiverseIds[i])
		if err != nil {
			return err
		}
		row.Close()
	}

	for i := 0; i < len(print.Games); i++ {
		row, err = db.Query("INSERT INTO print_game (card_name, set_id, game) VALUES ($1, $2, $3)", print.CardName, print.SetId, print.Games[i])
		if err != nil {
			return err
		}
		row.Close()
	}

	for i := 0; i < len(print.BorderEffects); i++ {
		row, err = db.Query("INSERT INTO print_border_effect (card_name, set_id, border_effect) VALUES ($1, $2, $3)", print.CardName, print.SetId, print.BorderEffects[i])
		if err != nil {
			return err
		}
		row.Close()
	}

	for i := 0; i < len(print.Related); i++ {
		related_card := print.Related[i]
		related_row := db.QueryRow("SELECT * FROM related WHERE related_id = $1", related_card.Id)
		if err := related_row.Scan(); err == sql.ErrNoRows {
			row, err = db.Query("INSERT INTO related (object_parts, id, component, card_name, type_line, uri) VALUES ($1, $2, $3, $4, $5, $6)", related_card.Object, related_card.Id, related_card.Component, related_card.Name, related_card.TypeLine, related_card.Uri)
			if err != nil {
				return err
			}
			row.Close()
			row, err = db.Query("INSERT INTO print_related (print_card_name, set_id, related_id) VALUES ($1, $2, $3)", print.CardName, print.SetId, related_card)
			if err != nil {
				return err
			}
			row.Close()
		}

	}
	return nil
}
