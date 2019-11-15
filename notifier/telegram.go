package notifier

import (
	"fmt"
	"strconv"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	c "github.com/aerialls/meereen/pkg/check"
	d "github.com/aerialls/meereen/pkg/data"
	n "github.com/aerialls/meereen/pkg/notifier"
	p "github.com/aerialls/meereen/pkg/processor"
)

// Telegram notifier
type Telegram struct {
	mu     sync.Mutex
	token  string
	chatID int64
}

// NewTelegram validates and returns a new Telegram notifier
func NewTelegram(data map[string]string) (n.Notifier, error) {
	token, err := d.GetRequiredParameter(data, "token")
	if err != nil {
		return nil, err
	}

	chatIDString, err := d.GetRequiredParameter(data, "chat_id")
	if err != nil {
		return nil, err
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
	t.mu.Lock()
	defer t.mu.Unlock()

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

	_, err = bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	n.RegisterNotifier("telegram", NewTelegram)
}
