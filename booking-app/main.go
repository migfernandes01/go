package main

import (
	"fmt"
	"strings"
)

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

	// array of strings that can hold up to 50 elements
	// var bookings [50]string

	// slice
	var bookings[]string

	// loop through user's input code
	// infinte loop
	for {
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

		// bool var for name validity
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		// bool var for email validity
		isEmailValid := strings.Contains(email, "@")
		// bool var for tickets amount validity
		isTicketsAmountValid := ticketAmount > 0 && ticketAmount < remainingTickets

		// check validations	
		if isValidName && isEmailValid && isTicketsAmountValid{
			remainingTickets = remainingTickets - ticketAmount

			// assign user's input to the first element of bookings array
			// bookings[0] = firstName + " " + lastName

			// append info to next position in slice
			bookings = append(bookings, firstName + " " + lastName)

			// slice info
			// fmt.Printf("\nBookings slice: %v\n", bookings)
			// fmt.Printf("First value: %v\n", bookings[0])
			// fmt.Printf("Slice type: %T\n", bookings)
			// fmt.Printf("Slice length: %v\n\n", len(bookings))

			// define string slice
			firstNames := []string{}

			// iterate through bookings slice
			// we get the index and value in each iteration
			for _, booking := range bookings {
				// splits string with white space as separator (we get an array with 2 elements)
				var names = strings.Fields(booking)
				// append first name(names[0]) to firstNames slice
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("First names of bookings: %v\n", firstNames)
			fmt.Printf("Thank you %v for booking %v tickets, you will receive them in your email: %v\n", bookings[0], ticketAmount, email)
			fmt.Printf("There are %v tickets left.\n\n", remainingTickets)

			// if no more remaining tickets
			if remainingTickets <= 0 {
				fmt.Println("Our conference is booked. Come back next year")
				// break out of loop (end program)
				break
			}	
		} else {
			// show error message
			if !isValidName {
				fmt.Println("Invalid name, please try again!")
			} else if !isEmailValid {
				fmt.Println("Invalid email, please try again!")
			} else {
				fmt.Println("Invalid ticket amount, please try again!")
			}
			// skipt to next loop iteration
			continue
		}
	}
}
