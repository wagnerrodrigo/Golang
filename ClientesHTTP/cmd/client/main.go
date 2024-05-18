package main

import (
	"flag"
	"fmt"
	"log"
)

const endpoint = "https://dummyjson.com"

func main() {
	action := flag.String("action", "", "Which action to perform")
	id := flag.Int("id", 0, "user id")
	username := flag.String("user", "", "username")
	password := flag.String("pass", "", "password")
	flag.Parse()

	switch *action {
	case "get-user":
		usr, err := getUserById(*id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(usr)
	case "login":
		usr, err := login(*username, *password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(usr)
	case "get-current-user":
		usr, err := login(*username, *password)
		if err != nil {
			log.Fatal(err)
		}

		current, err := getCurrentUser(usr.Token)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(current)
	}

}
