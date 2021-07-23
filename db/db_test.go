package db_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	highscoreDB "github.com/kimuson13/go-poker/db"
)

func TestCreateHighScore(t *testing.T) {
	testDB, cleanFunc := setupTestDB(t, "test1")
	defer cleanFunc()

	h := highscoreDB.HighScore{Name: "test", Score: 100}
	err := highscoreDB.CreateHighScore(testDB, h)
	if err != nil {
		t.Fatal("Error: ", err)
	}

	var gotHighScore highscoreDB.HighScore
	if err := testDB.QueryRow("SELECT * FROM highscore WHERE id = 1").Scan(&gotHighScore.ID, &gotHighScore.Name, &gotHighScore.Score); err != nil {
		t.Fatal("Error: ", err)
	}

	if h.Name != gotHighScore.Name || h.Score != gotHighScore.Score {
		t.Errorf("want set is (name, score) = (%s, %d), but got (%s, %d)",
			h.Name, h.Score, gotHighScore.Name, gotHighScore.Score)
	}
}

func TestUpdateHighScore(t *testing.T) {
	testDB, cleanFunc := setupTestDB(t, "test2")
	defer cleanFunc()

	originalHighScore := highscoreDB.HighScore{Name: "before", Score: 100}
	AfterUpdateHighScore := highscoreDB.HighScore{Name: "after", Score: 200}

	err := highscoreDB.CreateHighScore(testDB, originalHighScore)
	if err != nil {
		t.Fatal("Err: ", err)
	}

	if err := highscoreDB.UpdateHighScore(testDB, AfterUpdateHighScore); err != nil {
		t.Fatal("Err: ", err)
	}

	var gotHighScore highscoreDB.HighScore
	if err := testDB.QueryRow("SELECT * FROM highscore WHERE id = 1").Scan(&gotHighScore.ID, &gotHighScore.Name, &gotHighScore.Score); err != nil {
		t.Fatal("Error: ", err)
	}

	if AfterUpdateHighScore.Name != gotHighScore.Name || AfterUpdateHighScore.Score != gotHighScore.Score {
		t.Errorf("want set is (name, score) = (%s, %d), but got (%s, %d)",
			AfterUpdateHighScore.Name, AfterUpdateHighScore.Score, gotHighScore.Name, gotHighScore.Score)
	}
}

func TestGetHighScore(t *testing.T) {
	testDB, cleanfunc := setupTestDB(t, "test3")
	defer cleanfunc()

	wantHighScore := highscoreDB.HighScore{Name: "testget", Score: 100}
	if err := highscoreDB.CreateHighScore(testDB, wantHighScore); err != nil {
		t.Fatal("Error: ", err)
	}

	gotHighScore, err := highscoreDB.GetHighScore(testDB)
	if err != nil {
		t.Fatal("Error: ", err)
	}

	if wantHighScore.Name != gotHighScore.Name || wantHighScore.Score != gotHighScore.Score {
		t.Errorf("want set is (name, score) = (%s, %d), but got (%s, %d)",
			wantHighScore.Name, wantHighScore.Score, gotHighScore.Name, gotHighScore.Score)
	}
}

func TestIsRegisterWithTrue(t *testing.T) {
	testDB, cleanfunc := setupTestDB(t, "test4")
	defer cleanfunc()

	highScore := highscoreDB.HighScore{Name: "test", Score: 100}
	if err := highscoreDB.CreateHighScore(testDB, highScore); err != nil {
		t.Fatal("Error: ", err)
	}

	want := true
	got := highscoreDB.IsRegister(testDB)

	if want != got {
		t.Errorf("want is %v, but got %v", want, got)
	}
}

func TestIsRegisterWithFalse(t *testing.T) {
	testDB, cleanfunc := setupTestDB(t, "test5")
	defer cleanfunc()

	want := false
	got := highscoreDB.IsRegister(testDB)

	if want != got {
		t.Errorf("want is %v, but got %v", want, got)
	}
}

func setupTestDB(t *testing.T, name string) (*sql.DB, func()) {
	t.Helper()
	testDB, err := highscoreDB.GetDB(name)
	dbName := fmt.Sprintf("%s.db", name)
	if err != nil {
		t.Fatal("Error: ", err)
	}

	return testDB, func() {
		if err := testDB.Close(); err != nil {
			t.Fatal("Error: ", err)
		}
		if err := os.Remove(dbName); err != nil {
			t.Fatal("Error ", err)
		}
	}
}
