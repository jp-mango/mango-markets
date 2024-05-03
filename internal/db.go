package internal

import (
	"context"
	"database/sql"
	"time"
)

func LoadDB(DB_CONN string) (*sql.DB, error) {
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
