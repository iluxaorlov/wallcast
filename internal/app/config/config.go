package config

type Config struct {
	BotToken    string `toml:"bot_token"`
	ChatId      string `toml:"chat_id"`
	DatabaseUrl string `toml:"database_url"`
}

func New() *Config {
	return new(Config)
}
