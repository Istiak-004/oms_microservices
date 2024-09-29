package main

import (
	"fmt"
	"os"

	"github.com/istiak-004/common/config"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading.env file")
		os.Exit(1)
	}
	config := config.NewConfig()

	fmt.Println(config)
}
