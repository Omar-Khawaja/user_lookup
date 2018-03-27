package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/user"
)

var users []string

func PopulateUsers() {
	users = append(users, "root", "omar.khawaja")
}

func Lookup(users []string) {
    for _, i := range users {
        u, err := user.Lookup(i)
        if err != nil {
            log.Fatal(err)
        }
        data, err := json.MarshalIndent(u, "", "\t")
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%s\n\n", data)
    }
}

func main() {
	PopulateUsers()
	Lookup(users)
}
