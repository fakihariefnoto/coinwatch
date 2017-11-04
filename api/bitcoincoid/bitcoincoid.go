package bitcoincoid

import (
	"net/http"
	"time"
)

type (
	IBtcid interface {
		GetPriceMarket(string) (string, error)
	}

	Btcid struct {
		Key       string
		SecretKey string
	}

	TickerResponse struct {
		Ticker struct {
			Last string `json:"last"`
			High string `json:"high"`
			Low  string `json:"low"`
			Buy  string `json:"buy"`
			Sell string `json:"sell"`
		} `json:"ticker"`
	}
)

var (
	coinIdrURL map[string]string
	httpClient *http.Client
)

func init() {
	coinIdrURL = map[string]string{
		"btc-idr": "https://vip.bitcoin.co.id/api/btc_idr/ticker",
		"bch-idr": "https://vip.bitcoin.co.id/api/bch_idr/ticker",
		"xzc-idr": "https://vip.bitcoin.co.id/api/xzc_idr/ticker",
		"eth-idr": "https://vip.bitcoin.co.id/api/eth_idr/ticker",
		"etc-idr": "https://vip.bitcoin.co.id/api/etc_idr/ticker",
		"ltc-idr": "https://vip.bitcoin.co.id/api/ltc_idr/ticker",
	}

	// set time out for 5 sec
	httpClient = &http.Client{Timeout: 5 * time.Second}

}

func New() IBtcid {
	return &Btcid{}
}

func (b *Btcid) getHMAC(postData string) string {
	return ""
}
