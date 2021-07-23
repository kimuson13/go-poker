package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type HighScore struct {
	ID    int
	Name  string
	Score int
}

func GetDB(name string) (*sql.DB, error) {
	dbName := fmt.Sprintf("%s.db", name)
	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		return nil, err
	}
	const sql = `
	CREATE TABLE IF NOT EXISTS highscore (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		score INTEGER NOT NULL
	);
	`
	if _, err := db.Exec(sql); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateHighScore(db *sql.DB, h HighScore) error {
	const sql = "INSERT INTO highscore(name, score) values (?,?)"
	_, err := db.Exec(sql, h.Name, h.Score)
	if err != nil {
		return err
	}

	return nil
}

func UpdateHighScore(db *sql.DB, h HighScore) error {
	const sql = "UPDATE highscore SET (name, score) = (?,?) WHERE id = 1"
	_, err := db.Exec(sql, h.Name, h.Score)
	if err != nil {
		return err
	}

	return nil
}

func GetHighScore(db *sql.DB) (HighScore, error) {
	var h HighScore
	if err := db.QueryRow("SELECT * from highscore WHERE id = 1").Scan(&h.ID, &h.Name, &h.Score); err != nil {
		return h, err
	}

	return h, nil
}

func IsRegister(db *sql.DB) bool {
	var h HighScore
	if err := db.QueryRow("SELECT * from highscore WHERE id = 1").Scan(&h.ID, &h.Name, &h.Score); err != nil {
		return false
	}

	return true
}
