package pricechecker

import (
	"fmt"
	"log"
	"net/url"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func init() {
	var err error

	bot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}
}

func sendMsg(m tgbotapi.MessageConfig) {
	if _, err := bot.Send(m); err != nil {
		log.Printf("[Error] %s", err)
	}
}

func StartBot() {

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message != nil { // ignore any non-Message updates
			ul, err := url.ParseRequestURI(update.Message.Text)

			// TODO: URL parsing 후 ID를 뽑아낸다
			log.Printf("==================== %+v", ul.Path)

			if err != nil {
				textMsg := fmt.Sprintf("[Error] err=%+v, msg=%s", err, update.Message.Text)
				log.Print(textMsg)
				msg.Text = textMsg
				msg.ReplyToMessageID = update.Message.MessageID
				sendMsg(msg)
				continue
			}
		}

		if update.Message.IsCommand() { // ignore any non-command Messages
			switch update.Message.Command() {
			case "help":
				msg.Text = "이 봇은 App Store 가격 변동 알림을 위한 봇입니다."
			default:
				msg.Text = "I don't know that command"
			}
			sendMsg(msg)
		}
	}
}
