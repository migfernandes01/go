package main

import (
	"fmt"
	"booking-app/helper"
	"time"
	"sync"
)

// package level variables(all functions can access):

// define variables/const (type inference happens)
var conferenceName = "Go Conference" // alternative way to define a var using type inference
const conferenceTickets = 50
var remainingTickets uint = 50 							// uint is a positive int
var bookings = make([]UserData, 0)			// slice of maps

// create struct
type UserData struct {
	firstName string
	lastName string
	email string
	ticketAmount uint
}

var wg = sync.WaitGroup{}

func main() {

	// call greet users function passing params
	greetUsers()

	// loop through user's input code
	// infinte loop
	// for {
		// call function to get user inuput
		firstName, lastName, email, ticketAmount := getUserInput()		

		// call function to validate user input
		isValidName, isEmailValid, isTicketsAmountValid :=  helper.ValidateUserInput(firstName, lastName, email, ticketAmount, remainingTickets)

		// check validations	
		if isValidName && isEmailValid && isTicketsAmountValid{
			// call function to book tickets
			bookTicket(ticketAmount, firstName, lastName, email)

			// add 1 goroutine to waitgroup
			wg.Add(1)
			// call function to "send" ticket using cncurrency
			go sendTicket(ticketAmount, firstName, lastName, email)

			// call function to print first names
			firstNames := printFirstNames()

			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			fmt.Printf("Thank you %v for booking %v tickets, you will receive them in your email: %v\n", bookings[0], ticketAmount, email)
			fmt.Printf("There are %v tickets left.\n\n", remainingTickets)

			// if no more remaining tickets
			if remainingTickets <= 0 {
				fmt.Println("Our conference is booked. Come back next year")
				// break out of loop (end program)
				// break
			}	
		} else {
			if !isValidName {
				fmt.Println("Invalid name, please try again!")
			} else if !isEmailValid {
				fmt.Println("Invalid email, please try again!")
			} else {
				fmt.Println("Invalid ticket amount, please try again!")
			}
			// skipt to next loop iteration
			// continue
		}
		// waits all other threads on waitgroup to finish before exiting main
		wg.Wait()
	// }
}

// function to print greeting message
func greetUsers() {
	// print type of vars/consts
	fmt.Printf("conferenceName is %T, conferenceTickets is %T\n", conferenceName, conferenceTickets)

	// print template strings
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "of them are still available")
	fmt.Println("Get your tickets here to attend")
}

// function to print first names
// return slice of strings
func printFirstNames() []string {
	// define string slice
	firstNames := []string{}

	// iterate through bookings slice
	// we get the index and value in each iteration
	for _, booking := range bookings {
		// append first name(from struct) to firstNames slice
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

// function to get user input
// returns 3 strings and one positive int
func getUserInput() (string, string, string, uint) {
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

	// return user input
	return firstName, lastName, email, ticketAmount
}

// funtion to perform ticker booking
func bookTicket(ticketAmount uint, firstName string, lastName string, email string) {
	// decrement remaining tickets
	remainingTickets = remainingTickets - ticketAmount

	// populate userData struct
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		ticketAmount: ticketAmount,
	}

	// append map to next position in slice of UserData
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings: %v\n", bookings)
}

// function to simulate sending ticket (simulates async operation)
func sendTicket(ticketAmount uint, firstName string, lastName string, email string) {
	// sleep 10 seconds
	time.Sleep(10 * time.Second)
	// save formatted string into var
	var ticket = fmt.Sprintf("\n%v tickets for %v %v\n", ticketAmount, firstName, lastName)
	fmt.Printf("Sending ticket:\n %v \nto email address: %v\n", ticket, email)
	// tell waitgroup that this goroutine is done
	wg.Done()
}