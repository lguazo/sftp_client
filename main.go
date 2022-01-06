package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	sftp_client "github.com/lguazo/sftp_client/sftp"
	"github.com/pkg/sftp"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn := sftp_client.Conn()

	sc, err := sftp.NewClient(conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to start SFTP subsystem: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Connection Succesfully..")
	}
	defer conn.Close()
	defer sc.Close()

	sftp_client.CheckSftpFile(*sc, os.Getenv("FILE_PATH"))

}
