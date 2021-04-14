package main

//TODO: Display gas price in Eth/USD
//TODO: Experiment with other notification types - email, sms?
import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-toast/toast"
	"github.com/hrharder/go-gas"
)

func main() {
	var userPrice = flag.Int("price", 0, "The price to check for (in gwei)")
	flag.Parse()

	fmt.Printf("Checking for average gas prices under %v\n", *userPrice)
	notification := &toast.Notification{
		AppID: "GasChecker",
		Title: "Average gas price",
		Icon:  "C:/dev/go/gaschecker/icon.png",
	}

	prices := make(chan string)
	//Receive the price from Ethgasstation api, get only the first 3 digits.
	go func() {
		for {
			averagePrice, err := gas.SuggestGasPrice(gas.GasPriorityAverage)
			if err != nil {
				log.Fatal(err)
			}
			price := averagePrice.String()[:3]
			if price[2] == '0' && price[0] != 1 {
				price = price[:2]
			}

			prices <- price
			time.Sleep(10 * time.Second)

		}
	}()
	//Compare the current gas price to the users flag and notify when necessary.
	for price := range prices {
		intPrice, err := strconv.Atoi(price)
		if err != nil {
			log.Fatal(err)
		}
		notification.Message = fmt.Sprintf("The current average gas price is %v Gwei.", intPrice)
		fmt.Println(notification.Message)
		if *userPrice >= intPrice {
			err := notification.Push()
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}
