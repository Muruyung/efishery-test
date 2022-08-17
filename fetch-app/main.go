package main

import (
	"MyAPI/delivery"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	delivery.Run()
}
