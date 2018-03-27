package main

import (
	"testing"
)

func TestPopulateUsersNoEmpty(t *testing.T) {
	PopulateUsers()
	for _, v := range users {
		if v == "" {
			t.Error("users cannot contain empty string")
		}
	}
}
