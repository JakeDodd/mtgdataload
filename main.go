package main

import (
	//"fmt"
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/JakeDodd/mtgdataload/database"
	"github.com/JakeDodd/mtgdataload/models"

	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "mtgdataload"
)

// TODO
// Create struct for original json card table: FileCard
// For each created table we need a matching struct: Card, Printcard, Langcard, etc
// Make new Folder in dataload project called models - models.go (package models at top) - all structs go in here
func array2String(arr []string) string {
	var result string
	for i := 0; i < len(arr); i++ {
		result += arr[i]
		if i < len(arr)-1 {
			result += ","
		}
	}
	return result
}

func intArray2String(arr []int) string {
	var result string
	for i := 0; i < len(arr); i++ {
		result += strconv.Itoa(arr[i])
		if i < len(arr)-1 {
			result += ","
		}
	}
	return result
}

func ternary(b bool, t string, f string) string {
	if b {
		return t
	} else {
		return f
	}
}

func ternary_f(s string) bool {
	if s == "legal" {
		return true
	} else {
		return false
	}
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var card models.FileCard
	file, err := os.Open("all-cards-20250131102123.json")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	start := time.Now()
	var lines, inserts int
	for true {
		line, err := reader.ReadString('\n')
		if err != nil {

			break
		}
		//line starts with [ and ends with ], we dont want to unmarshal the first and last line
		if line != "[\n" && line != "]" {
			//removing new line character, remove the comma in all but the last line, theres no comma in last line
			if line[len(line)-2:] == ",\n" {
				line = line[:len(line)-2]
			} else {
				line = line[:len(line)-1]
			}

			err = json.Unmarshal([]byte(line), &card)
			if err != nil {
				log.Println(card.Name)
				log.Fatal(err)
			}

			//use card object to build all insert statements for card object and run those insert statements
			// if card.Name == "Balloon Stand" {
			// 	log.Print(card.AttractionLights)
			// }
			_, err := database.GetCardByName(card.Name, db)

			if err != nil && err != database.CardNotFound {
				log.Fatal(err)
			}

			if err == database.CardNotFound {
				rows, err := db.Query("INSERT into cards (\"object\", oracle_id, card_name, scryfall_uri, layout, mana_cost, cmc, type_line, oracle_text, power,"+
					" toughness, reserved, rulings_uri, standard_f, future_f, historic_f, timeless_f, gladiator_f, pioneer_f, explorer_f, modern_f, legacy_f, pauper_f, vintage_f, penny_f, commander_f, oathbreaker_f, standardbrawl_f, brawl_f, alchemy_f, paupercommander_f, duel_f, oldschool_f, premodern_f, predh_f, defense, loyalty, edhrec_rank, "+
					"hand_modifier, life_modifier, content_warning) "+
					"VALUES ($1,$2,$3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37,$38,$39,$40,$41)",
					card.Object, card.OracleId, card.Name, card.ScryfallUri, card.Layout, card.ManaCost, strconv.FormatFloat(card.Cmc, 'f', -1, 64), card.TypeLine, card.OracleText, card.Power,
					card.Toughness, card.Reserved, card.RulingsUri,
					ternary_f(card.Legalities.Standard), ternary_f(card.Legalities.Future), ternary_f(card.Legalities.Historic), ternary_f(card.Legalities.Timeless), ternary_f(card.Legalities.Gladiator), ternary_f(card.Legalities.Pioneer),
					ternary_f(card.Legalities.Explorer), ternary_f(card.Legalities.Modern), ternary_f(card.Legalities.Legacy),
					ternary_f(card.Legalities.Pauper), ternary_f(card.Legalities.Vintage), ternary_f(card.Legalities.Penny), ternary_f(card.Legalities.Commander), ternary_f(card.Legalities.Oathbreaker),
					ternary_f(card.Legalities.StandardBrawl), ternary_f(card.Legalities.Brawl), ternary_f(card.Legalities.Alchemy), ternary_f(card.Legalities.PauperCommander), ternary_f(card.Legalities.Duel),
					ternary_f(card.Legalities.Oldschool), ternary_f(card.Legalities.Premodern), ternary_f(card.Legalities.Predh), card.Defense, card.Loyalty,
					card.EdhrecRank, card.HandModifier, card.LifeModifier, card.ContentWarning)
				if err != nil {
					log.Fatal(err)
				}
				rows.Close()
				inserts++
			}
			lines++
			fmt.Printf("\rLines Complete: %d | Inserts complete: %d", lines, inserts)
			// _, err := db.Query(keywordsSql)
			// if err != nil {
			// 	log.Print(keywordsSql)
			// 	log.Fatal(err)
			// }
			//stmt, err := db.Prepare("INSERT INTO test (text) VALUES ($1)")
			// mtgSetsSql := fmt.Sprintf("INSERT into (set_id, set_code, set_name, set_type, set_uri, set_search_uri, scryfall_set_uri)" +
			// "VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s')", models.MtgSet.SetId, models.MtgSet.SetCode, models.MtgSet.SetName, models.MtgSet.SetType,
			// models.MtgSet.SetUri, models.MtgSet.ScryfallSetUri)
			// _, err := db.Query(mtgSetsSql)
			// if err != nil {
			// 	log.Print(mtgSetsSql)
			// 	log.Fatal(err)
			// }

		}

	}
	elapsed := time.Since(start)
	log.Println(elapsed)
}
