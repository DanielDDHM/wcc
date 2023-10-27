package main

import (
	"fmt"

	"github.com/DanielDDHM/world-coin-converter/config"
	"github.com/DanielDDHM/world-coin-converter/database"
	"github.com/DanielDDHM/world-coin-converter/server"
)

func main() {
	fmt.Println("Hello World")
	config.Init()
	database.StartDatabase()
	server := server.NewServer()
	server.Run()
}
