package api

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/PlatonovSerg/nickname-generato/internal/logic"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GenerateNicknameHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		gender := c.Query("gender")
		genders := []string{"male", "female", "neuter"}
		if gender == "" {
			rand.New(rand.NewSource(time.Now().UnixNano()))
			gender = genders[rand.Intn(len(genders))]
		}
		style := c.DefaultQuery("style", "classic")

		nickname, err := logic.GenerateNickname(db, gender, style)
		if err != nil || nickname == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no suitable words found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"nickname": nickname})
	}
}
