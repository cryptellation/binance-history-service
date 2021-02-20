package main

import (
	"fmt"

	"github.com/cryptellation/binance-messenger-service/pkg/server"
	"github.com/cryptellation/binance-messenger-service/pkg/service"
)

func main() {
	// Create binance client
	b := service.New("", "")

	// Create server
	d := server.New(b)

	// Print version
	fmt.Println("Galaxy Messenger Binance Daemon", d.Version())
}
