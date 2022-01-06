package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lguazo/sftp_client/sftp"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sftp.Conn()

}
