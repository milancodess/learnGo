package main

import "fmt"

func main() {
	messageLen := 30
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length: ", messageLen)

	if messageLen > maxMessageLen {
		fmt.Printf("Messsage not sent.")
	} else {
		fmt.Printf("Message sent")
	}
}
