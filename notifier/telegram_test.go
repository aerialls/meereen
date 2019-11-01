package notifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTelegram(t *testing.T) {
	TelegramWrapper := func(params map[string]string) error {
		_, err := NewTelegram(params)
		return err
	}

	errors := map[error]bool{
		// empty
		TelegramWrapper(map[string]string{}): false,
		// missing "token"
		TelegramWrapper(map[string]string{
			"foo":     "bar",
			"chat_id": "134",
		}): false,
		// chat_id should be an integer
		TelegramWrapper(map[string]string{
			"token":   "xxxx",
			"chat_id": "xxxx",
		}): false,
		// good
		TelegramWrapper(map[string]string{
			"token":   "xxxx",
			"chat_id": "1234",
		}): true,
	}

	for err, ok := range errors {
		if ok {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}
