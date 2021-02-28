package main

import (
	"fmt"

	"github.com/cryptellation/binance-history-service/pkg/server"
	binance "github.com/cryptellation/binance.go"
)

func main() {
	// Create binance client
	b := binance.New("", "")

	// Create server
	d := server.New(b)

	// Print version
	fmt.Println("Galaxy Messenger Binance Daemon", d.Version())
}
