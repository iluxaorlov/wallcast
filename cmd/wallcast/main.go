package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/iluxaorlov/wallcast/internal/app/config"
	"github.com/iluxaorlov/wallcast/internal/app/wallcast"
	"log"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config_file", "configs/wallcast.toml", "path to config file")
}

func main() {
	flag.Parse()

	c := config.New()

	_, err := toml.DecodeFile(configFile, &c)
	if err != nil {
		log.Fatal(err)
	}

	if err := wallcast.Start(c); err != nil {
		log.Fatal(err)
	}
}
