package bot

import (
	"bytes"
	"encoding/json"
	"github.com/iluxaorlov/wallcast/internal/app/model"
	"github.com/iluxaorlov/wallcast/internal/app/store"
	"net/http"
	"time"
)

func (b *Bot) SendMessage(chatId string, img *model.Image, s *store.Store) error {
	type message struct {
		ChatId              string `json:"chat_id"`
		Photo               string `json:"photo"`
		DisableNotification bool   `json:"disable_notification"`
	}

	js, err := json.Marshal(message{ChatId: chatId, Photo: img.Url, DisableNotification: true})
	if err != nil {
		return err
	}

	res, err := http.Post(b.baseUrl+"/sendPhoto", "application/json", bytes.NewBuffer(js))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		time.Sleep(time.Second * 5)

		err := b.SendMessage(chatId, img, s)
		if err != nil {
			return err
		}
	}

	img.Date = time.Now().Format("2006-01-02 15:04:05")

	if err := s.Image().Insert(img); err != nil {
		return err
	}

	return nil
}
