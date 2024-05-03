package errors

import "fmt"

// Returns an error if http request or reading the response body of request fails
func ErrDataPull(err error) error {
	return fmt.Errorf("error pulling data: %v", err)
}

// Returns an error if unmarshalling json into struct fails
func ErrUnmarshalJSON(err error) error {
	return fmt.Errorf("failed to unmarshal JSON response: %v", err)
}

func ERRInvalidAPIRequest(err string) error {
	return fmt.Errorf("API query error: %v", err)
}
