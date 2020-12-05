package db

import (
	"encoding/json"
	"github.com/jackc/pgx"
	"log"
	"os"
)

func ConfigUnmarshal(f string) *pgx.ConnConfig {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	config := new(pgx.ConnConfig)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
