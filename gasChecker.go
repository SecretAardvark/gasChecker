package main

//TODO: Figure out how to display the gas price in gwei without all the extra zeros. There might
//	function for this in the big pkg, or maybe convert the big.Int to a type I can do the
//	math conversion on?
//TODO: Display gas price in Eth/USD
//TODO: Loop the app to check current averageprice every ~30 minutes or so, and notify if
//	below your threshhold price.
//TODO: CLI frontend to take price threshhold parameters etc.
import (
	"fmt"
	"log"

	"github.com/go-toast/toast"
	"github.com/hrharder/go-gas"
)

func main() {

	averagePrice, err := gas.SuggestGasPrice(gas.GasPriorityAverage)
	if err != nil {
		log.Fatal(err)
	}

	notification := &toast.Notification{
		AppID:   "GasChecker",
		Title:   "Average gas price",
		Message: fmt.Sprintf("The current average gas price is %d", averagePrice.Uint64()),
	}

	if err := notification.Push(); err != nil {
		log.Fatal(err)
	}
}
