package main

import "testing"

func TestIsValid(t *testing.T) {
	zip := ZipCode("12345")
	if !zip.IsValid() {
		t.Errorf("12345 flagged as invalid")
	}
}

func TestIsValidWithLeadingZeros(t *testing.T) {
	zip := ZipCode("01234")
	if !zip.IsValid() {
		t.Errorf("01234 flagged as invalid")
	}
}

func TestIsValidWithZipPlusFour(t *testing.T) {
	zip := ZipCode("01234-1234")
	if !zip.IsValid() {
		t.Errorf("01234-1234 flagged as invalid")
	}
}

func TestIsValidWithLetters(t *testing.T) {
	zip := ZipCode("ABC12")
	if zip.IsValid() {
		t.Errorf("ABC12 flagged as valid")
	}
}

func TestIsValidWithWrongLength(t *testing.T) {
	zip := ZipCode("1234")
	if zip.IsValid() {
		t.Errorf("1234 flagged as valid")
	}
}
