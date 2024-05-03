package main

import (
	"fmt"
	"os"

	"mangomarkets/internal/load"

	_ "github.com/lib/pq"
)

func main() {
	// Set up logging
	logger, logFile, err := load.Logging()
	if err != nil {
		fmt.Println(fmt.Errorf("error loading logging: %v", err))
		os.Exit(1)
	}
	defer logFile.Close()

	// Load env variables
	_, DB_CONN, err := load.Env()
	if err != nil {
		logger.Error(err.Error())
	}

	// Connect to DB
	db, err := load.DB(DB_CONN)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	//bizniz logic

}
