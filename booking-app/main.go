package main

import "fmt"

func main() {
	// define variables/const (type inference happens)
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	// alternative way to define a var using type inference
	remainingTickets := 50

	// print type of vars/consts
	fmt.Printf("conferenceName is %T, conferenceTickets is %T\n", conferenceName, conferenceTickets)

	// print template strings
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "of them are still available")
	fmt.Println("Get your tickets here to attend")

	// define vars
	var userName string
	var ticketAmount int

	// assign content to vars
	userName = "Tom"
	ticketAmount = 2
	fmt.Printf("User %v booked %v tickets.", userName, ticketAmount)
}
