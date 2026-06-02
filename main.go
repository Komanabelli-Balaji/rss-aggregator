package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-chi/chi"
	_ "github.com/go-chi/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("PORT is not found in the environment")
	}

	portString := os.Getenv("PORT")
	fmt.Println("PORT:", portString)
}