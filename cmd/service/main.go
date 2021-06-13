package main

import (
	"fmt"
	"os"

	"github.com/cryptellation/binance-history-service/pkg/server"
	"github.com/cryptellation/binance.go/pkg/binance"
)

func run(path string) error {
	config, err := ConfigFromFile(path)
	if err != nil {
		return err
	}

	b := binance.New(config.APIKey, config.SecretKey)
	d := server.New(b)
	fmt.Println("Galaxy Messenger Binance Daemon", d.Version())

	return d.Serve("127.0.0.1:8080")
}

func main() {
	nb := 0

	if err := run("configs/binance-history-service.toml"); err != nil {
		fmt.Println(err)
		nb = 1
	}

	os.Exit(nb)
}
