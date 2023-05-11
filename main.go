package main

import (
	"log"

	pricechecker "github.com/yellowgg2/price_check_bot_go/pricechecker"
)

func main() {
	a := pricechecker.AppStore{}
	ra, err := a.LookupApp(pricechecker.RequireInfo{ID: "782438457"})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Price: %v err: %v", ra.Price, err)
}
