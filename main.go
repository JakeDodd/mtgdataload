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
	user     = "mtg"
	password = "mtg"
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
	file, err := os.Open("all-cards-20241022092120.json")
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
			existing, err := database.GetCardByName(card.Name, db)

			if err != nil && err != database.CardNotFound {
				log.Fatal(err)
			}

			if err == database.CardNotFound {
				err = database.SaveCard(card.FileCardToCard(), db)
				if err != nil {
					log.Fatal(err)
				}
				inserts++
			} else {
				if !existing.CompareCards(card.FileCardToCard()) {
					fmt.Printf("%+v\n", existing)
					fmt.Printf("%+v\n", card.FileCardToCard())
					log.Panic("AH")
				}
			}
			lines++
			fmt.Printf("\rLines Complete: %d | Inserts complete: %d", lines, inserts)

		}

	}
	elapsed := time.Since(start)
	log.Println(elapsed)
}
