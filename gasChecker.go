package main

//TODO: Display gas price in Eth/USD
//TODO: Experiment with other notification types - email, sms?
import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/go-toast/toast"
	"github.com/hrharder/go-gas"
)

func main() {
	//oneEth := big.NewInt(1000000000000000000)
	var userPrice = flag.Int("price", 0, "The price to check for (in gwei)")
	flag.Parse()

	fmt.Println(*userPrice)
	notification := &toast.Notification{
		AppID: "GasChecker",
		Title: "Average gas price",
		Icon:  "C:/dev/go/gaschecker/icon.png",
	}

	prices := make(chan string)
	go func() {
		for {
			averagePrice, err := gas.SuggestGasPrice(gas.GasPriorityAverage)
			if err != nil {
				log.Fatal(err)
			}
			//p := oneEth.Div(oneEth, averagePrice)
			prices <- averagePrice.String()
			time.Sleep(10 * time.Second)

		}
	}()
	for price := range prices {
		notification.Message = fmt.Sprintf("The current average gas price is %s", price[:3]+" Gwei.")
		fmt.Println(notification.Message)
		if fmt.Sprint(*userPrice) >= price[:3] {
			err := notification.Push()
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}
