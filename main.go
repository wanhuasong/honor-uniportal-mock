package main

import (
	"github.com/wanhuasong/uniportal/db"
	"log"

	"github.com/wanhuasong/uniportal/router"
)

func main() {
	if err := db.InitDefaultStore(); err != nil {
		panic(err)
	}
	defer db.DefaultStore.Close()

	log.SetFlags(log.Llongfile | log.LstdFlags)
	router.Run()
}
