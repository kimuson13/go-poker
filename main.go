package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/kimuson13/go-poker/db"
	"github.com/kimuson13/go-poker/system"
)

var (
	name string
	rate int
)

func init() {
	flag.StringVar(&name, "n", "guest", "player name")
	flag.IntVar(&rate, "r", 1, "poker rating please input 1 to 5")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
	if rate < 0 || 5 < rate {
		if _, err := fmt.Println("Please input 1 to 5"); err != nil {
			log.Fatal(err)
		}
	}
	userInput := system.New()
	dbName := "highscore"
	db, err := db.GetDB(dbName)
	if err != nil {
		if _, err := fmt.Fprintln(os.Stderr, err); err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}

	if err := system.Run(name, db, userInput.Stdin, rate); err != nil {
		if _, err := fmt.Fprintln(os.Stderr, err); err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
}
