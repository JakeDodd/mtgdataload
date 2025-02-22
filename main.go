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

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
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
	host := os.Getenv("pg_host")
	port, _ := strconv.Atoi(os.Getenv("pg_port"))
	user := os.Getenv("pg_user")
	password := os.Getenv("pg_password")
	dbname := os.Getenv("pg_dbname")
	filename := os.Getenv("data_filename")

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

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	start := time.Now()
	var lines, inserts, setInserts int
	for true {
		var card models.FileCard = models.FileCard{}
		line, err := reader.ReadString('\n')
		if err != nil {

			break
		}

		lines++
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
			if card.Layout != "token" && card.Layout != "art_series" && card.Layout != "double_faced_token" && card.SetType != "memorabilia" {
				existing, err := database.GetCardByOracleIdAndName(card.OracleId, card.Name, db)

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
						fmt.Printf("\n")
						fmt.Printf("Line: \n%s\n", line)
						fmt.Printf("Existing: \n%+v\n", existing)
						fmt.Printf("File Card Card: \n%+v\n", card.FileCardToCard())
						fmt.Printf("File Card: \n%+v\n", card)
						log.Panic("AH")
					}
				}
			}

			existingSet, err := database.GetSetBySetId(card.SetId, db)

			if err != nil && err != database.SetNotFound {
				log.Fatal(err)
			}

			if err == database.SetNotFound {
				err = database.SaveSet(card.FileCardToSet(), db)
				if err != nil {
					log.Fatal(err)
				}
				setInserts++
			} else {
				if !existingSet.CompareSets(card.FileCardToSet()) {
					fmt.Printf("\n")
					fmt.Printf("Line: \n%s\n", line)
					fmt.Printf("Existing: \n%+v\n", existingSet)
					fmt.Printf("File Card Card: \n%+v\n", card.FileCardToSet())
					fmt.Printf("File Card: \n%+v\n", card)
					log.Panic("AH")
				}
			}

			// TODO: put sets stuff here

			// end

			fmt.Printf("\rLines Complete: %d | Inserts complete: %d | Sets Inserted: %d", lines, inserts, setInserts)

		}

	}
	elapsed := time.Since(start)
	log.Println("\n" + elapsed)
}
