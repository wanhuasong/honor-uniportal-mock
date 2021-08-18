package main

import (
	"log"

	"github.com/wanhuasong/uniportal/router"
)

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	router.Run()
}
