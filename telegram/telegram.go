package telegram

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"gpparstel/parser"
	"log"
)

const TelApiKey = "111"

func UpdateChannelandBot() (tgbotapi.UpdatesChannel, *tgbotapi.BotAPI) {

	bot, err := tgbotapi.NewBotAPI(TelApiKey)
	if err != nil {
		log.Fatal(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updateChannel, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)

	}
	return updateChannel, bot
}

func PrepareMessage(p *parser.Post) string {
	return fmt.Sprintf(`
%s

%s

%s
%s
Отзыв:
<b>%s</b>

Модель: %s
Прикрепленные фото:
%s
`, p.Picture, p.Size, p.DateOfBuy, p.Measurements, p.Text, p.Model, p.Foto)
}

/*
file := "http://...../img.jpg"
msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, nil)
msg.FileID = file
msg.UseExisting = true
bot.Send(msg)

*/
