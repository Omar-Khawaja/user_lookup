package main

import (
	"testing"
)

func TestPopulateNoEmptyStrings(t *testing.T) {
	Populate()
	for _, v := range users {
		if v == "" {
			t.Error("users cannot contain empty string")
		}
	}
}
