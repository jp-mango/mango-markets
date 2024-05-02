package internal

import "fmt"

func ErrDataPull(err error) error {
	return fmt.Errorf("error pulling data: %v", err)
}

func ErrUnmarshalJSON(err error) error {
	return fmt.Errorf("failed to unmarshal JSON response: %v", err)
}
