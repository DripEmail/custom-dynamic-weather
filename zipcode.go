package main

import (
	"regexp"
)

// ZipCode is a United States postal code.
type ZipCode string

// IsValid returns boolean based on whether it is a valid zip code
func (zip ZipCode) IsValid() bool {
	valid, err := regexp.MatchString("^\\d{5}(-\\d{4})?$", string(zip))
	if err != nil {
		// TODO: Handle errors better.
		panic(err)
	}
	return valid
}
