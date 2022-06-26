package main

import "fmt"

func main() {
	// define variables/const (type inference happens)
	conferenceName := "Go Conference" // alternative way to define a var using type inference
	const conferenceTickets = 50
	var remainingTickets uint = 50 // uint is a positive int

	// print type of vars/consts
	fmt.Printf("conferenceName is %T, conferenceTickets is %T\n", conferenceName, conferenceTickets)

	// print template strings
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "of them are still available")
	fmt.Println("Get your tickets here to attend")

	// define vars
	var firstName string
	var lastName string
	var email string
	var ticketAmount uint // positive int

	fmt.Print("Enter your first name: ")
	// scan method takes pointer of var
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	// scan method takes pointer of var
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	// scan method takes pointer of var
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets to buy: ")
	// scan method takes pointer of var
	fmt.Scan(&ticketAmount)

	remainingTickets = remainingTickets - ticketAmount

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive them in your email: %v\n", firstName, lastName, ticketAmount, email)
	fmt.Printf("There are %v tickets left.\n", remainingTickets)
}
