package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jackc/pgx"
	"gpparstel/db"
	"gpparstel/parser"
	"gpparstel/telegram"
	"log"
	"os"
	"time"
)

func main() {
	//Unmarshal config for db connection to  pgx.ConConfig struct
	cn := db.ConfigUnmarshal("config.json")

	conn, err := pgx.Connect(*cn)

	fmt.Println("Connection successfully done")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	_, bot := telegram.UpdateChannelandBot()
	for {
		var firstPost parser.Post

		parser.ExampleScrape(&firstPost) // Unmarshal parsed info into struct

		hash := firstPost.GetMD5Hash() //Make hash of this struct.Text

		rows := conn.QueryRow("SELECT hash FROM gp WHERE hash = $1", hash).Scan(nil) //Check if this hash is in database
		// If it is database returns nil (we do not need this data)
		now := time.Now()
		if rows != nil {
			_, err = conn.Exec("INSERT INTO gp (text, size, date, measurement, model, picture, foto, hash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
				firstPost.Text, firstPost.Size, firstPost.DateOfBuy, firstPost.Measurements, firstPost.Model, firstPost.Picture, firstPost.Foto, hash)
			fmt.Println(now, " Element added to BD")
			text := telegram.PrepareMessage(&firstPost)
			msg := tgbotapi.NewMessage(-389334771, text)
			msg.ParseMode = "HTML"
			m, err := bot.Send(msg)
			fmt.Println(now, " Message send")
			if err != nil {
				log.Println(m, err)
			}

		} else {
			fmt.Println(now, " Already in DB")
		}
		time.Sleep(30 * time.Second)
	}
}
