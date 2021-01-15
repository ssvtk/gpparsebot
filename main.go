package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jackc/pgx"
	"gpparstel/db"
	"gpparstel/parser"
	"gpparstel/telegram"
	"log"
	"time"
)

func main() {
	//Unmarshal config for db connection to  pgx.ConConfig struct
	cn := db.ConfigUnmarshal("config.json")

	conn, err := pgx.Connect(*cn)

	log.Println("Connection successfully done")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, bot := telegram.UpdateChannelandBot()
	for {
		var firstPost parser.Post

		parser.ExampleScrape(&firstPost) // Unmarshal parsed info into struct

		hash := firstPost.GetMD5Hash() //Make hash of this struct.Text

		rows := conn.QueryRow("SELECT hash FROM gp WHERE hash = $1", hash).Scan(nil) //Check if this hash is in database
		// If it is database returns nil (we do not need this data)
		if rows != nil {
			_, err = conn.Exec("INSERT INTO gp (text, size, date, measurement, model, picture, foto, hash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
				firstPost.Text, firstPost.Size, firstPost.DateOfBuy, firstPost.Measurements, firstPost.Model, firstPost.Picture, firstPost.Foto, hash)
			log.Println("Element added to BD")
			text := telegram.PrepareMessage(&firstPost)
			msg := tgbotapi.NewMessage(-389334771, text)
			msg.ParseMode = "HTML"
			m, err := bot.Send(msg)
			log.Println(" Message send")
			if err != nil {
				log.Println(m, err)
			}

		} else {
			log.Println("Already in DB")
		}
		time.Sleep(1 * time.Second)
	}
}
