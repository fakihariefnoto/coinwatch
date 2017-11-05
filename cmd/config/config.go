package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	config "github.com/fakihariefnoto/coinwatch/config"
)

func main() {
	conf := config.Get()

	getField := os.Args[1]
	value := os.Args[2]

	if getField == "" && value == "" {
		fmt.Println("Nothing tobe updated..")
		fmt.Println("Bye..")
		time.Sleep(1000)
		os.Exit(1)
	}

	getField = strings.ToLower(getField)

	switch getField {
	case "key":
		conf.Config.API.Key = value
	case "secret":
		conf.Config.API.Secret = value
	case "path":
		conf.Path = value
	case "name":
		conf.FileName = value
	case "btc":
		val, _ := strconv.ParseInt(value, 10, 64)
		conf.Config.Watch.BTC = val
	case "bch":
		val, _ := strconv.ParseInt(value, 10, 64)
		conf.Config.Watch.BCH = val
	case "xzc":
		val, _ := strconv.ParseInt(value, 10, 64)
		conf.Config.Watch.XZC = val
	case "etc":
		val, _ := strconv.ParseInt(value, 10, 64)
		conf.Config.Watch.ETC = val
	case "ech":
		val, _ := strconv.ParseInt(value, 10, 64)
		conf.Config.Watch.ETH = val
	case "ltc":
		val, _ := strconv.ParseInt(value, 10, 64)
		conf.Config.Watch.LTC = val
	default:
		fmt.Println("Nothing tobe updated..")
		fmt.Println("Bye..")
		time.Sleep(1000)
		os.Exit(1)
	}

	err := config.Save(conf)
	if err != nil {
		fmt.Println("Error updating file got -> ", err)
	}

	config.Refresh()

	fmt.Printf("Updating config -%v- has been done, thank you\n", getField)
	fmt.Printf("Updating value to ->> %v \n", value)
	time.Sleep(1000)
	os.Exit(1)

}
