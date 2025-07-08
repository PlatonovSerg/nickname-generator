package main

import (
	"log"

	"github.com/PlatonovSerg/nickname-generato/internal/api"
	"github.com/PlatonovSerg/nickname-generato/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.InitDB("data/words.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	r := gin.Default()
	r.GET("/generate", api.GenerateNicknameHandler(database))
	log.Println("Starting server on :8080")
	r.Run(":8080")
}
