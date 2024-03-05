package Model

import (
	"database/sql"
	"errors"
	"log"
)

type Answer struct {
	ID             int    `json:"id"`
	UserID         int    `json:"id_user"`
	QuizID         int    `json:"id_quiz"`
	QuestionID     int    `json:"id_pertanyaan"`
	JawabanPeserta string `json:"jawaban_peserta"`
	Skor           int    `json:"skor"`
}

func (a *Answer) SaveAnswerByUser(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO jawaban_peserta (id_user, id_quiz, id_pertanyaan, jawaban_peserta) VALUES (?, ?, ?, ?)", a.UserID, a.QuizID, a.QuestionID, a.JawabanPeserta)
	return err
}

func (a *Answer) UpdateAnswerByAdmin(db *sql.DB) error {
	_, err := db.Exec("UPDATE jawaban_peserta SET skor = ? WHERE id = ?", a.Skor, a.ID)
	log.Print(err)
	return err
}

func (a *Answer) GetAnswerByID(db *sql.DB) error {
	row := db.QueryRow("SELECT id_user, id_quiz, id_pertanyaan, jawaban_peserta, skor FROM jawaban_peserta WHERE id = ?", a.ID)
	err := row.Scan(&a.UserID, &a.QuizID, &a.QuestionID, &a.JawabanPeserta, &a.Skor)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Answer not found")
		}
		return err
	}
	return nil
}

func (a *Answer) DeleteAnswerByAdmin(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM jawaban_peserta WHERE id = ?", a.ID)
	return err
}
