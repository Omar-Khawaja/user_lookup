package main

import (
	"testing"
)

func TestPopulate(t *testing.T) {
	Populate()
	for _, v := range users {
		if v == "" {
			t.Error("users cannot contain empty string")
		}
	}
}
