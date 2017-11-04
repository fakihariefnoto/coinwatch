package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

type (
	// Config struct for unite all config data
	Config struct {
		FileName string
		Path     string
		Config   struct {
			Notif `json:"notif"`
			API   `json:"api"`
			Watch `json:"watch"`
		} `json:"config"`
	}

	// Notif struct for data notifier
	Notif struct {
		Desktop bool `json:"desktop"`
		CLI     bool `json:"cli"`
	}

	// API struct for API key and secret data
	API struct {
		Key    string `json:"key"`
		Secret string `json:"secret"`
	}

	// Watch struct for btc list and watch timer
	Watch struct {
		BTC  int64 `json:"btc"`
		BCH  int64 `json:"bch"`
		XZC  int64 `json:"xzc"`
		ETH  int64 `json:"eth"`
		ETC  int64 `json:"etc"`
		LTC  int64 `json:"ltc"`
		Time int64 `json:"time"`
	}
)

var cfg Config

func init() {

	cfg = Config{
		FileName: "config",
		Path:     "files/config",
	}

	if ok := readJSONConfig(&cfg); ok != nil {
		log.Fatalln("failed to read config -> " + ok.Error())
	}
}

// Get Config struct
func Get() Config {
	return cfg
}

// Save to save config to file json
func Save(config Config) (err error) {

	configbyte, err := json.MarshalIndent(config, " ", "    ")
	if err != nil {
		return errors.New("Error when marshaling config struct: " + err.Error())
	}

	filename := cfg.Path + "/" + cfg.FileName + ".json"

	err = ioutil.WriteFile(filename, configbyte, 0644)
	if err != nil {
		return errors.New("Error when writing to file: " + err.Error())
	}

	return
}

func Refresh() {
	if ok := readJSONConfig(&cfg); ok != nil {
		log.Fatalln("failed to refresh config -> " + ok.Error())
	}

}

func readJSONConfig(config *Config) (err error) {

	fname := config.Path + "/" + config.FileName + ".json"

	bytefile, err := ioutil.ReadFile(fname)
	if err != nil {
		return errors.New("Error when open config err: " + err.Error())
	}

	err = json.Unmarshal(bytefile, &config)
	if err != nil {
		return errors.New("Error when unmarshal config err: " + err.Error())
	}

	return
}
