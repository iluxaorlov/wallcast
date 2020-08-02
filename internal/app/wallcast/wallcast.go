package wallcast

import (
	"database/sql"
	"github.com/iluxaorlov/wallcast/internal/app/bot"
	"github.com/iluxaorlov/wallcast/internal/app/config"
	"github.com/iluxaorlov/wallcast/internal/app/parse"
	"github.com/iluxaorlov/wallcast/internal/app/store"
	_ "github.com/lib/pq"
)

func Start(c *config.Config) error {
	db, err := newDatabase(c)
	if err != nil {
		return err
	}

	defer db.Close()

	s := store.New(db)
	p := parse.New()

	img, err := p.GetImage(s)
	if err != nil {
		return err
	}

	b := bot.New(c.BotToken)

	if err := b.SendMessage(c.ChatId, img, s); err != nil {
		return err
	}

	return nil
}

func newDatabase(c *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", c.DatabaseUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
