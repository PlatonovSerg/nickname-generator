package db

type NamePart struct {
	ID           uint `gorm:"primaryKey"`
	Word         string
	PartOfSpeech string
	Gender       string
	Style        string
	Language     string
}
