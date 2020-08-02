package bot

type Bot struct {
	baseUrl string
}

func New(botToken string) *Bot {
	return &Bot{
		baseUrl: "https://api.telegram.org/bot" + botToken,
	}
}
