package bitcoincoid

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	_ "net/http"
)

func (b *Btcid) GetPriceMarket(coin string) (result string, err error) {
	url := coinIdrURL[coin]
	ticker := TickerResponse{}

	if url == "" {
		return "", errors.New("your coin id was not match and result empty")
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return "", errors.New("fail to get response from server ->" + err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &ticker)
	if err != nil {
		return "", err
	}

	result = ticker.Ticker.Last

	return

}
