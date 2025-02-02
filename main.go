package main

import (
	//"fmt"
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

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
			// colors := array2String(card.Colors)
			// colorIdentity := array2String(card.ColorIdentity)
			// keywords := array2String(card.Keywords)
			// producedMana := array2String(card.ProducedMana)
			// colorIndicator := array2String(card.ColorIndicator)

			// stmt, err := db.Prepare("INSERT into cards (\"object\", oracle_id, card_name, scryfall_uri, layout, mana_cost, cmc, type_line, oracle_text, power,"+
			// 	" toughness, colors, color_identity, keywords, produced_mana, reserved, rulings_uri, legalities, defense, loyalty, color_indicator, card_faces, edhrec_rank, "+
			// 	"hand_modifier, life_modifier, attraction_lights, content_warning) "+
			// 	"VALUES ($1,$2,$3, $4, $5, $6, $7, $8, $9, $10, $11, '{%s}', '{%s}', '{%s}','{%s}', $16, $17, row('%s', '%s', '%s', '%s', '%s', '%s', "+
			// 	"'%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s','%s')::legality, $, $, '{%s}', '{%s}', $, $, $, '{%s}', $)",
			// 	card.Object, card.OracleId, card.Name, card.ScryfallUri, card.Layout, card.ManaCost, strconv.FormatFloat(card.Cmc, 'f', -1, 64), card.TypeLine, card.OracleText, card.Power,
			// 	card.Toughness, colors, colorIdentity, keywords, producedMana, ternary(card.Reserved, "true", "false"), card.RulingsUri,
			// 	card.Legalities.Standard, card.Legalities.Future, card.Legalities.Historic, card.Legalities.Timeless, card.Legalities.Gladiator, card.Legalities.Pioneer,
			// 	card.Legalities.Explorer, card.Legalities.Modern, card.Legalities.Legacy,
			// 	card.Legalities.Pauper, card.Legalities.Vintage, card.Legalities.Penny, card.Legalities.Commander, card.Legalities.Oathbreaker,
			// 	card.Legalities.StandardBrawl, card.Legalities.Brawl, card.Legalities.Alchemy, card.Legalities.PauperCommander, card.Legalities.Duel,
			// 	card.Legalities.Oldschool, card.Legalities.Premodern, card.Legalities.Predh, card.Defense, card.Loyalty, colorIndicator, createCardFacesString(card.CardFaces),
			// 	card.EdhrecRank, card.HandModifier, card.LifeModifier, intArray2String(card.AttractionLights), ternary(card.ContentWarning, "true", "false"))
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
