package load

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func Logging() (*slog.Logger, *os.File, error) {
	logLocation, err := os.OpenFile(`./log_records/logs.jsonl`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening log file: %v", err)
	}

	logger := slog.New(slog.NewJSONHandler(logLocation, nil))
	slog.SetDefault(logger)
	return logger, logLocation, nil
}

func DB(DB_CONN string) (*sql.DB, error) {
	db, err := sql.Open("postgres", DB_CONN)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func Env() (API_KEY, DB_CONN, ACTIVE_STOCKS string, err error) {
	err = godotenv.Load()
	//err = godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY = os.Getenv("API_KEY")
	DB_CONN = os.Getenv("DB_DSN")
	ACTIVE_STOCKS = os.Getenv("ACTIVE_STOCKS")

	return API_KEY, DB_CONN, ACTIVE_STOCKS, err
}
