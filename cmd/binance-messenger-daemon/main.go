package main

import (
	"fmt"

	"github.com/cryptellation/binance-messenger-service/pkg/server"

	binance "github.com/adshao/go-binance/v2"
)

func main() {
	// Create binance client
	b := binance.NewClient("", "")

	// Create server
	d := server.New(b)

	// Print version
	fmt.Println("Galaxy Messenger Binance Daemon", d.Version())
}
