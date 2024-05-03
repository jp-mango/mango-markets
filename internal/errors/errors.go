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

func ErrInvalidAPIRequest(err string) error {
	return fmt.Errorf("api query error: %v", err)
}

func ErrCompanyInfo(err string) error {
	return fmt.Errorf("unable to query company info: %v", err)
}
