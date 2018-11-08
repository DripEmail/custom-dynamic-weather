package zipcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	assert.True(t, ZipCode("12345").IsValid(), "should be valid")
}

func TestIsValidWithLeadingZeros(t *testing.T) {
	assert.True(t, ZipCode("01234").IsValid(), "should be valid")
}

func TestIsValidWithZipPlusFour(t *testing.T) {
	assert.True(t, ZipCode("01234-1234").IsValid(), "should be valid")
}

func TestIsValidWithLetters(t *testing.T) {
	assert.False(t, ZipCode("ABC123").IsValid(), "should not be valid")
}

func TestIsValidWithWrongLength(t *testing.T) {
	assert.False(t, ZipCode("1234").IsValid(), "should not be valid")
}
