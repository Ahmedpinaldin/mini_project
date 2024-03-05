package AnswearHandlers

import (
	"database/sql"
	"log"
	"mini-project/Model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAnswerHandlerByUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newAnswer Model.Answer
		if err := c.BindJSON(&newAnswer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode request body"})
			return
		}

		if newAnswer.JawabanPeserta == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Isi Jawaban Terlebih Dahulu!"})
			return
		}

		err := newAnswer.SaveAnswerByUser(db)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan jawaban"})
			return
		}

		c.JSON(200, newAnswer)
	}
}

func UpdateAnswerHandlerByAdmin(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("Role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User information not found in context"})
			return
		}

		// Cek apakah pengguna adalah admin
		if user != "Admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Only admins are allowed to access this resource"})
			return
		}

		answerID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid answer ID"})
			return
		}

		// Ambil data jawaban dari database berdasarkan ID
		existingAnswer, err := Model.GetAnsweByID(db, answerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve existing answer"})
			return
		}

		// Bind JSON ke struct Answer
		var answerData Model.Answer
		if err := c.BindJSON(&answerData); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode request body"})
			return
		}

		// Validasi data skor
		if answerData.Skor == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Skor harus diisi"})
			return
		}

		// Perbarui data jawaban yang ada
		existingAnswer.Skor = answerData.Skor

		err = existingAnswer.UpdateAnswerByAdmin(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Memberikan Skor"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Skor berhasil diberikan", "data": existingAnswer})
	}
}
