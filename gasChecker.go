package main

//TODO: Display gas price in Eth/USD
//TODO: Experiment with other notification types - email, sms?
//TODO: CLI frontend to take price threshhold parameters etc.
import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-toast/toast"
	"github.com/hrharder/go-gas"
)

func main() {

	averagePrice, err := gas.SuggestGasPrice(gas.GasPriorityAverage)
	if err != nil {
		log.Fatal(err)
	}
	priceString := strings.TrimRight(averagePrice.String(), "0")

	notification := &toast.Notification{
		AppID:   "GasChecker",
		Title:   "Average gas price",
		Message: fmt.Sprintf("The current average gas price is %s", priceString),
	}

	for {
		if priceString >= "100" {
			if err := notification.Push(); err != nil {
				log.Fatal(err)
			}
			time.Sleep(10 * time.Second)
		}
	}
}
