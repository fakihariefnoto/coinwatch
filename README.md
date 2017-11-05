# coinwatch

A bitcoin/alternative coin (alt coin)/cryptocurency watcher to give you annoying notification, it's like your desktop private assistant. and also have capability to auto order or sell (soon)

- currently just support api from -> bitcoin.co.id

# Feature
- fast (not to refresh browser)
- get notificatiin
- terminal friendly, can company you in your terminal tab
- auto update price market
- eye catching interface
- order (soon)
- buy (soon)
- compare price with bittrex, polonix etc. (soon)

# Requeirement
- golang 1.7+
- linux

# How To Build
- Open your terminal
- go get github.com/fakihariefnoto/coinwatch
- goto dir coinwatch
- type `make coin`


# Config outside coinwatch

you can do config inseide coinwatch also but if you want to update it just follow this instruction.
this will just updating config,you can update it from here or from app

## How to use
 Open your terminal and choose the setting do you want to updated

#### Updating bitcoin key to transaction (order, buy etc.)
- ./configcoin key your_key_here (to update bitcoin key app)
- ./configcoin secret your_secret_key_here (to update bitcon secret)

### Updating config path
default path name file is `files/config/config.json`
you can update it with,
`./configcoin path your_path_config`
`./configcoin name your_config_file_name`

### Updating price you want to follow
you can use it with coin shrotname like btc, bch, xzc, etc, eth, ltc-idr
example,
`./configcoin btc 10000000`
