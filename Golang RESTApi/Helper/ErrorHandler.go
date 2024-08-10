package helper

import (
	"fmt"
	"log"
)

func ServiceErr(err error, message string) error {
	// Log detailed error for internal debugging
	log.Printf("%s: %v", message, err)

	// Return a more generic error message
	return fmt.Errorf("%s", message)
}

func RepositoryErr(err error, context string) error {
	// Log detailed error for internal debugging
	log.Printf("%s: %v", context, err)

	// Return a more generic error message
	return fmt.Errorf("%s", context)
}
