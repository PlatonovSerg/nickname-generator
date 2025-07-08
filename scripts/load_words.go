package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/PlatonovSerg/nickname-generato/internal/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	f, err := os.Open("data/words.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	database, err := gorm.Open(sqlite.Open("data/words.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	database.AutoMigrate(&db.NamePart{})
	headers := records[0]
	for i, rec := range records {
		if i == 0 {
			continue
		} // skip title row
		row := make(map[string]string)
		for j, val := range rec {
			row[headers[j]] = strings.TrimSpace(val)
		}
		part := db.NamePart{
			Word:         row["word"],
			PartOfSpeech: row["part_of_speech"],
			Gender:       row["gender"],
			Style:        row["style"],
			Language:     row["language"],
		}
		database.Create(&part)
	}
	log.Println("Import done! Total records:", len(records)-1)

}
