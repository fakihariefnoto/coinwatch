package main

import (
	"fmt"
	bitcoincoid "github.com/fakihariefnoto/coinwatch/api/bitcoincoid"
	notif "github.com/fakihariefnoto/coinwatch/notifier"
	"time"
)

var (
	notify notif.INotifier
	sender notif.Notifier
)

func init() {
	notify = notif.New()
	sender = notify.NewConfig(notif.Config{
		Icon:       notif.ANGRY,
		FixedTitle: "Sekilas Info",
	})
}

func main() {
	fmt.Println("Trial send notification")

	bci := bitcoincoid.New()

	for true {

		bch, err := bci.GetPriceMarket("bch-idr")
		if err != nil {
			fmt.Println(err)
		}

		xzc, err := bci.GetPriceMarket("xzc-idr")
		if err != nil {
			fmt.Println(err)
		}

		etc, err := bci.GetPriceMarket("etc-idr")
		if err != nil {
			fmt.Println(err)
		}

		eth, err := bci.GetPriceMarket("eth-idr")
		if err != nil {
			fmt.Println(err)
		}

		ltc, err := bci.GetPriceMarket("ltc-idr")
		if err != nil {
			fmt.Println(err)
		}

		btc, err := bci.GetPriceMarket("btc-idr")
		if err != nil {
			fmt.Println(err)
		}

		body := fmt.Sprintf("BCH -> %v, XZC -> %v\nETH -> %v, ETC -> %v\nLTC -> %v, BTC -> %v\n", bch, xzc, eth, etc, ltc, btc)

		sender.Message("Sekilas Info", body).SendWithIcon(notif.LAUGH)

		time.Sleep(20 * time.Second)
	}
}
