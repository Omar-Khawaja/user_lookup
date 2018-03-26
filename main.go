package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/user"
)

func main() {
	
	users := []string{"root", "omar.khawaja"}

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
