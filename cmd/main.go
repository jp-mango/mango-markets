package main

import (
	"context"
	"database/sql"
	"log/slog"
	"mangomarkets/internal"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func openDB(DB_CONN string) (*sql.DB, error) {
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

func main() {
	_, DB_CONN, _ := internal.LoadEnv()
	
	logger := slog.New(slog.NewTextHandler(os.Stdout,nil))

	db, err := openDB(DB_CONN)
	if err != nil{
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

}
