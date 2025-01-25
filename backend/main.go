package main

import (
	"fmt"
	"log"
	"nxg/configs"
	"nxg/internal/api"
)

func main() {
	fmt.Println("Test pack")

	cfg, err := configs.SetupEnv()
	if err != nil {
		log.Fatal("config file is not loaded properly %v\n", err)
	}

	api.StartServer(cfg)

}
