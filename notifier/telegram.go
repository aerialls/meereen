package notifier

import (
	"errors"
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	c "github.com/aerialls/meereen/pkg/check"
	n "github.com/aerialls/meereen/pkg/notifier"
	p "github.com/aerialls/meereen/pkg/processor"
)

// Telegram notifier
type Telegram struct {
	token  string
	chatID int64
}

// NewTelegram validates and returns a new Telegram notifier
func NewTelegram(data map[string]string) (n.Notifier, error) {
	token, ok := data["token"]
	if !ok {
		return nil, errors.New("missing token parameter")
	}

	chatIDString, ok := data["chat_id"]
	if !ok {
		return nil, errors.New("missing chat_id parameter")
	}

	chatID, err := strconv.Atoi(chatIDString)
	if err != nil {
		return nil, err
	}

	return &Telegram{
		token:  token,
		chatID: int64(chatID),
	}, nil
}

// Notify Telegram
func (t *Telegram) Notify(check c.Check, state p.State, message string) error {
	bot, err := tgbotapi.NewBotAPI(t.token)
	if err != nil {
		return err
	}

	var tgMessage string
	if state == p.Error {
		tgMessage = fmt.Sprintf("\U0001F6A8 <b>%s</b>\n\n<code>%s</code>", check.GetTitle(), message)
	} else {
		tgMessage = fmt.Sprintf("\u2705 <b>%s</b>", check.GetTitle())
	}

	msg := tgbotapi.NewMessage(t.chatID, tgMessage)
	msg.ParseMode = "html"

	bot.Send(msg)

	return nil
}

func init() {
	n.RegisterNotifier("telegram", NewTelegram)
}
