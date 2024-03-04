package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func ForexVerification(currency string) bool {
	_, sourceFilePath, _, _ := runtime.Caller(0)
	dir := filepath.Dir(sourceFilePath)
	file, err := os.Open(filepath.Join(dir, "../reference/physical_currency_list.csv"))

	if err != nil {
		fmt.Println("Error while reading file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
	}

	for _, r := range records {
		if len(r) > 0 && r[0] == currency {
			return true
		}
	}

	return false
}
