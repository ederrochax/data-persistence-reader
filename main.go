package main

import (
	"log"
	"os"
	"reader/db"
	"reader/reader"
	"time"
)

func main() {
	filePath := "base_teste.txt" // Path of the input file

	file, err := openFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	conn, err := db.NewConnector("postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	startTime := time.Now()

	err = reader.PersistentScanner(file, conn)
	if err != nil {
		log.Fatal(err)
	}

	elapsedTime := time.Since(startTime)
	log.Printf("Total execution time: %s", elapsedTime)
}

func openFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return file, err
}
