package internal

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

func LoadLogging() (*slog.Logger, *os.File, error) {
	logLocation, err := os.OpenFile(`./log_records/logs.jsonl`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening log file: %v", err)
	}

	multi := io.MultiWriter(logLocation, os.Stdout)

	logger := slog.New(slog.NewJSONHandler(multi, nil))
	slog.SetDefault(logger)
	return logger, logLocation, nil
}
