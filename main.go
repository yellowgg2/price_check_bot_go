package main

import (
	"log"

	pricechecker "github.com/yellowgg2/price_check_bot_go/pricechecker"
)

func main() {
	r := pricechecker.AppQuery{ID: "782438457"}
	ra, err := r.LookupApp()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Price: %v err: %v", ra.Price, err)
}
