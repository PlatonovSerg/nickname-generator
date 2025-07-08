package logic

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/PlatonovSerg/nickname-generato/internal/db"
	"gorm.io/gorm"
)

func GenerateNickname(database *gorm.DB, gender string, style string) (string, error) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var adjectives []db.NamePart
	var nouns []db.NamePart

	resultAdj := database.Where("part_of_speech = ? AND gender = ? AND (style = ? OR style IS NULL)", "adjective", gender, style).Find(&adjectives)
	resultNoun := database.Where("part_of_speech = ? AND gender = ? AND (style = ? or style IS NULL)", "noun", gender, style).Find(&nouns)
	fmt.Printf("adjectives found: %d, error: %v\n", len(adjectives), resultAdj.Error)
	fmt.Printf("nouns found: %d, error: %v\n", len(nouns), resultNoun.Error)
	if len(adjectives) == 0 || len(nouns) == 0 {
		return "", nil
	}
	adj := adjectives[rand.Intn(len(adjectives))].Word
	noun := nouns[rand.Intn(len(nouns))].Word

	return adj + " " + noun, nil
}
