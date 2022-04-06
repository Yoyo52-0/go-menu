package main

import "fmt"

func main1() {
	var command string

	fmt.Println("type \"hello\"、\"help\"、\"quit\" for more")

	for {
		fmt.Scanf("%s", &command)
		switch command {
		case "hello":
			fmt.Println("\thi this is your demo :-)")
		case "help":
			fmt.Println("\there's help menu")
		case "quit":
			return
		default:
			fmt.Println("\twrong command!")
		}
	}
}
