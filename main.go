package main

import (
	"flag"
	"log"

	"github.com/wanhuasong/uniportal/config"
	"github.com/wanhuasong/uniportal/router"
)

var cfgFile string

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)

	flag.StringVar(&cfgFile, "c", "./config.json", "Config file")
	flag.Parse()

	if err := config.InitConfig(cfgFile); err != nil {
		panic(err)
	}
	router.Run()
}
