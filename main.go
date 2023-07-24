package main

import (
	"fmt"
	"golang-service/config"
	"log"

	"github.com/joho/godotenv"
)

type Request struct {
	TypeCard    string
	Description string
}

func main() {
	fmt.Println("Init Project")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memuat file .env")
	}

	db, err := config.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection open : ", db)

}
