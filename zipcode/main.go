package zipcode

import (
	"regexp"
)

// ZipCode is a United States postal code.
type ZipCode string

var (
	zipRegex = regexp.MustCompile("^\\d{5}(-\\d{4})?$")
)

// IsValid returns boolean based on whether it is a valid zip code
func (zip ZipCode) IsValid() bool {
	return zipRegex.MatchString(string(zip))
}
